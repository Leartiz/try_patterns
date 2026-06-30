package storage

import (
	"dao_v1/domain/table"
	"dao_v1/storage/dto"
)

type AlbumDao interface {
	SelectAllAlbums() ([]table.Album, error)
	SelectAlbumById(id int) (table.Album, error)
	InsertAlbum(artistId int, title string) (int, error)
	UpdateAlbum(input dto.UpdateAlbumInp) error
	DeleteAlbumById(id int) error
}

type Storage interface {
	AlbumDao
}
