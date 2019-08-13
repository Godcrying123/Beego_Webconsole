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

type Message struct {
	ID              int64
	MessageOwner    string
	MessageContent  string
	MessageSentTime time.Time
}
