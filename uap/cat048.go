package uap

// Cat048V127 User Application Profile
// version 1.27
var Cat048V127 = StandardUAP{
	Name:     "CAT048",
	Category: 48,
	Version:  1.27,
	Items: []DataField{
		{
			FRN:      1,
			DataItem: "I048/010",
			Type: TypeField{
				Name: "Fixed",
				Size: 2,
			},
		},
		{
			FRN:      2,
			DataItem: "I048/140",
			Type: TypeField{
				Name: "Fixed",
				Size: 3,
			},
		},
		{
			FRN:      3,
			DataItem: "I048/020",
			Type: TypeField{
				Name: "Extended",
				Size: 1,
			},
		},
		{
			FRN:      4,
			DataItem: "I048/040",
			Type: TypeField{
				Name: "Fixed",
				Size: 4,
			},
		},
		{
			FRN: 5, DataItem: "I048/070", Type: TypeField{Name: "Fixed", Size: 2},
		},
		{
			FRN:      6,
			DataItem: "I048/090",
			Type: TypeField{
				Name: "Fixed",
				Size: 2,
			},
		},
		{
			FRN:      7,
			DataItem: "I048/130",
			Type: TypeField{
				Name: "Compound",
				Meta: MetaField{
					8: {Name: "Fixed", Size: 1},
					7: {Name: "Fixed", Size: 1},
					6: {Name: "Fixed", Size: 1},
					5: {Name: "Fixed", Size: 1},
					4: {Name: "Fixed", Size: 1},
					3: {Name: "Fixed", Size: 1},
					2: {Name: "Fixed", Size: 1},
				},
			},
		},
		{
			FRN: 8, DataItem: "I048/220", Type: TypeField{Name: "Fixed", Size: 3},
		},
		{
			FRN: 9, DataItem: "I048/240", Type: TypeField{Name: "Fixed", Size: 6},
		},
		{
			FRN:      10,
			DataItem: "I048/250",
			Type: TypeField{
				Name: "Repetitive",
				Size: 8,
			},
		},
		{
			FRN: 11, DataItem: "I048/161", Type: TypeField{Name: "Fixed", Size: 2},
		},
		{
			FRN: 12, DataItem: "I048/042", Type: TypeField{Name: "Fixed", Size: 4},
		},
		{
			FRN: 13, DataItem: "I048/200", Type: TypeField{Name: "Fixed", Size: 4},
		},
		{
			FRN: 14, DataItem: "I048/170", Type: TypeField{Name: "Extended", Size: 1},
		},
		{
			FRN: 15, DataItem: "I048/210", Type: TypeField{Name: "Fixed", Size: 4},
		},
		{
			FRN: 16, DataItem: "I048/030", Type: TypeField{Name: "Extended", Size: 1},
		},
		{
			FRN: 17, DataItem: "I048/080", Type: TypeField{Name: "Fixed", Size: 2},
		},
		{
			FRN: 18, DataItem: "I048/100", Type: TypeField{Name: "Fixed", Size: 4},
		},
		{
			FRN: 19, DataItem: "I048/110", Type: TypeField{Name: "Fixed", Size: 2},
		},
		{
			FRN:      20,
			DataItem: "I048/120",
			Type: TypeField{
				Name: "Compound",
				Meta: MetaField{
					8: {Name: "Fixed", Size: 2},
					7: {Name: "Repetitive", Size: 6},
					6: {Name: "Spare"},
					5: {Name: "Spare"},
					4: {Name: "Spare"},
					3: {Name: "Spare"},
					2: {Name: "Spare"},
				},
			},
		},
		{
			FRN: 21, DataItem: "I048/230", Type: TypeField{Name: "Fixed", Size: 2},
		},
		{
			FRN: 22, DataItem: "I048/260", Type: TypeField{Name: "Fixed", Size: 7},
		},
		{
			FRN: 23, DataItem: "I048/055", Type: TypeField{Name: "Fixed", Size: 1},
		},
		{
			FRN: 24, DataItem: "I048/050", Type: TypeField{Name: "Fixed", Size: 2},
		},
		{
			FRN: 25, DataItem: "I048/065", Type: TypeField{Name: "Fixed", Size: 1},
		},
		{
			FRN: 26, DataItem: "I048/060", Type: TypeField{Name: "Fixed", Size: 2},
		},
		{
			FRN: 27, DataItem: "SP-Data Item", Type: TypeField{Name: "SP"},
		},
		{
			FRN: 28, DataItem: "RE-Data Item", Type: TypeField{Name: "RE"},
		},
	},
}
