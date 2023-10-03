package middleware

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"kratos-git/helper"
)

func Auth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				auth := tr.RequestHeader().Get("Authorization")
				if auth == "" {
					return nil, fmt.Errorf("No Auth!")
				}
				userClaims, err := helper.AnalyseToken(auth)
				if err != nil {
					return nil, err
				}
				if userClaims.Identity == "" {
					return nil, fmt.Errorf("No Auth!")
				}
				ctx = metadata.NewServerContext(ctx, metadata.New(map[string][]string{
					"identity": {userClaims.Identity},
					"username": {userClaims.Name},
				}))
			}
			return handler(ctx, req)
		}
	}
}
