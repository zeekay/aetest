package aetest

import (
	"io"
	"net/http"

	"golang.org/x/net/context"
)

// Instance represents a running instance of the development API Server.
type Instance interface {
	// Close kills the child api_server.py process, releasing its resources.
	io.Closer
	// NewRequest returns an *http.Request associated with this instance.
	NewRequest(method, urlStr string, body io.Reader) (*http.Request, error)
}

// Options is used to specify options when creating an Instance.
type Options struct {
	// AppID specifies the App ID to use during tests.
	// By default, "testapp".
	AppID string
	// StronglyConsistentDatastore is whether the local datastore should be
	// strongly consistent. This will diverge from production behaviour.
	StronglyConsistentDatastore bool
}

// NewContext starts an instance of the development API server, and returns
// a context that will route all API calls to that server, as well as a
// closure that must be called when the Context is no longer required.
func NewContext(opts *Options) (context.Context, *WrappedInstance, error) {
	inst, err := NewInstance(opts)
	if err != nil {
		return nil, nil, err
	}

	wrapped := &WrappedInstance{Instance: inst}

	ctx, err := wrapped.NewContext()
	return ctx, wrapped, err
}

// PrepareDevAppserver is a hook which, if set, will be called before the
// dev_appserver.py is started, each time it is started. If aetest.NewContext
// is invoked from the goapp test tool, this hook is unnecessary.
var PrepareDevAppserver func() error
