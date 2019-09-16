package models

import (
	"os"
	"time"
)

type File struct {
	FileName         string
	FileType         string
	FileAccess       os.FileMode
	FileContent      string
	FilePath         string
	FileSize         int64
	FileLastModified time.Time
}

type Directory struct {
	DirName         string
	DirAccess       os.FileMode
	DirSize         int64
	DirPath         string
	DirLastModified time.Time
}

type DirListing struct {
	Name          string
	ChildrenDirs  []Directory
	ChildrenFiles []File
}
