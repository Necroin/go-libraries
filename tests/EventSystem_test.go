package tests

import (
	"testing"

	"github.com/Necroin/go-libraries/EventSystem"
)

func TestEventSystem_AddHandler(t *testing.T) {
	output := make(map[string]string)
	output["handler"] = "Handler not called"
	event := EventSystem.CreateEvent()
	event.AddHandler("handler", func(name string, arguments ...interface{}) { output[name] = "Handler called" })
	event.Call()
	if output["handler"] == "Handler not called" {
		t.Error("Handler hasn't been added")
	}
}
