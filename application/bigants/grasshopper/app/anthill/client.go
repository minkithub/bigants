// Package anthill 은 anthill 서비스의 클라이언트와 도메인 어댑터를 구현한다.
package anthill

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Endpoint string
	hc       *http.Client
	// hcs        []*http.Client

	// rr         int
	// clientLock sync.Mutex
}

func NewClient(endpoint string) *Client {
	return &Client{Endpoint: endpoint, hc: new(http.Client)}
}

func (c *Client) doAPIGetRequest(ctx context.Context, path string, params map[string]string, dest interface{}) (err error) {
	q := url.Values{}
	for k, v := range params {
		q.Set(k, v)
	}

	println(c.Endpoint + path + "?" + q.Encode())
	hreq, err := c.newRequest(ctx, "GET", c.Endpoint+path+"?"+q.Encode(), nil)
	if err != nil {
		return err
	}

	var resb bytes.Buffer
	if err = c.sendRequest(hreq, &resb); err != nil {
		return err
	}
	if err = json.Unmarshal(resb.Bytes(), dest); err != nil {
		return fmt.Errorf("Invalid response format(%w): %s", err, resb.String())
	}
	println(resb.String())

	return nil
}

func (c *Client) doAPIPostRequest(ctx context.Context, path string, param interface{}, dest interface{}) (err error) {
	var b bytes.Buffer
	if err = json.NewEncoder(&b).Encode(param); err != nil {
		return err
	}
	hreq, err := c.newRequest(ctx, "POST", c.Endpoint+path, &b)
	if err != nil {
		return err
	}

	var resb bytes.Buffer
	if err = c.sendRequest(hreq, &resb); err != nil {
		return err
	}
	if err = json.Unmarshal(resb.Bytes(), dest); err != nil {
		return fmt.Errorf("Invalid response format(%w): %s", err, resb.String())
	}

	return nil
}

func (c *Client) newRequest(ctx context.Context, method, url string, body io.Reader) (r *http.Request, err error) {
	r, err = http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) getClient() (hc *http.Client) {
	return c.hc
}

func (c *Client) sendRequest(req *http.Request, buff *bytes.Buffer) (err error) {
	r, err := c.getClient().Do(req)
	if err != nil {
		return err
	}
	// var buff bytes.Buffer
	buff.ReadFrom(r.Body) // try to read
	r.Body.Close()

	if 400 <= r.StatusCode && r.StatusCode < 500 {
		return fmt.Errorf("%w: %d(%s, %s)", ErrClient, r.StatusCode, r.Request.URL, buff.String())
	}
	if 500 <= r.StatusCode && r.StatusCode < 600 {
		return fmt.Errorf("%w: %d(%s, %s)", ErrServer, r.StatusCode, r.Request.URL, buff.String())
	}

	return nil
}

var (
	// ErrClient 클라이언트 요청에 문제가 있었다
	ErrClient = fmt.Errorf("anthill: error client")
	// ErrServer 현재 서버 상태에 문제가 있다
	ErrServer = fmt.Errorf("anthill: error server")
)
