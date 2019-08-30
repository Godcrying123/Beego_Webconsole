package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"webconsole_sma/models"

	"github.com/bitly/go-simplejson"
	"github.com/mitchellh/mapstructure"
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

func HostInfoRead() (hostoutput models.Machine, err error) {
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

func HostJsonGenerator() (message string, err error) {
	var builder strings.Builder
	for _, hostentity := range SSHHosts {
		output, err := json.MarshalIndent(hostentity, "", "\t")
		if err != nil {
			beego.Error(err)
			return "", err
		}
		builder.Write(output)
		builder.WriteString(",\n")
	}
	err = WriteJson(builder, "json/all_host_infos.json")
	return "all hosts info have been exported to JSON successfully!", nil
}

func HostSave(nodenames, hostips, hostnames, users, passwords, authtypes, sshports, keyfiles []string) (err error) {
	for index := 0; index < len(nodenames); index++ {
		sshHost := models.MachineSSH{
			NodeName: strings.Trim(nodenames[index], " "),
			HostName: strings.Trim(hostnames[index], " "),
			HostIp:   strings.Trim(hostips[index], " "),
			UserName: strings.Trim(users[index], " "),
			Password: strings.Trim(passwords[index], " "),
			SSHPort:  strings.Trim(sshports[index], " "),
			AuthType: strings.Trim(authtypes[index], " "),
		}
		if sshports[index] == "" {
			sshHost.SSHPort = "22"
		}
		SSHHosts[sshHost.NodeName] = sshHost
	}
	return
}

func HostJsonRead(filePath string) (jsonStruct map[string]models.MachineSSH, err error) {
	var byter bytes.Buffer
	jsonFile, err := ioutil.ReadFile(filePath)
	jsonStruct = make(map[string]models.MachineSSH)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	byter.Write([]byte("["))
	byter.Write(jsonFile)
	byter.Write([]byte("]"))
	jsons, _ := simplejson.NewJson(byter.Bytes())
	for _, jsonmap := range jsons.MustArray() {
		machine := models.MachineSSH{}
		err = mapstructure.WeakDecode(jsonmap.(map[string]interface{}), &machine)
		if err != nil {
			beego.Error(err)
		}
		jsonStruct[machine.NodeName] = machine
	}
	return jsonStruct, nil
}
