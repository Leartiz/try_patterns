package storage

type Album struct {
	AlbumId  int
	Title    string
	ArtistId int
}

func NewPart(title string, artistId int) *Album {
	return &Album{Title: title, ArtistId: artistId}
}

func FindOne(id int) *Album {

}

func FindAll(id int) []*Album {

} 

// -----------------------------------------------------------------------

func (p *Album) Create() error {

}

func (p *Album) Update() error {

}

func (p *Album) Delete() error {

}
