// Package models: data models for albums, artists etc
package models

type Album struct {
	Artist       string
	AlbumName    string
	RecordLabels []string
	ReleaseYear  string
	Genre        string
	Description  string
	TrackList    []Track
}
