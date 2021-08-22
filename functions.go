package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func apiEncode(w http.ResponseWriter, p interface{}) {
	jsonOut, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		e := apiErrors("fail to ident json", err)
		log.Println(e)
	}
	_, errWrite := w.Write(jsonOut)
	if errWrite != nil {
		e := apiErrors("fail to encode json", err)
		log.Println(e)
	}
}

func apiDecode(r *http.Request, p interface{}) {
	errDecode := json.NewDecoder(r.Body).Decode(p)
	if errDecode != nil {
		e := apiErrors("fail to decode json to struct", errDecode)
		log.Println(e)
	}
}