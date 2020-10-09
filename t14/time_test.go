package t14

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTime_String(t *testing.T) {
	v := struct {
		T14 Time
	}{
		T14: Time(time.Now()),
	}

	r, e := json.Marshal(v)
	if e != nil {
		t.Error(e)
		return
	}
	t.Log(string(r))

	obj := struct {
		T14 Time
	}{}
	json.Unmarshal(r, &obj)

}
