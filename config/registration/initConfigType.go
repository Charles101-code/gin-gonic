package registration

import "time"

type OwnerInfo struct {
	Name  string
	Email string
	DOB   time.Time
}

type Database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type Server struct {
	IP   string
	DC   string
	Port string
}

type Clients struct {
	Data  [][]interface{}
	Hosts []string
}
