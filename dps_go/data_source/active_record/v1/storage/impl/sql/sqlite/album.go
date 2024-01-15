package sqlite

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Album struct {
	AlbumId  *int
	Title    string
	ArtistId int
}

// ctors
// -----------------------------------------------------------------------

func NewAlbum(title string, artistId int) *Album {
	return &Album{Title: title, ArtistId: artistId}
}

func FindOne(id int) (*Album, error) {
	db, err := openDatabase() // <--- unsafe!
	defer db.Close()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(
		fmt.Sprintf("SELECT * FROM Album WHERE AlbumId = %v;", id))
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		var album Album
		album.AlbumId = new(int)
		err = rows.Scan(album.AlbumId, &album.Title, &album.ArtistId)
		if err != nil {
			return nil, err
		}

		return &album, nil
	}
	return nil, fmt.Errorf("Album with id %v not found", id)
}

func (p *Album) String() string {
	strAlbumId := "<unknown>"
	if p.AlbumId != nil {
		strAlbumId = strconv.Itoa(*p.AlbumId)
	}

	return fmt.Sprintf("{%v %v %v}",
		strAlbumId, p.Title, p.ArtistId)
}

// active record
// -----------------------------------------------------------------------

func (p *Album) Create() error {
	if err := validateTitle(p.Title); err != nil {
		return err
	}

	db, err := openDatabase()
	defer db.Close()
	if err != nil {
		return err
	}

	result, err := db.Exec(`INSERT INTO Album (Title, ArtistId) VALUES  ($1, $2)`,
		p.Title, p.ArtistId)
	if err != nil {
		return err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.AlbumId = new(int)
	*p.AlbumId = int(lastInsertedId)
	return nil
}

func (p *Album) Update() error {
	if p.AlbumId == nil {
		return fmt.Errorf("Album id is nil")
	}
	if err := validateTitle(p.Title); err != nil { // <--- domain logic?
		return err
	}

	db, err := openDatabase() // <--- unsafe!
	defer db.Close()
	if err != nil {
		return err
	}

	_, err = db.Exec(`UPDATE Album SET Title = $1 WHERE AlbumId = $2;`,
		p.Title, *p.AlbumId)
	if err != nil {
		return err
	}

	return nil
}

func (p *Album) Delete() error {
	if p.AlbumId == nil {
		return fmt.Errorf("Album id is nil")
	}

	db, err := openDatabase() // <--- unsafe!
	defer db.Close()
	if err != nil {
		return err
	}

	_, err = db.Exec(`DELETE FROM Album WHERE AlbumId = $1;`,
		*p.AlbumId)
	if err != nil {
		return err
	}

	return nil
}

// private
// -----------------------------------------------------------------------

func validateTitle(title string) error {
	if len(title) == 0 {
		return fmt.Errorf("Title is empty")
	}
	return nil
}

func openDatabase() (*sql.DB, error) {
	return sql.Open("sqlite3", "../../../../chinook.sqlite")
}
