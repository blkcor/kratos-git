// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.21.12
// source: api/git/repo.proto

package git

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationRepoListRepo = "/api.git.Repo/ListRepo"

type RepoHTTPServer interface {
	ListRepo(context.Context, *ListRepoRequest) (*ListRepoReply, error)
}

func RegisterRepoHTTPServer(s *http.Server, srv RepoHTTPServer) {
	r := s.Route("/")
	r.GET("/repo/list", _Repo_ListRepo0_HTTP_Handler(srv))
}

func _Repo_ListRepo0_HTTP_Handler(srv RepoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRepoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRepoListRepo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRepo(ctx, req.(*ListRepoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListRepoReply)
		return ctx.Result(200, reply)
	}
}

type RepoHTTPClient interface {
	ListRepo(ctx context.Context, req *ListRepoRequest, opts ...http.CallOption) (rsp *ListRepoReply, err error)
}

type RepoHTTPClientImpl struct {
	cc *http.Client
}

func NewRepoHTTPClient(client *http.Client) RepoHTTPClient {
	return &RepoHTTPClientImpl{client}
}

func (c *RepoHTTPClientImpl) ListRepo(ctx context.Context, in *ListRepoRequest, opts ...http.CallOption) (*ListRepoReply, error) {
	var out ListRepoReply
	pattern := "/repo/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRepoListRepo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
