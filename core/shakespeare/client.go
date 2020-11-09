package shakespeare

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// SPClient is an interface for shakespeare api integration
type SPClient interface {
	Translate(context.Context, string) (*Translation, error)
}

// NewClient returns the default client implementation
func NewClient(uri string) SPClient {
	return &clientImpl{
		uri:         uri,
		restyClient: resty.New(), // No retry set as the limits of the shakespeare api is very limited (5 in an hour)
	}
}

type clientImpl struct {
	uri         string
	restyClient *resty.Client
}

func (c *clientImpl) Translate(ctx context.Context, text string) (*Translation, error) {
	n := time.Now()
	logrus.Infof("Starting POST /translate")
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

	duration := time.Since(n)
	logrus.WithField("duration", duration).Infof("Finished POST /translate without issues in %s", duration)

	return &t, nil
}
