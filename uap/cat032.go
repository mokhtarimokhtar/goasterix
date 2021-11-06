package uap

// Cat032StrV70 User Application Profile
// version 5.1
var Cat032StrV70 = StandardUAP{
	Category: 32,
	Version:  7.0,
	Items: []DataField{
		{
			FRN: 1, DataItem: "I032/010", Type: TypeField{Name: Fixed, Size: 2},
		},
		{
			FRN: 2, DataItem: "I032/020", Type: TypeField{Name: Fixed, Size: 3},
		},
		{
			FRN: 3, DataItem: "I032/060", Type: TypeField{Name: Fixed, Size: 4},
		},
		{
			FRN: 4, DataItem: "I032/070", Type: TypeField{Name: Fixed, Size: 15},
		},
		{
			FRN: 5, DataItem: "I032/080", Type: TypeField{Name: Fixed, Size: 12}, //todo: add field extended
		},
	},
}
