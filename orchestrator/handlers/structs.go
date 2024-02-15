package handlers 

// for all request structs to be jsoned

// /chtimeout

type ChtimeReqIn struct{
	Sign string `json:"sign"`
	Ms int `json:"ms"`
}

type ChtimeReqOut struct{
	Errmsg string `json:"errmsg"`
}

// /timeouts

type TimeoutsReqOut struct{
	Plus ""
}