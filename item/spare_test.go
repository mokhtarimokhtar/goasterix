package item

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/util"
	"reflect"
	"testing"
)

func TestSpareClone(t *testing.T) {
	// Arrange
	input := Spare{
		Base: Base{
			FRN:  1,
			Type: SpareField,
		},
	}
	output := &Spare{
		Base: Base{
			FRN:  1,
			Type: SpareField,
		},
	}
	// Act
	res := input.Clone()

	// Assert
	if reflect.DeepEqual(res, output) == false {
		t.Errorf(util.MsgFailInValue, "", res, output)
	} else {
		t.Logf(util.MsgSuccessInValue, "", res, output)
	}

}

func TestSpareReader(t *testing.T) {
	// Arrange
	rb := bytes.NewReader([]byte{0xff})
	input := &Spare{
		Base{
			FRN:  1,
			Type: SpareField,
		},
	}
	// Act
	err := input.Reader(rb)

	// Assert
	if err != nil {
		t.Errorf(util.MsgFailInValue, "", err, nil)
	} else {
		t.Logf(util.MsgSuccessInValue, "", err, nil)
	}
}

func TestSpareString(t *testing.T) {
	// Arrange
	input := &Spare{
		Base{
			FRN:  1,
			Type: SpareField,
		},
	}
	output := ""
	// Act
	res := input.String()

	// Assert
	if res != output {
		t.Errorf(util.MsgFailInValue, "", res, output)
	} else {
		t.Logf(util.MsgSuccessInValue, "", res, output)
	}
}

func TestSparePayload(t *testing.T) {
	// Arrange
	input := &Spare{
		Base{
			FRN:  1,
			Type: SpareField,
		},
	}
	var output []byte
	// Act
	res := input.Payload()

	// Assert
	if bytes.Equal(res, output) == false {
		t.Errorf(util.MsgFailInValue, "", res, output)
	} else {
		t.Logf(util.MsgSuccessInValue, "", res, output)
	}
}