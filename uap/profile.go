package uap

// DefaultProfiles contains the defaults User Application Profiles version.
var DefaultProfiles = map[uint8]StandardUAP{
	1:   Cat001V12,
	2:   Cat002V10,
	21:  Cat021v10,
	30:  Cat030StrV51,
	32:  Cat032StrV70,
	34:  Cat034V127,
	48:  Cat048V127,
	255: Cat255StrV51,
	62:  Cat062V119,
	63:  Cat063V16,
	// Category for testing not exist
	26: Cat4Test,
}
