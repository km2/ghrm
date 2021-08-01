package ghrm

import (
	"context"
	"fmt"

	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
)

type GHRM struct {
	gh *github.Client
}

type Token struct {
	Token string `json:"token"`
}

func New(token *Token) *GHRM {
	ctx := context.Background()
	sts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token.Token})

	return &GHRM{
		gh: github.NewClient(oauth2.NewClient(ctx, sts)),
	}
}

func (g *GHRM) RemoveRepository(owner, repo string) error {
	ctx := context.Background()
	if _, err := g.gh.Repositories.Delete(ctx, owner, repo); err != nil {
		return fmt.Errorf("failed to delete repository: %w", err)
	}

	return nil
}
