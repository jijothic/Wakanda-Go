package main_test

import (
	"os"
	"tesing"
)

func TestMain() {
	os.Exit(m.Run())
}

func TestVersion(t *tesing.T) {

}

func TestMissingArgsErr(t *testing.T) {
	err := NewMissingArgsErr("test-cmd")
	assert.Equal(t, "(test-cmd) command is missing required arguments", err.Error())
}