package sdk

import (
	"net/http"
	"bitbucket.org/smartystreets/smartystreets-go-sdk"
)

type SigningClient struct {
	inner      httpClient
	credential smarty_sdk.Credential
}

func NewSigningClient(inner httpClient, credential smarty_sdk.Credential) *SigningClient {
	return &SigningClient{
		inner:      inner,
		credential: credential,
	}
}

func (c *SigningClient) Do(request *http.Request) (*http.Response, error) {
	err := c.credential.Sign(request)
	if err != nil {
		return nil, err
	}
	return c.inner.Do(request)
}
