package response

import (
	"github.com/go-chi/render"
	"github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"github.com/rianekacahya/boilerplate/middleware/rest"
	"github.com/rianekacahya/boilerplate/pkg/goerror"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool `json:"success"`
	Error   struct {
		Code    string      `json:"code"`
		Message interface{} `json:"message"`
	} `json:"error,omitempty"`
}

func (SuccessResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ErrorResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func Yay(rw http.ResponseWriter, r *http.Request, data interface{}) {
	var response = new(SuccessResponse)

	response.Success = true
	response.Data = data

	render.Status(r, http.StatusOK)
	if err := render.Render(rw, r, response); err != nil {
		Nay(rw, r, goerror.WithCode(err, goerror.ErrCodeFormatting))
	}
}

func Nay(rw http.ResponseWriter, r *http.Request, err error) {
	var(
		response = new(ErrorResponse)
		debug = r.Context().Value(rest.DebugContextKey).(bool)
	)

	response.Success = false

	switch err.(type) {
	case *goerror.Error:
		render.Status(r, goerror.RestCode(err.(*goerror.Error).Code()))
		response.Error.Code = err.(*goerror.Error).Code()

		switch err.(*goerror.Error).Unwrap().(type) {
		case validation.Errors:
			response.Error.Message = err.(*goerror.Error).Unwrap()
		default:
			response.Error.Message = err.(*goerror.Error).Message()
			if debug {
				response.Error.Message = err.(*goerror.Error).Unwrap().Error()
			}
		}
	default:
		render.Status(r, http.StatusInternalServerError)
		response.Error.Code = goerror.ErrCodeInternalServer
		response.Error.Message = goerror.Message(goerror.ErrCodeInternalServer)
		if debug {
			response.Error.Message = err.Error()
		}
	}

	if err := render.Render(rw, r, response); err != nil {
		http.Error(rw, "unexpected error occurred while processing your request", http.StatusInternalServerError)
	}
}

func NotFound(rw http.ResponseWriter, r *http.Request) {
	var response = new(ErrorResponse)
	response.Success = false
	response.Error.Code = goerror.ErrCodeNotFound
	response.Error.Message = goerror.Message(goerror.ErrCodeNotFound)

	render.Status(r, http.StatusNotFound)
	if err := render.Render(rw, r, response); err != nil {
		http.Error(rw, "unexpected error occurred while processing your request", http.StatusInternalServerError)
	}
}

func MethodNotAllowed(rw http.ResponseWriter, r *http.Request) {
	var response = new(ErrorResponse)
	response.Success = false
	response.Error.Code = goerror.ErrCodeMethodNotAllowed
	response.Error.Message = goerror.Message(goerror.ErrCodeMethodNotAllowed)

	render.Status(r, http.StatusMethodNotAllowed)
	if err := render.Render(rw, r, response); err != nil {
		http.Error(rw, "unexpected error occurred while processing your request", http.StatusInternalServerError)
	}
}
