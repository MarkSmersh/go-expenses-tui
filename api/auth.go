package api

import (
	"net/http"

	"github.com/MarkSmersh/go-expenses-tui.git/api/models"
)

func (s *Server) AuthRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.AuthGet(w, r)
	case "PUT":
		s.AuthPut(w, r)
	case "POST":
		s.AuthPost(w, r)
	case "OPTIONS":
		s.AuthOptions(w, r)
	default:
		w.WriteHeader(400)
	}
}

// Returns user's access token
func (s *Server) AuthGet(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	user := models.NewUser(s.Conn)

	apierr := user.Auth(cookie.Value)

	if apierr != nil {
		w.WriteHeader(apierr.Code())
		w.Write(apierr.ErrorBytes())
		return
	}

	w.Write([]byte(cookie.Value))
}

// Registration
func (s *Server) AuthPut(w http.ResponseWriter, r *http.Request) {
	userData, err := models.NewUserData(r.Body)

	if err != nil {
		w.WriteHeader(err.Code())
		w.Write(err.ErrorBytes())
		return
	}

	user := models.NewUser(s.Conn)

	accessToken, err := user.SignUp(userData.Username, userData.Password)

	if err != nil {
		w.WriteHeader(err.Code())
		w.Write(err.ErrorBytes())
		return
	}

	cookie := http.Cookie{
		Name:     "access-token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)

	w.WriteHeader(201)
	w.Write([]byte("Account has been created succesfully"))
}

// Login
func (s *Server) AuthPost(w http.ResponseWriter, r *http.Request) {
	userData, err := models.NewUserData(r.Body)

	if err != nil {
		w.WriteHeader(err.Code())
		w.Write(err.ErrorBytes())
		return
	}

	user := models.NewUser(s.Conn)

	accessToken, err := user.LogIn(userData.Username, userData.Password)

	if err != nil {
		w.WriteHeader(err.Code())
		w.Write(err.ErrorBytes())
		return
	}

	cookie := http.Cookie{
		Name:     "access-token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)

	w.WriteHeader(200)
	w.Write([]byte("Loggen in successfully"))
}

func (s *Server) AuthOptions(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("access-token")

	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized"))
		return
	}

	cookie := http.Cookie{
		Name:     "access-token",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	w.WriteHeader(200)
	w.Write([]byte("Logged out"))
}
