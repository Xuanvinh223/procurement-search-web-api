package types

type YWCP struct {
	DDBH      string `json:"DDBH" form:"DDBH"`
	StartDate string `json:"StartDate" form:"StartDate"`
	EndDate   string `json:"EndDate" form:"EndDate"`
	DiffDay   int    `json:"DiffDay" form:"DiffDay"`
}
