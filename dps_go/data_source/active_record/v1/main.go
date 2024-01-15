package main

import (
	"dps_go/data_source/active_record/v1/storage/impl/sql/sqlite"
	"fmt"
)

func main() {
	album, err := sqlite.FindOne(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(album)

	// ***

	album.Title = "New Title"
	err = album.Update()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(album)

	// ***

	album.Title = "New Title 2"
	album.ArtistId = 4
	err = album.Create()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(album)

	// ***

	fmt.Println("[OK]")
}
