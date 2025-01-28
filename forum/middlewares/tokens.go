package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func GenerateId() string {
	uuid := uuid.NewString()

	return uuid
}

func SetCookies(w http.ResponseWriter, username string) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session-cookies",
		Value:   username,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	})
}

func GetCookies(r *http.Request, name string) (*http.Cookie, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			return nil, errors.New("cookie not foud")
		default:
			return nil, err
		}
	}
	return cookie, nil
}
