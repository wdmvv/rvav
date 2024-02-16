package handlers 

// for all request structs to be jsoned

// /chtime

type ChtimeReqIn struct{
	Sign string `json:"sign"`
	Ms int `json:"ms"`
}

type ChtimeReqOut struct{
	Error string `json:"error"`
}

// /status

type StatusReqOut struct{
	Msg string `json:"msg"`
}