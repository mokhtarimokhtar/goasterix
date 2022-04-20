package uap

type TypeField uint8

const (
	Fixed TypeField = iota + 1
	Extended
	Compound
	Repetitive
	Explicit
	SP
	RE
	RFS
	Spare
)

// StandardUAP is User Application Profile
// Cat is ASTERIX Category number (integer)
// Version is ASTERIX version for a category
type StandardUAP struct {
	Name     string
	Category uint8
	Version  float64
	Items    []DataField
}

type IDataField interface {
	GetFrn() uint8
	GetDataItem() string
	GetDescription() string
	GetType() TypeField
	GetSize() Size
	GetCompound() []DataField
	GetRFS() []DataField
}

// DataField describes FRN(Field Reference Number)
type DataField struct {
	FRN         uint8
	DataItem    string
	Description string
	Type        TypeField
	Compound    []DataField
	RFS         []DataField
	Conditional bool
	Size        Size
}

func DataFieldFactory(frn uint8, dataItem string, desc string, t TypeField, s Size) IDataField {
	var d = &DataField{}
	d.FRN = frn
	d.DataItem = dataItem
	d.Description = desc
	d.Type = t
	d.Size = s
	return d
}

func (d DataField) GetFrn() uint8 {
	return d.FRN
}
func (d DataField) GetDataItem() string {
	return d.DataItem
}
func (d DataField) GetDescription() string {
	return d.Description
}
func (d DataField) GetType() TypeField {
	return d.Type
}
func (d DataField) GetSize() Size {
	return d.Size
}
func (d DataField) GetCompound() []DataField {
	return d.Compound
}
func (d DataField) GetRFS() []DataField {
	return d.RFS
}

type Size struct {
	ForFixed             uint8
	ForExtendedPrimary   uint8
	ForExtendedSecondary uint8
	ForRepetitive        uint8
}
