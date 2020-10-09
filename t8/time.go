package t8

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	YYYYMMDD = "20060102"
)

func NowF8() string {
	return Time(time.Now()).String()
}

type Time time.Time

func (nt Time) String() string {
	return time.Time(nt).Local().Format(YYYYMMDD)
}

func (nt Time) MarshalJSON() ([]byte, error) {
	t := time.Time(nt).Local()

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(YYYYMMDD)+2)
	b = append(b, '"')
	if !t.IsZero() {
		b = t.AppendFormat(b, YYYYMMDD)
	}
	b = append(b, '"')
	return b, nil
}

func (nt *Time) UnmarshalJSON(data []byte) error {
	v := strings.Trim(string(data), "\"")
	if data == nil || v == "" {
		*nt = Time(time.Time{})
		return nil
	}
	t, err := time.Parse(YYYYMMDD, v)
	if err != nil {
		return fmt.Errorf("解析时间字符串'%s'出错", v)
	}
	*nt = Time(t.Local())
	return err
}
