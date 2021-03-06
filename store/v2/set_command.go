package v2

import (
	"time"

	"github.com/coreos/etcd/log"
	"github.com/coreos/etcd/store"
	"github.com/coreos/raft"
)

func init() {
	raft.RegisterCommand(&SetCommand{})
}

// Create command
type SetCommand struct {
	Key        string    `json:"key"`
	Value      string    `json:"value"`
	ExpireTime time.Time `json:"expireTime"`
	Dir        bool      `json:"dir"`
}

// The name of the create command in the log
func (c *SetCommand) CommandName() string {
	return "etcd:set"
}

// Create node
func (c *SetCommand) Apply(context raft.Context) (interface{}, error) {
	s, _ := context.Server().StateMachine().(store.Store)

	// create a new node or replace the old node.
	e, err := s.Set(c.Key, c.Dir, c.Value, c.ExpireTime)

	if err != nil {
		log.Debug(err)
		return nil, err
	}

	return e, nil
}
