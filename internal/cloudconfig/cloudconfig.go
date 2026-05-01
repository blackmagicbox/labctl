package cloudconfig

type UserData struct {
	UserName string
	SSHKey   string
}

type MetaData struct {
	VMName   string
	Hostname string
}
