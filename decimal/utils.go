package decimal

import (
	"errors"

	inf "gopkg.in/inf.v0"
)

func GetDecimal(s string) (d *inf.Dec, err error) {
	var success bool
	if d, success = d.SetString(s); !success {
		err = errors.New("Set string to decimal failed")
	}
	return
}
