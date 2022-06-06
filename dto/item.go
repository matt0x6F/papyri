package dto

import (
	"encoding/json"

	"github.com/google/uuid"
)

const (
	Directory = iota
	Bookmark
)

type Source struct {
	ID   uuid.UUID
	Name string // Chrome, Firefox, GitHub, HackerNews
}

type Item struct {
	ID       uuid.UUID
	Title    string
	Author   string
	Parent   uuid.UUID
	Metadata json.RawMessage
	Type     int // Directory, Bookmark
}
