package item

type FieldReferenceNumber uint8

const (
	FRN1 FieldReferenceNumber = iota + 1
	FRN2
	FRN3
	FRN4
	FRN5
	FRN6
	FRN7
	FRN8
	FRN9
	FRN10
	FRN11
	FRN12
	FRN13
	FRN14
	FRN15
	FRN16
	FRN17
	FRN18
	FRN19
	FRN20
	FRN21
	FRN22
	FRN23
	FRN24
	FRN25
	FRN26
	FRN27
	FRN28
	FRN29
	FRN30
	FRN31
	FRN32
	FRN33
	FRN34
	FRN35
)

// UAP is a User Application Profile ASTERIX
type UAP struct {
	Name      string
	Category  uint8
	Version   float32
	DataItems []DataItem
}

type Base struct {
	FRN          FieldReferenceNumber
	DataItemName string
	Description  string
	Type         TypeField
}

// GetFrn returns FieldReferenceNumber number of dataField from UAP
func (b Base) GetFrn() FieldReferenceNumber {
	return b.FRN
}
func (b Base) GetType() TypeField {
	return b.Type
}

func (b Base) GetDataItemName() string {
	return b.DataItemName
}
func (b Base) GetDescription() string {
	return b.Description
}
