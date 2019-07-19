package models

import (
	"time"

	_ "github.com/astaxie/beego"
)

type Machine struct {
	ID         int64
	CPU        string
	CPUCores   int
	Memory     int64
	DiskSpace  int64
	SWAPStatus string
	FireFall   string
}

type Services struct {
	ID               int64
	ServiceName      string
	ServiceVersion   string
	Status           string
	StatusCommand    string
	LogPath          string
	LastModifiedTime time.Time
}

type Steps struct {
	ID              int64
	StepSummary     string
	StepDescription string
	StepResult      string
}

type Message struct {
	ID              int64
	MessageOwner    string
	MessageContent  string
	MessageSentTime time.Time
}

type Files struct {
	ID               int64
	FileName         string
	FileType         string
	FileLastModified time.Time
	FileOwnerShip    string
	FileContent      string
}
