package table

type Album struct {
	AlbumId  int
	Title    string
	ArtistId int
}

func MakeAlbum(id, artistId int, title string) Album {
	return Album{
		AlbumId:  id,
		ArtistId: artistId,
		Title:    title,
	}
}
