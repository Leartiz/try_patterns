package storage

import (
	"dps_go/data_source/data_access_object/v1/domain/table"
	"dps_go/data_source/data_access_object/v1/storage/dto"
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
