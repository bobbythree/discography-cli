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
	{
		Artist:       "Miles Davis",
		AlbumName:    "Seven Steps to Heaven",
		RecordLabels: []string{"Columbia", "CBS", "Sony"},
		ReleaseYear:  "1963",
		Genre:        "Jazz",
		Style:        "Hard Bop",
		TrackList: []models.Track{
			{TrackNumber: 1, Title: "Basin Street Blues"},
			{TrackNumber: 2, Title: "Seven Steps To Heaven"},
			{TrackNumber: 3, Title: "I Fall In Love Too Easily"},
			{TrackNumber: 4, Title: "So Near, So Far"},
			{TrackNumber: 5, Title: "Baby Won't You Please Come Home"},
			{TrackNumber: 6, Title: "Joshua"},
		},
	},
	{
		Artist:       "Miles Davis",
		AlbumName:    "Bitches Brew",
		RecordLabels: []string{"Columbia", "CBS", "Sony"},
		ReleaseYear:  "1970",
		Genre:        "Jazz",
		Style:        "Fusion",
		TrackList: []models.Track{
			{TrackNumber: 1, Title: "Pharaoh's Dance"},
			{TrackNumber: 2, Title: "Bitches Brew"},
			{TrackNumber: 3, Title: "Spanish Key"},
			{TrackNumber: 4, Title: "John McLaughlin"},
			{TrackNumber: 5, Title: "Miles Runs The Voodoo Down"},
			{TrackNumber: 6, Title: "Sactuary"},
		},
	},
	{
		Artist:       "John Coltrane",
		AlbumName:    "Giant Steps",
		RecordLabels: []string{"Atlantic"},
		ReleaseYear:  "1960",
		Genre:        "Jazz",
		Style:        "Hard Bop",
		TrackList: []models.Track{
			{TrackNumber: 1, Title: "Giant Steps"},
			{TrackNumber: 2, Title: "Cousin Mary"},
			{TrackNumber: 3, Title: "Countdown"},
			{TrackNumber: 4, Title: "Spiral"},
			{TrackNumber: 5, Title: "Syeeda's Song Flute"},
			{TrackNumber: 6, Title: "Naima"},
			{TrackNumber: 7, Title: "Mr. P.C."},
		},
	},
	{
		Artist:       "John Coltrane",
		AlbumName:    "A Love Supreme",
		RecordLabels: []string{"Impulse!", "ABC Records", "MCA"},
		ReleaseYear:  "1965",
		Genre:        "Jazz",
		Style:        "Modal, Spiritual",
		TrackList: []models.Track{
			{TrackNumber: 1, Title: "Part I - Acknowledgement"},
			{TrackNumber: 2, Title: "Part II - Resolution"},
			{TrackNumber: 3, Title: "Part III - Pursuance / Part IV - Psalm"},
		},
	},
}
