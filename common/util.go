package common

import (
	"net/http"
	"log"
	"image/jpeg"
	"encoding/json"
)

type(
	appError struct {
		Error string `json:"error"`
		Message string `json:"message"`
		HttpStatus int `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)
//TODO: DisplayAppError - En estos momentos cuando ocurre un error la app devuelve un jason que indica el error, se debe cambiar para que de una respuesta HTML.
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int)  {
	errObj := appError{
		Error: handlerError.Error(),
		Message: message,
		HttpStatus: code,
	}

	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err {
		w.Write(j)
	}

}
