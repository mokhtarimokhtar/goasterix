package goasterix

import (
	"github.com/mokhtarimokhtar/goasterix/item"
)

// ProfileRegistry contains the defaults User Application Profiles version.
var ProfileRegistry = map[uint8]item.StandardUAP{
	//1: Cat001V12,
	//2: Cat002V10,
	//4: Cat004V112,
	////21:  Cat021v10,
	//30:  Cat030StrV51,
	//32:  Cat032StrV70,
	//34:  Cat034V127,
	48: Cat048V127,
	//255: Cat255StrV51,
	//62:  Cat062V119,
	//63:  Cat063V16,
	//65:  Cat065V15,

	// Category for testing not exist
	26: CatForTest,
}
