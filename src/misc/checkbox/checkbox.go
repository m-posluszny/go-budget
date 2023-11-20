package checkbox

import "bytes"

var (
	checkboxTrue = []byte(`"on"`)
)

type Checkbox bool

func (c Checkbox) Bool() bool {
	return bool(c)
}

func (c *Checkbox) UnmarshalForm(data []byte) error {
	*c = false
	if bytes.Equal(data, checkboxTrue) || bytes.Equal(data, []byte("true")) {
		*c = true
	}
	return nil
}
