package types

type Entity struct {
	Name string
	Parent *string
	Limit int
	Utilised int
}

type Report struct {
	Name string
	Entries []string
	Allocation int
	DirectUsage int
	Usage int
	SubTotalLimit int
	SubReports []*Report
}