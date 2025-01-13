package gapi

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"log"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	xForwardedForHeader        = "x-forwarded-for"
	userAgentHeader            = "user-agent"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	// gRPC 메타데이터 가져오기
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("md: %+v\n", md)
		// HTTP 클라이언트 메타데이터
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		// gRPC 클라이언트 메타데이터
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			mtdt.ClientIP = clientIPs[0]
		}
	}

	// Peer 패키지로 gRPC 클라이언트 IP 가져오기
	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt

}
