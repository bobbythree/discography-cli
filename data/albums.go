// Package data: a mock database for discography data
package data

import "github.com/bobbythree/discography-cli/models"

var Albums = []models.Album{
	{
		Artist:       "Miles Davis",
		AlbumName:    "Kind of Blue",
		RecordLabels: []string{"Columbia", "CBS", "Fontana", "Sony"},
		ReleaseYear:  "1959",
		Genre:        "Jazz",
		Style:        "Modal",
		TrackList: []models.Track{
			{TrackNumber: 1, Title: "So What"},
			{TrackNumber: 2, Title: "Freddie Freeloader"},
			{TrackNumber: 3, Title: "Blue in Green"},
			{TrackNumber: 4, Title: "All Blues"},
			{TrackNumber: 5, Title: "Flamenco Sketches"},
		},
	},
}
