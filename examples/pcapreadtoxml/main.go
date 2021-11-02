package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"goasterix"
	"goasterix/model"
	"log"
)

var (
	pcapFile = "../data/sample.pcap"
	handle   *pcap.Handle
	err      error
	llcLayer layers.LLC
	ethLayer layers.Ethernet
)

func main() {
	// Open file instead of device
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	indexPacket := 0

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		m := packet.Metadata()
		capLen := m.CaptureLength
		dt := m.Timestamp

		indexPacket++

		parser := gopacket.NewDecodingLayerParser(
			layers.LayerTypeEthernet,
			&ethLayer,
			&llcLayer,
		)
		var foundLayerTypes []gopacket.LayerType

		err := parser.DecodeLayers(packet.Data(), &foundLayerTypes)
		if err != nil {
			fmt.Println("Trouble decoding layers: ", err)
		}
		loop:
		for _, layerType := range foundLayerTypes {
			if layerType == layers.LayerTypeLLC {
				data := llcLayer.Payload // it contains only payload of LLC
				w := new(goasterix.WrapperDataBlock) // wrapper of asterix datablock, it contains one or more datablocks
				_, err = w.Decode(data) // decode method one or more datablocks
				if err != nil {
					fmt.Printf("Trouble decoding wrapper packet: %s, Frame n°%v, Time = %v\n", err, indexPacket, dt)
					break loop
				}
				fmt.Printf("\n# Frame n°%v, Size = %v bytes, Time = %v\n", indexPacket, capLen, dt)

				for _, dataB := range w.DataBlocks {
					fmt.Printf("Category: %v, Len: %v\n", dataB.Category, dataB.Len)
					// Parsing JSON datablock for each record

					if dataB.Category == 48 {
						for _, record := range dataB.Records {
							catModel := new(model.Cat048Model)
							catXml, _ := model.WriteModelXML(catModel, record.Items)
							fmt.Println(string(catXml))
						}
					} else if dataB.Category == 34 {
						for _, record := range dataB.Records {
							catModel := new(model.Cat034Model)
							catXml, _ := model.WriteModelXML(catModel, record.Items)
							fmt.Println(string(catXml))
						}
					} else if dataB.Category == 30 {
						for _, record := range dataB.Records {
							catModel := new(model.Cat030STRModel)
							catXml, _ := model.WriteModelXML(catModel, record.Items)
							fmt.Println(string(catXml))
						}
					} else if dataB.Category == 255 {
						for _, record := range dataB.Records {
							catModel := new(model.Cat255STRModel)
							catXml, _ := model.WriteModelXML(catModel, record.Items)
							fmt.Println(string(catXml))
						}
					}
				}
			}
		}
	}

}



