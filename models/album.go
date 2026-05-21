// Package models: data models for albums, artists etc
package models

type Album struct {
	Artist       string
	AlbumName    string
	RecordLabels []string
	ReleaseYear  string
	Genres       []string
	Styles       []string
	TrackList    []Track
}
