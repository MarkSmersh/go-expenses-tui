package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	res  *http.Response
	body []byte
}

func NewResponse(res *http.Response) Response {
	body, _ := io.ReadAll(res.Body)

	return Response{
		res:  res,
		body: body,
	}
}

func (r *Response) Res() *http.Response {
	return r.res
}

func (r Response) Body() []byte {
	return r.body
}

func (r Response) Unmarshall(dest any) error {
	return json.Unmarshal(r.body, &dest)
}

func (r Response) IsStatusSuccess() bool {
	if r.res.StatusCode > 299 {
		return false
	}

	return true
}

func (r Response) GetCookie(name string) (*http.Cookie, bool) {
	for _, c := range r.res.Cookies() {
		if c.Name == name {
			return c, true
		}
	}

	return &http.Cookie{}, false
}
