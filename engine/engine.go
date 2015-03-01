package engine

import (
	"fmt"
	"os"

	"github.com/denkhaus/imap2spam/rspamd"
	"github.com/denkhaus/tcgl/applog"
)

////////////////////////////////////////////////////////////////////////////////
type Engine struct {
	RspamdConfig *rspamd.Config
}

////////////////////////////////////////////////////////////////////////////////
type ContextBase struct {
	Username  string
	Password  string
	Host      string
	Port      int
	Ephemeral bool
}

type EngineFunc func(engine *Engine) error

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) Execute(fn EngineFunc) error {
	return fn(e)
}

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) initDataStore(ephemeral bool, args ...interface{}) (*DataStore, error) {
	dbPath, err := GetDBPathByArgs(args...)
	if err != nil {
		return nil, fmt.Errorf("Store::GetDBPathByArgs::%s", err)
	}

	if ephemeral {
		applog.Infof("Store::Remove database %s", dbPath)
		os.Remove(dbPath)
	}

	store, err := NewDatastore(dbPath, "UIDMap")
	if err != nil {
		return nil, fmt.Errorf("Storage::%s", err)
	}

	if ephemeral {
		applog.Infof("Store::Start with new database %s", dbPath)
	} else {
		applog.Infof("Store::Use database %s", dbPath)
	}

	return store, nil
}

////////////////////////////////////////////////////////////////////////////////
func NewEngine(config *rspamd.Config) (*Engine, error) {
	e := &Engine{RspamdConfig: config}
	return e, nil
}