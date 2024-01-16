package sqlite

import (
	"dps_go/data_source/data_access_object/v1/domain/table"
	"dps_go/data_source/data_access_object/v1/storage/dto"

	_ "github.com/mattn/go-sqlite3"
)

type AlbumStorage struct{}

func NewAlbumStorage() *AlbumStorage {
	return &AlbumStorage{}
}

// -----------------------------------------------------------------------

func (a *AlbumStorage) SelectAllAlbums() ([]table.Album, error) {
	db, err := openDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM Album ORDER BY AlbumId;`)
	if err != nil {
		return nil, err
	}

	albums := []table.Album{}
	for rows.Next() {
		album := table.Album{}
		err := rows.Scan(&album.AlbumId, &album.Title, &album.ArtistId)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func (a *AlbumStorage) SelectAlbumById(id int) (table.Album, error) {
	db, err := openDatabase()
	if err != nil {
		return table.Album{}, err
	}
	defer db.Close()

	row := db.QueryRow(
		`SELECT * FROM Album WHERE AlbumId = $1`,
		id,
	)
	album := table.Album{}
	err = row.Scan(&album.AlbumId, &album.Title, &album.ArtistId)
	if err != nil {
		return table.Album{}, err
	}

	return album, nil
}

func (a *AlbumStorage) InsertAlbum(artistId int, title string) (int, error) {
	//...

	return 0, ErrNotImplemented
}

func (a *AlbumStorage) UpdateAlbum(input dto.UpdateAlbumInp) error {
	//...

	return ErrNotImplemented
}

func (a *AlbumStorage) DeleteAlbumById(id int) error {
	//...

	return ErrNotImplemented
}
