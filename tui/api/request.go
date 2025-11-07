package api

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/settings"
)

var logger = modules.Logger{File: "app.log"}

func RawRequest(address string, method string, endpoint string, data []byte, cookies []*http.Cookie) (Response, error) {

	client := &http.Client{
		// CheckRedirect: redirectPolicyFunc,
	}

	reader := bytes.NewReader(data)

	req, _ := http.NewRequest(
		method,
		fmt.Sprintf("https://%s%s", address, endpoint),
		reader,
	)

	for _, c := range cookies {
		req.AddCookie(c)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		logger.Logf("%s. Trying to connect with HTTP", err.Error())

		req, err = http.NewRequest(
			method,
			fmt.Sprintf("http://%s%s", address, endpoint),
			reader,
		)

		req.Header.Add("Content-Type", "application/json")

		for _, c := range cookies {
			req.AddCookie(c)
		}

		res, err = client.Do(req)

		if err != nil {
			logger.Logf("%s", err.Error())
			return Response{}, err
		}
	}

	return NewResponse(res), nil
}

// Includes cookies with an access-token and an address (a server) from settings
func Request(method string, endpoint string, data []byte) (Response, error) {
	address, err := settings.GetServer()

	if err != nil {
		logger.Logf("%s", err.Error())
		return Response{}, err
	}

	cookies := []*http.Cookie{}

	accessToken, err := settings.GetAccessToken()

	if err != nil {
		logger.Logf("%s", err.Error())
		return Response{}, err
	}

	cookies = append(cookies, &http.Cookie{
		Name:     "access-token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})

	return RawRequest(address, method, endpoint, data, cookies)
}
