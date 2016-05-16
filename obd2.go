package obd2

import (
	"bufio"
	"errors"
	"io"
)

type Request struct {
	Mode byte
	Args []byte
}
type Response []byte

type Transport interface {
	RoundTrip(req *Request) (*Response, error)
}

type Client struct {
	t Transport
}

func NewClient(t Transport) *Client {
	return &Client{t: t}
}
