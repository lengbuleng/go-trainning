package server

import (
	"errors"
	"strings"
)

type StringRequest struct {
	A string
	B string
}

type Service interface {
	//Concat a and b
	Concat(req StringRequest, ret *string) error

	//a,b common string value
	Diff(req StringRequest, ret *string) error
}

type StringService struct {
}

var StrMaxSize = 20

func (s StringService) Concat(req StringRequest, ret *string) error {
	//test for length overflow
	if len(req.A)+len(req.B) > StrMaxSize {
		*ret = ""
		return errors.New("ErrMaxSize")
	}
	*ret = req.A + req.B
	return nil
}

func (s StringService) Diff(req StringRequest, ret *string) error {
	if len(req.A) < 1 || len(req.B) < 1 {
		*ret = ""
		return nil
	}
	res := ""
	if len(req.A) >= len(req.B) {
		for _, char := range req.B {
			if strings.Contains(req.A, string(char)) {
				res = res + string(char)
			}
		}
	} else {
		for _, char := range req.A {
			if strings.Contains(req.B, string(char)) {
				res = res + string(char)
			}
		}
	}
	*ret = res
	return nil
}
