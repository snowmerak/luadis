package luadis

import "github.com/redis/rueidis"

type Client struct {
	conn rueidis.Client
}
