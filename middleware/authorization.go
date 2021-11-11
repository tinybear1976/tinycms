package middleware

import (
	"encoding/base64"
	"errors"
	"strings"

	crsa "github.com/tinybear1976/cross_authenticate"
)

const (
	AUTH_RED  = "Red"
	AUTH_BLUE = "Blue"
)

type Author struct {
	Mode string
	User string
	E    string
}

func (a *Author) authorization() (ok bool, err error) {
	ok = false
	err = nil
	if a == nil {
		return
	}
	if a.Mode == AUTH_RED {
		ts := get_TS(a.User)
		ok, err = crsa.Authenticate(ts, a.E)
	} else {
		t := get_T()
		ok, err = crsa.Authenticate(t, a.E)
	}
	return
}

func get_TS(userid string) string {
	//todo
	return ""
}

func get_T() string {
	//todo
	return ""
}

func getAuthor(auth_str string) (*Author, error) {
	if len(auth_str) <= 0 {
		return nil, errors.New("authorization info is empty")
	}
	spl := strings.Split(auth_str, " ")
	if len(spl) != 2 {
		return nil, errors.New("authorization info split must be to 2 parts")
	}
	if spl[0] != AUTH_RED && spl[0] != AUTH_BLUE {
		return nil, errors.New("unrecognized authentication mode: " + spl[0])
	}
	bytes, err := base64.StdEncoding.DecodeString(spl[1])
	if err != nil {
		return nil, err
	}
	data := strings.Split(string(bytes), ":")
	if len(data) != 2 {
		return nil, errors.New("auth_data split must be to 2 parts")
	}
	// data[0] -> userid
	// data[1] -> E(len=36)
	a := &Author{
		Mode: string(spl[0]),
		User: string(data[0]),
		E:    string(data[1]),
	}
	return a, nil
}

func AuthenticateByRequest(sAuth string) (ok bool, err error) {
	// request.head["authorization"]
	//sAuth = "Red dXNlcjE6cGFzc3dvcmQ="
	a, err := getAuthor(sAuth)
	if err != nil {
		return false, err
	}
	ok, err = a.authorization()
	return
}
