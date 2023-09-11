package supabase

import (
	"errors"

	"github.com/supabase/postgrest-go"
)

const REST_URL = "/rest/v1"

type Client struct {
	rest *postgrest.Client
}

type RestOptions struct {
	Schema string
}

type ClientOptions struct {
	Headers map[string]string
	Db      *RestOptions
}

func NewClient(url, key, bearerKey string, options *ClientOptions) (*Client, error) {
	if url == "" || key == "" {
		return nil, errors.New("url and key are required")
	}

	// If bearerKey is not set, use key
	if bearerKey == "" {
		bearerKey = key
	}

	defaultHeaders := map[string]string{
		"Authorization": "Bearer " + bearerKey,
		"apikey":        key,
	}

	client := &Client{}
	schema := "public"
	if options != nil && options.Db != nil {
		if len(options.Headers) > 0 {
			for k, v := range options.Headers {
				defaultHeaders[k] = v
			}
		}
	}
	client.rest = postgrest.NewClient(url+REST_URL, schema, defaultHeaders)

	return client, nil
}

// Wrap postgrest From method
// From returns a QueryBuilder for the specified table.
func (c *Client) From(table string) *postgrest.QueryBuilder {
	return c.rest.From(table)
}

// Wrap postgrest Rpc method
// Rpc returns a string for the specified function.
func (c *Client) Rpc(name, count string, rpcBody interface{}) string {
	return c.rest.Rpc(name, count, rpcBody)
}
