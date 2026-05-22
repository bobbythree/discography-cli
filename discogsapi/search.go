package discogsapi

import (
	"strings"

	"github.com/bobbythree/discography-cli/models"
)

func QueryAlbums(q string, albums []models.Album) []models.Album {
	var result []models.Album

	query := normalize(q)

	for _, a := range albums {
		artist := strings.ToLower(a.Artist)
		albumName := strings.ToLower(a.AlbumName)

		if strings.Contains(artist, query) ||
			strings.Contains(albumName, query) {
			result = append(result, a)
		}
	}
	return result
}
