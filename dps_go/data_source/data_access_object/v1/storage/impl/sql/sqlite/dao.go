package sqlite

type Storage struct {
	*AlbumStorage
}

func NewStorage() *Storage {
	return &Storage{
		AlbumStorage: NewAlbumStorage(),
	}
}
