package goasterix

import (
	"github.com/mokhtarimokhtar/goasterix/uap"
)

/*
<Fixed length="2">
	<Bits from="16" to="9">
		<BitsShortName>SAC</BitsShortName>
		<BitsName>System Area Code</BitsName>
	</Bits>
	<Bits from="8" to="1">
		<BitsShortName>SIC</BitsShortName>
		<BitsName>System Identification Code</BitsName>
	</Bits>
</Fixed>

<Bits bit="5">
	<BitsShortName>SIM</BitsShortName>
	<BitsValue val="0">Actual target report</BitsValue>
	<BitsValue val="1">Simulated target report</BitsValue>
</Bits>
<Bits bit="4">
	<BitsShortName>RDP</BitsShortName>
	<BitsValue val="0">Report from RDP Chain 1</BitsValue>
	<BitsValue val="1">Report from RDP Chain 2</BitsValue>
</Bits>
*/

// data 0xff, 0xff
// size field = 2 bytes
// (from=16, to=9) (from=8, to=1)

// data 0xff, 0xff
// size field = 2 bytes
// (bit=16, bit=15, bit=14) (from=12, to=1)

// data 0xff
// size field = 1 byte
// (bit=8, bit=7, bit=6) (from=5, to=1)

// data 0xff, 0xff [1][111-1111 1111]-[1][111]
// size field = 2 bytes
// (bit=16) (from=15, to=5) (bit=4) (from=3, to=1)

// err = binary.Read(rb, binary.BigEndian, &f.Data)

type SubItem interface {
	Reader(data []byte) error
	//GetType() uap.TypeField
}

type SubItemBit struct {
	Name string
	Type uap.TypeField
	Pos  uint8
	Data []byte
}

func newSubItemBit(field uap.IDataField) SubItem {
	f := &SubItemBit{}
	f.Type = field.GetType()
	return f
}

func (s *SubItemBit) Reader(data []byte) error {
	var err error
	nbBits := uint8(len(data)) * 8
	//littleEndianPos := nbBits - s.Pos
	index := (nbBits - s.Pos) / 8
	relativePos := s.Pos % 8

	//fmt.Println("s.Pos", s.Pos)
	//fmt.Println("nbBits", nbBits)
	////fmt.Println("littleEndianPos", littleEndianPos)
	//fmt.Println("index", index)
	//fmt.Println("relativePos", relativePos)

	/*
		data = 0x80 = 0[1]00-000
		pos = 7
		relativePos = 7 % 8 = 7
		nbBits = 8
		littleEndianPos = 8 - 7 = 1
		index = (8-7)/8 = 0


			[1111-1111 1111-1[1]11]
			nbBits = 16, pos = 3
			littleEndianPos = 16 - 3 = 13
			relativePos = 3%8= 3 pos%8
						relativeLittleEndianPos = 13 % 8 = 5
						nbBits - pos = 16 - 3 = 13, 13/8 = 1 , data[1]
						data[1]=[1111-1[1]11]

			[1111-1[1]11 1111-1111] length = 16, pos = 11
						littleEndianPos = 16 - 11 = 5, 11%8 = 3
						length - pos = 16 - 11 = 5, 5/8 = 0, data[0] 11-8=3
						data[0]=[1111-1[1]11]
	*/

	//s.Data = OneBitReader(data[index], relativePos)
	s.Data = make([]byte, 1)
	s.Data[0] = OneBitReader(data[index], relativePos)

	/*
		use case: 1 bit type
			1-use case: 1 bit + 1 byte
			get value of bit: OneBitReader
			return 1 byte = 0 or 1

			2-use case: 1 bit + 2 bytes = data[0] + data[1]
			get position of this bit inside []byte ?
				pos = 16 to 1
					if pos <= 8 = OneBitReader of data[1]
					if pos > 8 = OneBitReader of data[0]

			return 1 byte = 0 or 1

			ex: v = pos / 8
			[1111-1111 1111-1[1]11] length = 16, pos = 3
			littleEndianPos = 16 - 3 = 13, 3%8= 3
			relativeLittleEndianPos = 13 % 8 = 5
			length - pos = 16 - 3 = 13, 13/8 = 1 , data[1]
			data[1]=[1111-1[1]11]

			[1111-1[1]11 1111-1111] length = 16, pos = 11
			littleEndianPos = 16 - 11 = 5, 11%8 = 3
			length - pos = 16 - 11 = 5, 5/8 = 0, data[0] 11-8=3
			data[0]=[1111-1[1]11]

	*/
	return err
}

func OneBitReader(data byte, b uint8) byte {
	//func OneBitReader(data []byte, b uint8) []byte {
	//func OneBitReader(data byte, b uint8) []byte {
	pos := b - 1
	//tmp := make([]byte, 1)
	//tmp[0] = data >> pos & 0x01
	//return tmp
	return data >> pos & 0x01
}

type SubItemFromTo struct {
	Name string
	Type uap.TypeField
	From uint8
	To   uint8
	Data []byte
}

func (s *SubItemFromTo) Reader(data []byte) error {
	var err error

	return err
}

type SubItem2 struct {
	Name string
	Type uap.TypeField
	Pos  BitPosition
	Data []byte
}

type BitPosition struct {
	Bit  uint8
	From uint8
	To   uint8
}

func GetBitsFromTo(data []byte, from uint8, to uint8) []byte {
	//var d []byte
	length := uint8(len(data))
	//totalBits := uint8(8) * uint8(len(data))
	nbBits := from - to + 1
	sizeBytes := (nbBits + 7) / 8
	//fmt.Println(sizeBytes)
	d := make([]byte, sizeBytes)

	// ex
	// data 0xff, 0xff [1][111-1111 1111]-[1][111]
	// (from=15, to=5)
	// totalBits= 8*2 = 16
	// nbBits = 15-5+1 = 11
	// res= 0000-0111 1111-1111 = 07ff

	switch length {
	case 1:
		//
	case 2:
		//tmp := binary.BigEndian.Uint16(data)
		//fmt.Println(tmp)
		d = FromToBitReader16(data, from, to)
		//tmp = FromToBitReader16(tmp, from, to)
		//fmt.Printf("%x\n", tmp)
		//binary.BigEndian.PutUint16(d, tmp)
	case 3:
		//
	case 4:
		//
	}

	return d
}

func (s *SubItem2) Reader(data []byte) error {
	var err error

	//nbBits := s.Pos.From - s.Pos.To + 1 // 15-5+1 = 11 or 3-1+1 = 3
	//sizeBytes := (nbBits + 7) / 8       // (3+7)/8

	//tmp := make([]byte, 0, sizeBytes)
	//binary.BigEndian

	/*if s.Type == uap.Bit {
		//index := totalBits - s.Pos.Bit // 16 - 4 = 12
		//tmp := data[index]
		//fmt.Println(tmp)
	}*/
	if s.Type == uap.FromTo {
		s.Data = GetBitsFromTo(data, s.Pos.From, s.Pos.To)
	}

	//if s.Type == uap.FromTo {
	//	diff := s.Pos.From - s.Pos.To
	//
	//}
	return err
}
