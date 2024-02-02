package client

import (
	"context"
	"fmt"

	"github.com/razin99/cq-source-obsidian/obsidian"
	"github.com/rs/zerolog"
)

type Client struct {
	logger   zerolog.Logger
	Spec     Spec
	Obsidian *obsidian.ObsidianClient
}

func (c *Client) ID() string {
	return "obsidian"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	if s.Token == "" {
		return Client{}, fmt.Errorf("Token not provided")
	}
	c := Client{
		logger: logger,
		Spec:   *s,
		Obsidian: &obsidian.ObsidianClient{
			Token:     s.Token,
			UserAgent: s.UserAgent,
		},
	}

	return c, nil
}
