package test

import (
	"testing"
)

func TestAgenda(t *testing.T) {
	result := Healthz()

	if result != "All good on Agenda API" {
		t.Error("Healthz failed")
	}
}
