package glim

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func New() (engine *Engine) {
	engine = &Engine{}

	return
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
