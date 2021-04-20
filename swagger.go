package swag

import (
	"errors"
	"sync"
)

// Name is a unique name be used to register swag instance.
const Name = "swagger"

var (
	swaggerMu sync.RWMutex
	swag      map[string]Swagger
)

// Swagger is a interface to read swagger document.
type Swagger interface {
	ReadDoc() string
}

// Register registers swagger for given name.
func Register(name string, swagger Swagger) {
	swaggerMu.Lock()
	defer swaggerMu.Unlock()
	if swagger == nil {
		panic("swagger is nil")
	}

	if swag == nil {
		swag = make(map[string]Swagger)
	}

	swag[name] = swagger
}

// ReadDoc reads swagger document.
func ReadDoc() (string, error) {
	if rv, ok := swag[Name]; ok {
		return rv.ReadDoc(), nil
	}
	return "", errors.New("not yet registered swag")
}

// Read Named Swagger Doc
func ReadDocName(name string) (string, error) {
	if swag == nil {
		return "", errors.New("no swag has yet been registered")
	}
	if rv, ok := swag[name]; ok {
		return rv.ReadDoc(), nil
	}
	return "", errors.New("swag not registered: " + name)
}
