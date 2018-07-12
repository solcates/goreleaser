package client

import (
	"github.com/goreleaser/goreleaser/config"
	"bytes"
	"os"
	"github.com/goreleaser/goreleaser/context"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/oauth2"
	"net/url"
)

type gitlabClient struct {
	client *gitlab.Client
}

func NewGitlab(ctx *context.Context) (Client, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ctx.Token},
	)
	hClient := oauth2.NewClient(ctx, ts)
	client := gitlab.NewClient(hClient, ctx.Token)
	if ctx.Config.GitLabURLs.API != "" {
		apiurl, err := url.Parse(ctx.Config.GitLabURLs.API)
		if err != nil {
			return &gitlabClient{}, err
		}
		client.SetBaseURL(apiurl.String())
	}

	return &gitlabClient{client: client}, nil
}

func (c *gitlabClient) CreateRelease(ctx *context.Context, body string) (releaseID int64, err error) {
	panic("implement me")
}

func (c *gitlabClient) CreateFile(ctx *context.Context, commitAuthor config.CommitAuthor, repo config.Repo, content bytes.Buffer, path, message string) (err error) {
	panic("implement me")
}

func (c *gitlabClient) Upload(ctx *context.Context, releaseID int64, name string, file *os.File) (err error) {
	panic("implement me")
}
