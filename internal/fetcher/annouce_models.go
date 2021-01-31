package fetcher

type SHAnnouce struct { //
	PageHelp   PageHelp `json:"pageHelp"`
	ReportType string   `json:"reportType"`
	Results    []Result `json:"result"`
}

type PageHelp struct { //
	Total     uint16 `json:"total"`
	PageCount uint16 `json:"pageCount"`
	PageNo    uint16 `json:"pageNo"`
	PageSize  uint16 `json:"pageSize"`
}

type Result struct {
	ADDDATE       string `json:"ADDDATE"`
	BULLETIN_TYPE string `json:"BULLETIN_TYPE"`
	SECURITY_CODE string `json:"SECURITY_CODE"`
	TITLE         string `json:"TITLE"`
}
