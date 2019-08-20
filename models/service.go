package models

import "time"

type Service struct {
	ID             int64  `form:"-"`
	ServiceName    string `form:"service_name"`
	ServiceVersion string `form:"service_version"`
	ActiveStatus   string
	RunningStatus  string
	ServiceStatus  string
}

type UnitServices struct {
	StatusCommand          []Command
	File                   []File
	LastStatusModifiedTime time.Time
	LastFiledModifiedTime  time.Time
}
