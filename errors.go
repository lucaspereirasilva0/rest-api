package main

import (
	"fmt"
	"runtime"
)

type Error struct {
	Op    string `json:"op,omitempty"`
	Msg   string `json:"msg,omitempty"`
	Error string `json:"error,omitempty"`
}

func apiErrors(msg string, err error) Error {
	var nameFunction string
	errReturn := err.Error()

	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		nameFunction = fmt.Sprintf("called from %s", details.Name())
	}

	e := Error{
		Op:    nameFunction,
		Msg:   msg,
		Error: errReturn,
	}

	//return fmt.Errorf("%#v", e)
	return e
}
