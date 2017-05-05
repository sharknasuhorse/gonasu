package main

import (
	"testing"
)

func TestAppName(t *testing.T) {
	expect := "Zoo Application"
	actual := Appname()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
