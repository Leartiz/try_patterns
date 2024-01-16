package main

import (
	"dps_go/data_source/data_access_object/v1/storage/impl/sql/sqlite"
	"fmt"
)

func main() {
	client := NewClient(sqlite.NewStorage())
	err := client.PrintAllAlbums()
	if err != nil {
		fmt.Printf("Client err: %v", err)
		return
	}

	// ***

	err = client.PrintOneAlbum(100)
	if err != nil {
		fmt.Printf("Client err: %v", err)
		return
	}

	//...
}
