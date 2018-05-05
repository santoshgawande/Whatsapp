package servicehandlers

import (
	"net/http"
)

type ChatHandler struct {
}

func (p ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

func (p ChatHandler) Get(r *http.Request) (string, int) {
	return "GET Called", 200
}

func (p ChatHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

func (p ChatHandler) Post(r *http.Request) (string, int) {
	return "POST Called", 200

}
