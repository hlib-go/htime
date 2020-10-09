package rfc3339

import (
	"testing"
	"time"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	//2020-04-25T17:41:52+08:00
	t.Log(time.Now().Format(time.RFC3339))
	t.Log(time.Parse(time.RFC3339, "2006-01-02T15:04:05+08:00"))

}
