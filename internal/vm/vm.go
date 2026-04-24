package vm

import "fmt"

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

func (c Config) String() string {
	return fmt.Sprintf(
		"VMConfig:\n"+
			"  Distro: %s\n"+
			"  Image: %s\n"+
			"  VMName: %s\n"+
			"  Username: %s\n"+
			"  Hostname: %s\n"+
			"  Disk: %s\n"+
			"  Memory: %s\n"+
			"  CPU: %s\n",
		c.Distro,
		c.Image,
		c.VMName,
		c.Username,
		c.Hostname,
		c.Disk,
		c.Memory,
		c.CPU,
	)
}
