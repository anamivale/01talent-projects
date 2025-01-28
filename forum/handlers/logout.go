package handlers

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "session-cookies",
		Value:   "",
		Path:    "/",
		Expires: time.Now().Add(-24 * time.Hour),
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
