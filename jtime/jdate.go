package jtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	JsonDateFormat = "2006-01-02"
)

type JsonDate time.Time

func NewNowJsonDate() JsonDate {
	return JsonDate(time.Now())
}

func (t JsonDate) GetTime() time.Time {
	return time.Time(t)
}

func (t JsonDate) String() string {
	if t.GetTime().IsZero() {
		return ""
	}
	return t.GetTime().Format(JsonDateFormat)
}

func (t *JsonDate) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+JsonDateFormat+`"`, string(data), time.Local)
	*t = JsonDate(now)
	return
}

func (t JsonDate) MarshalText() (text []byte, err error) {
	b := make([]byte, 0, len(JsonDateFormat))
	//b = append(b, '"')
	b = time.Time(t).AppendFormat(b, JsonDateFormat)
	//b = append(b, '"')
	if string(b) == `0001-01-01 00:00:00` {
		b = []byte(``)
	}
	return b, nil
}

func (t JsonDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.GetTime().UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.GetTime(), nil
}

func (t *JsonDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonDate(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
