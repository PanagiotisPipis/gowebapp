package httpsrv

import (
	"net/http"
)

func (s *Server) handlerHealth(w http.ResponseWriter, r *http.Request) {
	session, _ := s.store.Get(r, "session.id")
	authenticated := session.Values["authenticated"]
	if authenticated != nil && authenticated != false {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
    return
	}
}