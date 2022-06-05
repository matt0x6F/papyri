package handlers

import (
	"context"
)

type Settings struct {
	ctx context.Context
}

func NewSettings() *Settings {
	return &Settings{}
}

func (s *Settings) startup(ctx context.Context) {
	s.ctx = ctx
}

func (s *Settings) shutdown(ctx context.Context) {}
