package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"

	"github.com/MarkSmersh/go-expenses-tui.git/api/components"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserData(body io.ReadCloser) (UserData, components.ApiError) {
	bytes, _ := io.ReadAll(body)

	var user UserData

	json.Unmarshal(bytes, &user)

	if user.Password == "" || user.Username == "" {
		return user, components.NewApiError(400, "Field password or username is missing")
	}

	return user, nil
}

type User struct {
	username string
	auth     bool
	conn     *pgx.Conn
}

func NewUser(conn *pgx.Conn) User {
	return User{
		username: "",
		auth:     false,
		conn:     conn,
	}
}

// Creates a new user with username and password. Returns a new accessToken and authenticates the user.
func (u *User) SignUp(username string, password string) (string, components.ApiError) {
	hashedPassword := components.HashPassword(password)

	u.auth = true
	u.username = username

	accessToken, err := u.GenerateAccessToken()

	if err != nil {
		return "", err
	}

	_, sqlerr := u.conn.Exec(
		context.Background(),
		"INSERT INTO public.users (username, password) VALUES ($1, $2)",
		username,
		hashedPassword,
	)

	if sqlerr != nil {
		var pgErr *pgconn.PgError
		if errors.As(sqlerr, &pgErr) {
			if pgErr.Code == "23505" {
				return "", components.NewApiError(400, "User with such username already exists")
			} else {
				slog.Error(pgErr.Message)
				return "", components.InternalServerError()
			}
		}
	}

	return accessToken, nil
}

// Authenticates the user with an access token
func (u *User) Auth(accessToken string) components.ApiError {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (any, error) {
		key, err := components.GetJwtSecretKey()

		if err != nil {
			return nil, nil
		}

		return []byte(key), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil || !token.Valid {
		return components.NewApiError(400, "Invalid access token")
	}

	u.username, err = token.Claims.GetSubject()

	if err != nil {
		return components.NewApiError(400, "Invalid access token. Missing claim for subject.")
	}

	u.auth = true

	return nil
}

// Authenticates the user and return an access token
func (u User) LogIn(username string, password string) (string, components.ApiError) {
	var dbPassword string

	err := u.conn.QueryRow(context.Background(), "SELECT password FROM users WHERE username = $1", username).Scan(&dbPassword)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", components.NewApiError(400, "No user with such username")
		}

		slog.Error(err.Error())
		return "", components.InternalServerError()
	}

	if dbPassword != components.HashPassword(password) {
		return "", components.NewApiError(400, "Wrong credentials")
	}

	u.auth = true
	u.username = username

	accessToken, err := u.GenerateAccessToken()

	if err != nil {
		return "", components.InternalServerError()
	}

	return accessToken, nil
}

func (u *User) GenerateAccessToken() (string, components.ApiError) {
	if !u.Authenticated() {
		return "", components.NewApiError(401, "User is not autheticated.")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.username,
	})

	key, err := components.GetJwtSecretKey()

	if err != nil {
		return "", components.InternalServerError()
	}

	accessToken, err := token.SignedString([]byte(key))

	if err != nil {
		slog.Error(
			fmt.Sprintf("Unable to sign a JWT token. %s", err.Error()),
		)
		return "", components.InternalServerError()
	}

	return accessToken, nil
}

func (u User) Authenticated() bool {
	return u.auth
}

func (u User) GetUsername() (string, components.ApiError) {
	if !u.Authenticated() || len(u.username) <= 0 {
		return "", components.Unauthorized()
	}

	return u.username, nil
}
