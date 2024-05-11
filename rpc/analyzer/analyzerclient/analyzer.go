// Code generated by goctl. DO NOT EDIT.
// Source: analyzer.proto

package analyzerclient

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = analyzer.Request
	Response = analyzer.Response

	Analyzer interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultAnalyzer struct {
		cli zrpc.Client
	}
)

func NewAnalyzer(cli zrpc.Client) Analyzer {
	return &defaultAnalyzer{
		cli: cli,
	}
}

func (m *defaultAnalyzer) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := analyzer.NewAnalyzerClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
