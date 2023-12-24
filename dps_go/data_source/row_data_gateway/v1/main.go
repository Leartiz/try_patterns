package main

import (
	storage "dps_go/data_source/row_data_gateway/v1/storage"
	rowGateway "dps_go/data_source/row_data_gateway/v1/storage/gateway/row/impl"
	storageImpl "dps_go/data_source/row_data_gateway/v1/storage/impl"
	"fmt"
	"log"
)

func fillStorage(storageInstance storage.Storage) {
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
	storageType := storage.MEMORY
	storageInstance := storageImpl.Instance(storage.MEMORY)

	fillStorage(storageInstance)

	// ***

	{
		person, err := rowGateway.MakePerson(storageType, "Joshua", "Bloch", 1)
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

		person, err := rowGateway.FindPersonById(storageType, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("person:", person)
	}
}
