package luadis

import (
	"context"
	"fmt"

	"github.com/redis/rueidis"
)

type ScriptRequest interface {
	Keys() []string
	Args() []string
}

type ScriptResponse interface {
	FromReply(reply rueidis.RedisResult) error
}

type Script[Req ScriptRequest, Res ScriptResponse] struct {
	script *rueidis.Lua
}

func NewScript[Req ScriptRequest, Res ScriptResponse](script string) *Script[Req, Res] {
	return &Script[Req, Res]{
		script: rueidis.NewLuaScript(script),
	}
}

func (s *Script[Req, Res]) Exec(ctx context.Context, client *Client, request Req) (Res, error) {
	var zero Res
	result := s.script.Exec(ctx, client.conn, request.Keys(), request.Args())
	if err := result.Error(); err != nil {
		return zero, fmt.Errorf("failed to execute script: %w", err)
	}

	if err := zero.FromReply(result); err != nil {
		return zero, fmt.Errorf("failed to parse script response: %w", err)
	}

	return zero, nil
}
