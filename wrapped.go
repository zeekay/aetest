package aetest

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

// Wrap Instance so we can have nice things.
type WrappedInstance struct {
	Instance
}

// Gets a new app engine context from an instance
func (i *WrappedInstance) NewContext() (context.Context, error) {
	req, err := i.NewRequest("GET", "/", nil)
	if err != nil {
		i.Close()
		return nil, err
	}
	return appengine.NewContext(req), nil
}
