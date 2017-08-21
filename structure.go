package main

type Result struct {
	Host []Host `xml:"host"`
}

type Host struct {
	EndTime string `xml:"endtime,attr"`
	StartTime string `xml:"starttime,attr"`
	Status Status `xml:"status"`
	Address []Address `xml:"address"`
	Os Os `xml:"os"`
}

type Status struct {
	Reason string `xml:"reason,attr"`
	Reason_TTL string `xml:"reason_ttl,attr"`
	State string `xml:"state,attr"`
}

type Address struct {
	Addr string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
}

type Os struct {
	OsMatch []OsMatch `xml:"osmatch"`
}

type OsMatch struct {
	Accuracy string `xml:"accuracy,attr"`
	Line string `xml:"line,attr"`
	Name string `xml:"name,attr"`
}