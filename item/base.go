package item

type StandardUAP struct {
	Name      string
	Category  uint8
	Version   float64
	DataItems []DataItem
}

type Base struct {
	FRN          uint8
	DataItemName string
	Description  string
	Type         TypeField
}

func (b *Base) NewBase(field DataItem) {
	b.FRN = field.GetFrn()
	b.DataItemName = field.GetDataItemName()
	b.Description = field.GetDescription()
	b.Type = field.GetType()
}

// GetFrn returns FRN number of dataField from UAP
func (b Base) GetFrn() uint8 {
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
