package types

type YWCP struct {
	DDBH      string `json:"DDBH" form:"DDBH"`
	Startdate string `json:"Startdate" form:"Startdate"`
	Enddate   string `json:"Enddate" form:"Enddate"`
	Diffday   int    `json:"Diffday" form:"Diffday"`
}
