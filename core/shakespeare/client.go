package shakespeare

import (
	"context"
	"encoding/json"
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

// Client is an interface for shakespeare api integration
type Client interface {
	Translate(context.Context, string) (*Translation, error)
}

// NewClient returns the default client implementation
func NewClient(uri string) Client {
	return &clientImpl{
		uri:         uri,
		restyClient: resty.New(),
	}
}

type clientImpl struct {
	uri         string
	restyClient *resty.Client
}

func (c *clientImpl) Translate(ctx context.Context, text string) (*Translation, error) {
	resp, err := c.restyClient.R().
		SetContext(ctx).
		EnableTrace().
		SetBody(translationRequest{Text: text}).
		Post(c.uri)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("Expected 200 from shakespeare api, got %d", resp.StatusCode())
	}

	var t Translation
	err = json.Unmarshal(resp.Body(), &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
