package uap

// Cat002V10 User Application Profile CAT002
// version 1.0
var Cat002V10 = StandardUAP{
	Category: 2,
	Version:  1.0,
	Items: []DataField{
		{
			FRN: 1, DataItem: "I002/010", Type: TypeField{Name: "Fixed", Size: 2},
		},
		{
			FRN: 2, DataItem: "I002/000", Type: TypeField{Name: "Fixed", Size: 1},
		},
		{
			FRN: 3, DataItem: "I002/020", Type: TypeField{Name: "Fixed", Size: 1},
		},
		{
			FRN: 4, DataItem: "I002/030", Type: TypeField{Name: "Fixed", Size: 3},
		},
		{
			FRN: 5, DataItem: "I002/041", Type: TypeField{Name: "Fixed", Size: 2},
		},
		{
			FRN: 6, DataItem: "I002/050", Type: TypeField{Name: "Extended", Size: 1},
		},
		{
			FRN: 7, DataItem: "I002/060", Type: TypeField{Name: "Extended", Size: 1},
		},
		{
			FRN: 8, DataItem: "I002/070", Type: TypeField{Name: "Repetitive", Size: 2},
		},
		{
			FRN: 9, DataItem: "I002/100", Type: TypeField{Name: "Fixed", Size: 8},
		},
		{
			FRN: 10, DataItem: "I002/090", Type: TypeField{Name: "Fixed", Size: 2},
		},
		{
			FRN: 11, DataItem: "I002/080", Type: TypeField{Name: "Extended", Size: 1},
		},
		{
			FRN: 12, DataItem: "NA", Type: TypeField{Name: "NA"},
		},
		{
			FRN: 13, DataItem: "SP-Data Item", Type: TypeField{Name: "SP"},
		},
		{
			FRN: 14, DataItem: "Random Field Sequencing", Type: TypeField{Name: "RFS"},
		},
	},
}
