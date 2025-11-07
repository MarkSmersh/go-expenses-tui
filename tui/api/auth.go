package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LogIn(destination string, username string, password string) (string, error) {

	jsonString, _ := json.Marshal(Credentials{
		Username: username,
		Password: password,
	})

	res, err := RawRequest(
		destination,
		"POST",
		"/auth/",
		jsonString,
		[]*http.Cookie{},
	)

	if err != nil {
		logger.Logf("%s", err.Error())
		return "", err
	}

	if !res.IsStatusSuccess() {
		err = errors.New(
			fmt.Sprintf("Unable to log in into. %s. %s", res.Res().Status, string(res.Body())),
		)
		logger.Logf("%s", err.Error())
		return "", err
	}

	cookie, ok := res.GetCookie("access-token")

	if !ok || len(cookie.Value) <= 0 {
		errStr := "Unable to retrieve an access-token cookie while logging in"
		logger.Logf("%s", errStr)
		return "", errors.New(errStr)
	}

	return cookie.Value, nil
}

func SignUp(destination string, username string, password string) (string, error) {
	jsonString, _ := json.Marshal(Credentials{
		Username: username,
		Password: password,
	})

	res, err := RawRequest(
		destination,
		"PUT",
		"/auth/",
		jsonString,
		[]*http.Cookie{},
	)

	if err != nil {
		logger.Logf("%s", err.Error())
		return "", err
	}

	if !res.IsStatusSuccess() {
		err = errors.New(
			fmt.Sprintf("Unable to sign up. %s. %s", res.Res().Status, string(res.Body())),
		)
		logger.Logf("%s", err.Error())
		return "", err
	}

	cookie, ok := res.GetCookie("access-token")

	if !ok || len(cookie.Value) <= 0 {
		errStr := "Unable to retrieve an access-token cookie while sign up"
		logger.Logf("%s", errStr)
		return "", errors.New(errStr)
	}

	return cookie.Value, nil
}
