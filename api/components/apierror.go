package components

type ApiError interface {
	// Code    int
	// Content string
	Code() int
	Error() string
	ErrorBytes() []byte
}

type errorData struct {
	code    int
	content string
}

func newError(code int, content string) errorData {
	return errorData{
		code:    code,
		content: content,
	}
}

func (e errorData) Code() int {
	return e.code
}

func (e errorData) Error() string {
	return e.content
}

func (e errorData) ErrorBytes() []byte {
	return []byte(e.content)
}

func NewApiError(code int, content string) ApiError {
	return newError(code, content)
}

func BadRequest() ApiError {
	return NewApiError(400, "Bad Request")
}

func Unauthorized() ApiError {
	return NewApiError(401, "Unauthorized")
}

func NotFound() ApiError {
	return NewApiError(404, "Not Found")
}

func InternalServerError() ApiError {
	return NewApiError(500, "Internal Server Error")
}
