package discogsapi

import (
	"github.com/bobbythree/discography-cli/models"
)

func searchApi(q string) []models.Album {
	var result []models.Album

	query := normalize(q)

	return result
}
