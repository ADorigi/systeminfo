package main

import (
	"fmt"

	"github.com/adorigi/systeminfo"
)

func main() {
	info := systeminfo.CollectStats()

	fmt.Println("CPU name: " + info.CpuName)
	fmt.Println("CPU Architecture: ", info.CpuArch)
	fmt.Println("Operating System", info.OperatingSystem)
	fmt.Println("Available Disk Storage: ", info.DiskAvailable, info.StorageUnit)
	fmt.Println("Used Disk Storage: ", info.DiskUsed, info.StorageUnit)
	fmt.Println("Disk Used %: ", info.DiskUsedPercent, "%")
	fmt.Println("Local IPv4: ", info.LocalIPv4)
	fmt.Println("Global IP: ", info.GlobalIP)

}
