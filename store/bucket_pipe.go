// Copyright © 2017 thingful

package store

import (
	"github.com/thingful/device-hub/engine"
	"github.com/thingful/device-hub/proto"
)

type pipeBucket struct {
	bucket
}

type Pipe struct {
	Listener  *proto.Entity
	Endpoints []*proto.Entity
	Profile   Profile
	Uri       string
}

type Profile struct {
	Uid         string
	Name        string
	Description string
	// TODO : make this a semantic triple
	Version string
	Script  engine.Script
}

func (p pipeBucket) CreateOrUpdate(pipe Pipe) error {
	return p.store.Insert(p.bucket, []byte(pipe.Uri), pipe)
}

func (p pipeBucket) List() ([]Pipe, error) {

	pipes := []Pipe{}
	err := p.store.List(p.bucket, &pipes)

	if err != nil {
		return nil, err
	}

	return pipes, nil
}

func (p pipeBucket) Delete(uid string) error {
	return p.store.Delete(p.bucket, []byte(uid))
}
