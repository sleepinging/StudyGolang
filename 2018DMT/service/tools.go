package service

import (
	"net/http"
	"strconv"
	"errors"
	"net/url"
)

func GetPostInt(name string, w http.ResponseWriter, r *http.Request) (v int, err error) {
	ids, ok := r.PostForm[name]
	if !ok || len(ids) == 0 {
		err = errors.New("缺少" + name + "参数")
		return
	}
	id, err := strconv.ParseInt(ids[0], 10, 32)
	if err != nil {
		err = errors.New(name + "参数需要为整数")
		return
	}
	v = int(id)
	return
}

func GetPostString(name string, w http.ResponseWriter, r *http.Request) (v string, err error) {
	vs, ok := r.PostForm[name]
	if !ok || len(vs) == 0 || vs[0] == "" {
		err = errors.New("缺少" + name + "参数")
		return
	}
	v = vs[0]
	return
}

func GetGetInt(name string, queryForm url.Values) (v int, err error) {
	vs, ok := queryForm[name]
	if len(vs) == 0 || !ok {
		err = errors.New("缺少" + name + "参数")
		return
	}
	v64, err := strconv.ParseInt(vs[0], 10, 32)
	if err != nil {
		err = errors.New(name + "参数需要为整数")
		return
	}
	v = int(v64)
	return
}

func GetGetString(name string, queryForm url.Values) (v string, err error) {
	vs, ok := queryForm[name]
	if !ok || len(vs) == 0 || vs[0] == "" {
		err = errors.New("缺少" + name + "参数")
		return
	}
	v = vs[0]
	return
}
