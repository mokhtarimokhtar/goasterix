package uap

// Cat255StrV51 User Application Profile
// version 5.1
// French ANSP specific category
var Cat255StrV51 = StandardUAP{
	Category: 255,
	Version:  5.1,
	Items: []DataField{
		{
			FRN: 1, DataItem: "I255/010", Type: TypeField{Name: Fixed, Size: 2},
		},
		{
			FRN: 2, DataItem: "I255/020", Type: TypeField{Name: Fixed, Size: 3},
		},
		{
			FRN: 3, DataItem: "I255/030", Type: TypeField{Name: Extended, Size: 1},
		},
		{
			FRN: 4, DataItem: "I255/050", Type: TypeField{Name: Fixed, Size: 4},
		},
		{
			FRN:      5,
			DataItem: "I255/060",
			Type: TypeField{
				Name: Repetitive,
				Size: 1,
			},
		},
		{
			FRN:      6,
			DataItem: "I255/070",
			Type: TypeField{
				Name: Fixed,
				Size: 9,
			},
		},
		{
			FRN:      7,
			DataItem: "I255/040",
			Type: TypeField{
				Name: Repetitive,
				Size: 10,
			},
		},
	},
}
