package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	db "github.com/soojae/simplebank/db/sqlc"
	"github.com/soojae/simplebank/token"
	"github.com/soojae/simplebank/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setUpRouter()

	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.New()
	router.Use(gin.Recovery())

	// zerolog 미들웨어 추가
	router.Use(func(c *gin.Context) {
		startTime := time.Now()

		// 다음 핸들러 처리
		c.Next()

		// 요청 처리 완료 후 로깅
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		if c.Writer.Status() >= 400 {
			log.Error().
				Str("client_ip", clientIP).
				Str("method", method).
				Str("path", path).
				Int("status_code", statusCode).
				Dur("duration (ns)", duration).
				Str("error", c.Errors.String()).
				Msg("HTTP request failed")
		} else {
			log.Info().
				Str("client_ip", clientIP).
				Str("method", method).
				Str("path", path).
				Int("status_code", statusCode).
				Dur("duration (ns)", duration).
				Msg("HTTP request processed")
		}
	})

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specified address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
