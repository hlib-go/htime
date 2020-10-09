package t19

import (
	"testing"
	"time"
)

func TestTime_String(t *testing.T) {
	t.Log(time.Now().Add(time.Duration(7200 * 1e9)).Add(-5 * time.Minute).String())
}
