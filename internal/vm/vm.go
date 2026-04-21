package vm

import (
	"fmt"
	"reflect"
)

type Config struct {
	Distro   string
	Image    string
	VMName   string
	Username string
	Hostname string
	Disk     string
	Memory   string
	CPU      string
}

func NewVMConfig(distro, image, vmName, username, hostname, disk, memory, cpu string) *Config {
	return &Config{
		Distro:   distro,
		Image:    image,
		VMName:   vmName,
		Username: username,
		Hostname: hostname,
		Disk:     disk,
		Memory:   memory,
		CPU:      cpu,
	}
}

func (c *Config) String() string {
	out := "VM Config: \n"
	val := reflect.ValueOf(c).Elem()
	typeOfStruct := val.Type()

	for i := 0; i < val.NumField(); i++ {
		out = fmt.Sprintf("%s  %s: %v\n", out, typeOfStruct.Field(i).Name, val.Field(i).String())
	}
	return out
}
