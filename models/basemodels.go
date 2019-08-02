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

type Service struct {
	ID                     int64  `form:"-"`
	ServiceName            string `form:"service_name"`
	ServiceVersion         string `form:"service_version"`
	Status                 bool
	LastStatusModifiedTime time.Time
}

type UnitServices struct {
	Service               Service
	StatusCommand         string
	LogPath               string
	LastFiledModifiedTime time.Time
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

func init() {}
