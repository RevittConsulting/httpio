package httpio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
)

type ErrResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Data      interface{} `json:"data"`
	FromCache bool        `json:"from_cache"`
}

func ReadJSON(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, v interface{}) {
	if v == nil {
		WriteErr(w, fmt.Errorf("not found"), http.StatusNotFound)
		return
	}

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(v); err != nil {
		WriteErr(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(buffer.Bytes())
	if err != nil {
		log.Fatalf("error writing JSON to response: %v", err)
	}
}

func WriteCachedJSON(w http.ResponseWriter, v interface{}, fromCache bool) {
	if isNil(v) {
		WriteErr(w, fmt.Errorf("not found"), http.StatusNotFound)
		return
	}

	WriteJSON(w, Response{Data: v, FromCache: fromCache})
}

func WriteErr(w http.ResponseWriter, err error, code int) {
	if err == nil {
		WriteErr(w, fmt.Errorf(http.StatusText(code)), code)
		return
	}

	switch err.Error() {
	case "unauthorised":
		w.WriteHeader(http.StatusUnauthorized)
		WriteJSON(w, ErrResponse{Error: err.Error()})
	default:
		w.WriteHeader(code)
		WriteJSON(w, ErrResponse{Error: err.Error()})
	}
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	value := reflect.ValueOf(i)
	kind := value.Kind()
	return (kind == reflect.Ptr || kind == reflect.Slice || kind == reflect.Map || kind == reflect.Func || kind == reflect.Chan || kind == reflect.Interface) && value.IsNil()
}
