package main

import (
	"dps_go/data_source/row_data_gateway/v1/storage"
	"fmt"
	"log"
)

func fillStorage() {
	storageInstance := storage.Instance()
	_, err := storageInstance.InsertCompany("msi")
	if err != nil {
		log.Fatal(err)
	}
	_, err = storageInstance.InsertCompany("envision")
	if err != nil {
		log.Fatal(err)
	}
	//...

	companies, err := storageInstance.GetCompanies()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(companies)
}

func main() {
	fillStorage()
	storageInstance := storage.Instance()

	// ***

	{
		person, err := storage.MakePerson("Joshua", "Bloch", 1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("person:", person)

		person.UpdateWithCompanyId(0)
		fmt.Println("person:", person)

		person.SetFirstName("Alex")
		fmt.Println("person:", person)
		person.Update() // <--- commit changes!
	}

	// ***

	{
		persons, err := storageInstance.GetPersons()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("persons:", persons)

		person, err := storage.FindPersonById(0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("person:", person)
	}
}
