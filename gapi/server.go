package gapi

import (
	"fmt"
	db "github.com/soojae/simplebank/db/sqlc"
	"github.com/soojae/simplebank/pb"
	"github.com/soojae/simplebank/token"
	"github.com/soojae/simplebank/util"
	"github.com/soojae/simplebank/worker"
)

// Server serves gRPC requests for the simplebank service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
