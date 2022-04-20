/*
Package goasterix provides an ASTERIX (All Purpose Structured EUROCONTROL Surveillance Information Exchange)
packet decoding and marshalling JSON/XML for the Go language.
datablock.go and record.go contain the logic built into goasterix for decoding packet datagram.

goasterix contains some sub-packages including:
 * dataField: This contains all definition ASTERIX Profile (User Application Profile).
 * transform: This contains of the logic marshalling in JSON or XML format.
 * commbds: It is for decoding Comm-B Data Selector of transponder (ICAO Doc 9871:Technical Provisions for Mode S
	Services and Extended Squitter)

Data Block corresponds at general message structure.
It contains:
  * a one-octet dataField Data Category (CAT) indicating to which Category the data transmitted belongs;
  * two-octet dataField Length Indicator (LEN) indicating the total length (in octets) of the Data Block, including
  	the CAT and LEN fields;
  * one or more Record(s) containing data of the same Category.

Data Record contains:
  * a Field Specification (FSPEC) dataField of variable length, considered as a table of contents, in the form of a bit
	sequence, where every individual bit signals the presence (bit set to one) or absence (bit set to zero) of a
	well-defined Data Field assigned to it;
  * a variable number of Data Fields. Each Data Field is associated with one and only one Data Item, as defined by the
	UAP.

Describe an UAP definition: todo
Basic Usage: todo

*/
package goasterix
