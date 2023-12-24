package storage

import (
	"fmt"
	"testing"
)

func Test_Storage(t *testing.T) {
	storage := Instance()

	companies, err := storage.GetCompanies()
	if err != nil {
		t.Errorf("something went wrong")
	}
	if len(companies) != 0 {
		t.Errorf("something went wrong")
	}

	// ***

	_, err = storage.InsertCompany("msi")
	if err != nil {
		t.Errorf("something went wrong")
	}
	_, err = storage.InsertCompany("envision")
	if err != nil {
		t.Errorf("something went wrong")
	}

	// ***

	companies, err = storage.GetCompanies()
	if err != nil {
		t.Errorf("something went wrong")
	}
	if len(companies) != 2 {
		t.Errorf("something went wrong")
	}
	fmt.Println(companies[0])
	fmt.Println(companies[1])

	//...
}
