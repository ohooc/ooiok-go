package options_partten

import "testing"

func TestOptions(t *testing.T) {
	client := NewHttpClient(WithTimeout(10), WithMaxIdle(30), WithCallback(func(err error) {
		t.Logf(err.Error())
	}))
	t.Logf("%+v", client)
}
