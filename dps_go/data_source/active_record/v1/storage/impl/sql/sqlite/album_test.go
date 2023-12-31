package sqlite

import (
	"fmt"
	"testing"
)

func Test_Album_String(t *testing.T) {
	album := NewAlbum("title", 5)
	fmt.Println(album)

	album.AlbumId = new(int)
	*album.AlbumId = 1
	fmt.Println(album)
}
