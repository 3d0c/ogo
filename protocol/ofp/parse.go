package ofp

import (
	"github.com/3d0c/ogo/protocol/ofp10"
	"github.com/3d0c/ogo/protocol/ofp13"
	"github.com/3d0c/ogo/protocol/util"
)

func Parse(b []byte) (message util.Message, err error) {
	switch b[0] {
	case 1:
		message, err = ofp10.Parse(b)
	case 4:
		message, err = ofp13.Parse(b)
	}
	return
}
