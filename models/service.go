package models

import "time"

type Service struct {
	ID             int64  `form:"-"`
	ServiceName    string `form:"service_name"`
	ServiceVersion string `form:"service_version"`
	ActiveStatus   bool
	RunningStatus  bool
}

type UnitServices struct {
	Service                Service
	StatusCommand          []Command
	File                   []Files
	LastStatusModifiedTime time.Time
	LastFiledModifiedTime  time.Time
}
