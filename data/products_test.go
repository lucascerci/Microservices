package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "cerci",
		Price: 1.00,
		SKU:   "abs-abc-sef",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
