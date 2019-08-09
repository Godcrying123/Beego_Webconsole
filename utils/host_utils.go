package utils

import (
	"strings"
	"webconsole_sma/models"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"

	"github.com/astaxie/beego"

	"github.com/shirou/gopsutil/mem"
)

func SSHConnectHost() {

}

func beegoerr(err error) {
	if err != nil {
		beego.Error(err)
	}
}

func HostInfoRead(hostinput models.Machine) (hostoutput models.Machine, err error) {
	var builder strings.Builder
	vmStat, err := mem.VirtualMemory()
	beegoerr(err)
	diskStat, err := disk.Usage("/")
	beegoerr(err)
	cpuStat, err := cpu.Info()
	beegoerr(err)
	percentage, err := cpu.Percent(0, true)
	beegoerr(err)
	hostStat, err := host.Info()
	beegoerr(err)
	interfStat, err := net.Interfaces()
	beegoerr(err)
	builder.WriteString(hostStat.OS + " ")
	builder.WriteString(hostStat.Platform + " ")
	builder.WriteString(hostStat.PlatformVersion + " ")
	builder.WriteString(hostStat.KernelVersion)
	hostoutput.CPU.CPUModelandFrequency = cpuStat[0].ModelName
	hostoutput.CPU.CPUCores = len(cpuStat)
	for _, cpupercentage := range percentage {
		hostoutput.CPU.CPUPercentage = append(hostoutput.CPU.CPUPercentage, cpupercentage)
	}
	hostoutput.Memory.TotalMemory = vmStat.Total
	hostoutput.Memory.UsedMemory = vmStat.Used
	hostoutput.Memory.MemoryPercentage = vmStat.UsedPercent
	hostoutput.Memory.SWAPonoff = (vmStat.SwapTotal != 0)
	hostoutput.DiskSpace.TotalDisk = diskStat.Total
	hostoutput.DiskSpace.UsedDisk = diskStat.Used
	hostoutput.DiskSpace.AvaileDisk = diskStat.Free
	hostoutput.DiskSpace.DiskPercentage = diskStat.UsedPercent
	hostoutput.HostName = hostStat.Hostname
	hostoutput.OS = builder.String()
	for _, interf := range interfStat {
		for _, addr := range interf.Addrs {
			hostoutput.HostIp = append(hostoutput.HostIp, addr.String())
		}
	}
	return hostoutput, nil
}
