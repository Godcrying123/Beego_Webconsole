package models

type Machine struct {
	CPU       CPUDetail    `json:"CPU"`
	Memory    MemoryDetail `json:"Memory"`
	DiskSpace DiskDetail   `json:"DiskSpace"`
	OS        string       `json:"OS"`
	HostIp    []string
	HostName  string
}

type MachineSSH struct {
	NodeName string `json:"nodename"`
	HostIp   string `json:"hostip"`
	HostName string `json:"hostname"`
	UserName string `json:"username"`
	Password string `json:"password"`
	AuthType string `json:"authtype"`
	SSHPort  string `json:"sshport"`
	KeyFile  File
}

type CPUDetail struct {
	CPUModelandFrequency string
	CPUCores             int
	CPUPercentage        []float64
}

type MemoryDetail struct {
	TotalMemory      uint64
	UsedMemory       uint64
	MemoryPercentage float64
	SWAPonoff        bool
}

type DiskDetail struct {
	TotalDisk      uint64
	AvaileDisk     uint64
	UsedDisk       uint64
	DiskPercentage float64
}
