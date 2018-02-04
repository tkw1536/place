// auth package sets a cookie on the first request sent without a cookie, and then expects it on every subsequent call

package auth

import (
	"github.com/google/uuid"
	"log"
	"net/http"
)

const (
	cookieName = "sticky"
)

// StickyUser is an HTTP handler that will use a cookie to ensure only the browser that sent the first request can access any resource
type StickyUser struct {
	cookieValue string
	Next        http.Handler
}

func (s *StickyUser) sendUnauthorized(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "You're not authorized to view this page.", http.StatusUnauthorized)
}

func (s *StickyUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userCookie, cookieErr := r.Cookie(cookieName)

	if s.cookieValue == "" {
		uuid, err := uuid.NewRandom()
		if err != nil {
			log.Printf("could not create sticky cookie")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.cookieValue = uuid.String()

		stickyCookie := http.Cookie{
			Name:  cookieName,
			Value: s.cookieValue,
		}
		http.SetCookie(w, &stickyCookie)
	} else {
		if cookieErr != nil {
			s.sendUnauthorized(w, r)
			return
		}

		if userCookie.Value != s.cookieValue {
			s.sendUnauthorized(w, r)
			return
		}
	}

	s.Next.ServeHTTP(w, r)
}
