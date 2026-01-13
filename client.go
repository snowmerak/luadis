package luadis

import (
	"fmt"

	"github.com/redis/rueidis"
)

type Options = rueidis.ClientOption

type Client struct {
	conn rueidis.Client
}

func New(opt Options) (*Client, error) {
	conn, err := rueidis.NewClient(opt)
	if err != nil {
		return nil, fmt.Errorf("failed to create redis client: %w", err)
	}

	return &Client{conn: conn}, nil
}

func With(conn rueidis.Client) *Client {
	return &Client{conn: conn}
}
