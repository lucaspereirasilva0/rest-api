package tools

import (
	"encoding/json"
	"github.com/lucaspereirasilva0/rest-api/internal/errors"
	"log"
	"net/http"
)

func ApiEncode(w http.ResponseWriter, p interface{}) {
	jsonOut, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		e := errors.New("fail to ident json", err)
		log.Println(e)
	}
	_, errWrite := w.Write(jsonOut)
	if errWrite != nil {
		e := errors.New("fail to encode json", err)
		log.Println(e)
	}
}

func ApiDecode(r *http.Request, p interface{}) {
	errDecode := json.NewDecoder(r.Body).Decode(p)
	if errDecode != nil {
		e := errors.New("fail to decode json to struct", errDecode)
		log.Println(e)
	}
}
