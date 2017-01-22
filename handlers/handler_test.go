package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"io"
)

func newBogusGlobal() *config.Global {
	return &config.Global{}
}

type bogusConsumer struct {
	ctx ctx.Global
}

func newBogusConsumer() *bogusConsumer {
	return &bogusConsumer{}
}

func (bc *bogusConsumer) Consume(_ io.Writer, ctx ctx.Global) error {
	bc.ctx = ctx
	return nil
}

type bogusRequest struct{}

func newBogusRequest() *bogusRequest {
	return &bogusRequest{}
}

func (br *bogusRequest) Path() string                { return "/" }
func (br *bogusRequest) Subtree() string             { return "/" }
func (br *bogusRequest) Redirect(_ string)           {}
func (br *bogusRequest) Write(b []byte) (int, error) { return len(b), nil }
func (br *bogusRequest) Error(_ error)               {}

type bogusReverser struct{}

func newBogusReverser() *bogusReverser {
	return &bogusReverser{}
}

func (br *bogusReverser) Root() string                                    { return "/" }
func (br *bogusReverser) Static(_ string) string                          { return "/" }
func (br *bogusReverser) RepoRoot(_ models.Repo) string                   { return "/" }
func (br *bogusReverser) RepoPlain(_ models.Repo, _ string) string        { return "/" }
func (br *bogusReverser) RepoTree(_ models.Repo, _ string, _ bool) string { return "/" }
func (br *bogusReverser) RepoLog(_ models.Repo) string                    { return "/" }
func (br *bogusReverser) RepoRefs(_ models.Repo) string                   { return "/" }
