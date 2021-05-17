package main

import (
	"testing"

	"github.com/mirlabraga/gymshark-challenge-golang/src/services"
)

func TestIsOrderIs250(t *testing.T) {
	packages := services.Calculation(250)

	if packages[0].Package != 250 {
		t.Error("Calculation failed, expected 250", packages)
	}

}

func TestIsOrderIs251(t *testing.T) {
	packages := services.Calculation(251)

	if packages[0].Package != 500 {
		t.Error("Calculation failed, expected 500", packages)
	}

}
