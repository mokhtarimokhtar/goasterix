package uap

// Cat034V127 User Application Profile CAT034
// version 1.27
var Cat034V127 = StandardUAP{
	Category: 34,
	Version:  1.27,
	Items: []DataField{
		{
			FRN: 1, DataItem: "I034/010", Type: TypeField{NameType: Fixed, Size: 2},
		},
		{
			FRN: 2, DataItem: "I034/000", Type: TypeField{NameType: Fixed, Size: 1},
		},
		{
			FRN: 3, DataItem: "I034/030", Type: TypeField{NameType: Fixed, Size: 3},
		},
		{
			FRN: 4, DataItem: "I034/020", Type: TypeField{NameType: Fixed, Size: 1},
		},
		{
			FRN: 5, DataItem: "I034/041", Type: TypeField{NameType: Fixed, Size: 2},
		},
		{
			FRN: 6, DataItem: "I034/050",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {NameType: Fixed, Size: 1},
						7: {NameType: Spare},
						6: {NameType: Spare},
						5: {NameType: Fixed, Size: 1},
						4: {NameType: Fixed, Size: 1},
						3: {NameType: Fixed, Size: 2},
						2: {NameType: Spare},
					},
				},
			},
		},
		{
			FRN: 7, DataItem: "I034/060",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {NameType: Fixed, Size: 1},
						7: {NameType: Spare},
						6: {NameType: Spare},
						5: {NameType: Fixed, Size: 1},
						4: {NameType: Fixed, Size: 1},
						3: {NameType: Fixed, Size: 1},
						2: {NameType: Spare},
					},
				},
			},
		},
		{
			FRN: 8, DataItem: "I034/070", Type: TypeField{NameType: Repetitive, Size: 2},
		},
		{
			FRN: 9, DataItem: "I034/100", Type: TypeField{NameType: Fixed, Size: 8},
		},
		{
			FRN: 10, DataItem: "I034/110", Type: TypeField{NameType: Fixed, Size: 1},
		},
		{
			FRN: 11, DataItem: "I034/120", Type: TypeField{NameType: Fixed, Size: 8},
		},
		{
			FRN: 12, DataItem: "I034/090", Type: TypeField{NameType: Fixed, Size: 2},
		},
		{
			FRN: 13, DataItem: "RE-Data Item", Type: TypeField{NameType: RE},
		},
		{
			FRN: 14, DataItem: "SP-Data Item", Type: TypeField{NameType: SP},
		},
	},
}
