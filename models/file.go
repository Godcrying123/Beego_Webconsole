package models

import "time"

type Files struct {
	ID               int64
	FileName         string
	FileType         string
	FileLastModified time.Time
	FileOwnerShip    string
	FileContent      string
	FilePath         string
}

type DirListing struct {
	Name           string
	Children_dir   []string
	Children_files []string
}
