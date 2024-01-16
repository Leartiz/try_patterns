package main

import (
	"dps_go/data_source/data_access_object/v1/storage"
	"fmt"
)

type Client struct { // <--- BusinessObject
	storage storage.Storage
}

func NewClient(storage storage.Storage) *Client {
	return &Client{
		storage: storage,
	}
}

// public
// -----------------------------------------------------------------------

func (c *Client) PrintAllAlbums() error {
	albums, err := c.storage.SelectAllAlbums()
	if err != nil {
		return err
	}

	for _, album := range albums {
		fmt.Println(album)
	}
	return nil
}

func (c *Client) PrintOneAlbum(id int) error {
	album, err := c.storage.SelectAlbumById(id)
	if err != nil {
		return err
	}
	fmt.Println(album)
	return nil
}
