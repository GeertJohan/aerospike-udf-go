package main

import "github.com/GeertJohan/go.rice"

var boxTemplates *rice.Box

func setupRice() {
	boxTemplates = rice.MustFindBox("templates")
}
