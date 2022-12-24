package jtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	JsonTimeFormat = "2006-01-02 15:04:05"
)

type JsonTime time.Time

func NewNowJsonTime() JsonTime {
	return JsonTime(time.Now())
}

func (t JsonTime) GetTime() time.Time {
	return time.Time(t)
}

func (t JsonTime) String() string {
	if t.GetTime().IsZero() {
		return ""
	}
	return t.GetTime().Format(JsonTimeFormat)
}

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return
	}
	now, err := time.ParseInLocation(`"`+JsonTimeFormat+`"`, string(data), time.Local)
	*t = JsonTime(now)
	return
}

func (t JsonTime) MarshalText() (text []byte, err error) {
	b := make([]byte, 0, len(JsonTimeFormat))
	//b = append(b, '"')
	b = time.Time(t).AppendFormat(b, JsonTimeFormat)
	//b = append(b, '"')
	if string(b) == `0001-01-01 00:00:00` {
		b = []byte(``)
	}
	return b, nil
}

func (t JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.GetTime().UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.GetTime(), nil
}

func (t *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
