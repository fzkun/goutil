package magicapi

type ApiHook interface {
	RecordFunc(record RecordData) error //记录api
}

type RecordData struct {
	Url      string `json:"url" gorm:"column:url;type:text;"`
	Method   string `json:"method"` //请求方式
	Body     string `json:"body" gorm:"column:body;type:longtext;"`
	Response string `json:"response" gorm:"column:response;type:longtext;"`
	Error    string `json:"error" gorm:"column:error;type:text;"`
	Status   int    `json:"status" gorm:"column:status;"` //http code
	Duration int64  `json:"duration"`                     //请求时长 ms
}
