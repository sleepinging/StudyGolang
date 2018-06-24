// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"twt/mytools"
)

func main() {
	var le *walk.LineEdit
	var wv *walk.WebView
	cp,_:=mytools.GetCurrentPath()
	MainWindow{
		Icon:    Bind("'../img/' + icon(wv.URL) + '.ico'"),
		Title:   "Walk WebView Example'",
		MinSize: Size{800, 600},
		Layout:  VBox{MarginsZero: true},
		Children: []Widget{
			LineEdit{
				AssignTo: &le,
				Text:     Bind("wv.URL"),
				OnKeyDown: func(key walk.Key) {
					if key == walk.KeyReturn {
						wv.SetURL(le.Text())
					}
				},
			},
			WebView{
				AssignTo: &wv,
				Name:     "wv",
				URL:      `file:///`+cp+`index.html`,
			},
		},
		Functions: map[string]func(args ...interface{}) (interface{}, error){
			"icon": func(args ...interface{}) (interface{}, error) {
				if strings.HasPrefix(args[0].(string), "https") {
					return "check", nil
				}

				return "stop", nil
			},
		},
	}.Run()
}
