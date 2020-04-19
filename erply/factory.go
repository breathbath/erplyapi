package erply

import (
	"github.com/breathbath/go_utils/utils/env"
	"github.com/erply/api-go-wrapper/pkg/api"
)

//ClientFactory interface that abstracts creation of Erply API
type ClientFactory interface {
	CreateClient() (cl api.IClient, err error)
}

//DefaultClientFactory default implementation of ClientFactory
type DefaultClientFactory struct {
	AuthProvider AuthProvider
}

//CreateClient implements ClientFactory interface
func (c DefaultClientFactory) CreateClient() (cl api.IClient, err error) {
	session, err := c.AuthProvider.GetSession()
	if err != nil {
		return
	}

	clientCode, err := env.ReadEnvOrError("ERPLY_CLIENT_CODE")
	if err != nil {
		return
	}

	cl = api.NewClient(session.sessionID, clientCode, nil)

	return
}
