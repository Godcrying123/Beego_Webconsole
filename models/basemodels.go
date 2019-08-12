package models

import (
	"time"

	_ "github.com/astaxie/beego"
)

type Command struct {
	ID             int64
	Command        string
	CommandResults string
	ExcutedTime    time.Time
}

type Files struct {
	ID               int64
	FileName         string
	FileType         string
	FileLastModified time.Time
	FileOwnerShip    string
	FileContent      string
	FilePath         string
}

type Message struct {
	ID              int64
	MessageOwner    string
	MessageContent  string
	MessageSentTime time.Time
}
