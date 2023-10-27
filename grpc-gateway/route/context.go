package httprouter

import "net/http"

// HandlersChain defines a HandlerFunc slice.
type HandlersChain []Handle

// Handle defines the handler used by middleware as return value.
func (h *HandlersChain) Handle(w http.ResponseWriter, r *http.Request, p Params) {
	for _, handler := range *h {
		handler(w, r, p)
	}
}
