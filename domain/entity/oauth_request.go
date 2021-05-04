package entity

import "net/http"

func (e *RequestToken) Bind(r *http.Request) error {
	// validate request body
	if err := e.Validate(); err != nil {
		return err
	}

	return nil
}
