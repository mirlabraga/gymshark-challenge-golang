package main

import (
	"testing"

	"github.com/mirlabraga/gymshark-challenge-golang/src/services"
)

func TestIsOrderIs250(t *testing.T) {
	packages := services.Calculation(250)

	if packages[0].Package != 250 {
		t.Error("Calculation failed, the pack expected is 1 pack of 250", packages)
	}

}

func TestIsOrderIs251(t *testing.T) {
	packages := services.Calculation(251)

	if packages[0].Package != 500 {
		t.Error("Calculation failed, the pack expected is 1 pack of 500", packages)
	}

}

func TestIsOrderIs500(t *testing.T) {
	packages := services.Calculation(500)

	if packages[0].Package != 500 {
		t.Error("Calculation failed, the pack expected is 1 pack of 500", packages)
	}

}

func TestIsOrderIs750(t *testing.T) {
	packages := services.Calculation(750)

	if packages[0].Package != 250 || packages[0].Quantity != 1 ||
		packages[1].Package != 500 || packages[1].Quantity != 1 {
		t.Error("Calculation failed, the pack expected is 1 pack of 500", packages)
	}
}

func TestIsOrderIs751(t *testing.T) {
	packages := services.Calculation(751)

	if packages[0].Package != 1000 || packages[0].Quantity != 1 {
		t.Error("Calculation failed, the pack expected is 1 pack of 1000", packages)
	}

}

func TestIsOrderIs1001(t *testing.T) {
	packages := services.Calculation(1001)

	if packages[0].Package != 250 || packages[0].Quantity != 1 ||
		packages[1].Package != 1000 || packages[1].Quantity != 1 {
		t.Error("Calculation failed, the pack expected is 1 pack of 1000 and 1 pack of 250", packages)
	}

}

func TestIsOrderIs1250(t *testing.T) {
	packages := services.Calculation(1250)

	if packages[0].Package != 250 || packages[0].Quantity != 1 ||
		packages[1].Package != 1000 || packages[1].Quantity != 1 {
		t.Error("Calculation failed, the pack expected is 1 pack of 1000 and 1 pack of 250", packages)
	}

}

func TestIsOrderIs9000(t *testing.T) {
	packages := services.Calculation(9000)

	if packages[0].Package != 2000 || packages[0].Quantity != 2 ||
		packages[1].Package != 5000 || packages[1].Quantity != 1 {
		t.Error("Calculation failed, the pack expected is 1 pack of 1000 and 1 pack of 250", packages)
	}

}

func TestIsOrderIs12000(t *testing.T) {
	packages := services.Calculation(12000)

	if packages[0].Package != 2000 || packages[0].Quantity != 1 ||
		packages[1].Package != 5000 || packages[1].Quantity != 2 {
		t.Error("Calculation failed, the pack expected is 2 pack of 5000 and 1 pack of 2000", packages)
	}

}

func TestIsOrderIs12001(t *testing.T) {
	packages := services.Calculation(12001)

	if packages[0].Package != 250 || packages[0].Quantity != 1 ||
		packages[1].Package != 2000 || packages[1].Quantity != 1 ||
		packages[2].Package != 5000 || packages[2].Quantity != 2 {
		t.Error("Calculation failed, the pack expected is 2 pack of 5000 and 1 pack of 2000", packages)
	}

}
