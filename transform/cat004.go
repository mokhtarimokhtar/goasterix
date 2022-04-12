package transform

import (
	"github.com/mokhtarimokhtar/goasterix"
	"math"
	"strconv"
)

const (
	msgTypeCode000, msgTypeDesc000 string = "UNDEFINED", "undefined_message_type"
	msgTypeCode001, msgTypeDesc001 string = "AM", "alive_message"
	msgTypeCode002, msgTypeDesc002 string = "RAMLD", "route_adherence_monitor_longitudinal_deviation"
	msgTypeCode003, msgTypeDesc003 string = "RAMHD", "route_adherence_monitor_heading_deviation"
	msgTypeCode004, msgTypeDesc004 string = "MSAW", "minimum_safe_altitude_warning"
	msgTypeCode005, msgTypeDesc005 string = "APW", "area_proximity_warning"
	msgTypeCode006, msgTypeDesc006 string = "CLAM", "clearance_level_adherence_monitor"
	msgTypeCode007, msgTypeDesc007 string = "STCA", "Short Term Conflict Alert ()"
	msgTypeCode008, msgTypeDesc008 string = "APM", "approach_path_monitor"
	msgTypeCode009, msgTypeDesc009 string = "ALM", "RIMCAS Arrival / Landing Monitor ()"
	msgTypeCode010, msgTypeDesc010 string = "WRA", "RIMCAS Arrival / Departure Wrong Runway Alert ()"
	msgTypeCode011, msgTypeDesc011 string = "OTA", "RIMCAS Arrival / Departure Opposite Traffic Alert ()"
	msgTypeCode012, msgTypeDesc012 string = "RDM", "RIMCAS Departure Monitor ()"
	msgTypeCode013, msgTypeDesc013 string = "RCM", "RIMCAS Runway / Taxiway Crossing Monitor ()"
	msgTypeCode014, msgTypeDesc014 string = "TSM", "RIMCAS Taxiway Separation Monitor ()"
	msgTypeCode015, msgTypeDesc015 string = "UTMM", "RIMCAS Unauthorized Taxiway Movement Monitor()"
	msgTypeCode016, msgTypeDesc016 string = "SBOA", "RIMCAS Stop Bar Overrun Alert ()"
	msgTypeCode017, msgTypeDesc017 string = "EOC", "End Of Conflict ()"
	msgTypeCode018, msgTypeDesc018 string = "ACASRA", "ACAS Resolution Advisory ()"
	msgTypeCode019, msgTypeDesc019 string = "NTCA", "Near Term Conflict Alert ()"
	msgTypeCode020, msgTypeDesc020 string = "DBPSM", "Downlinked Barometric Pressure Setting Monitor ()"
	msgTypeCode021, msgTypeDesc021 string = "SAM", "Speed Adherence Monitor ()"
	msgTypeCode022, msgTypeDesc022 string = "OCAT", "Outside Controlled Airspace Tool ()"
	msgTypeCode023, msgTypeDesc023 string = "VCD", "Vertical Conflict Detection ()"
	msgTypeCode024, msgTypeDesc024 string = "VRAM", "Vertical Rate Adherence Monitor ()"
	msgTypeCode025, msgTypeDesc025 string = "CHAM", "Cleared Heading Adherence Monitor ()"
	msgTypeCode026, msgTypeDesc026 string = "DSAM", "Downlinked Selected Altitude Monitor ()"
	msgTypeCode027, msgTypeDesc027 string = "HAM", "Holding Adherence Monitor ()"
	msgTypeCode028, msgTypeDesc028 string = "VPM", "Vertical Path Monitor ()"
	msgTypeCode029, msgTypeDesc029 string = "TTA", "RIMCAS Taxiway Traffic Alert ()"
	msgTypeCode030, msgTypeDesc030 string = "CRA", "RIMCAS Arrival/Departure Close Runway Alert ()"
	msgTypeCode031, msgTypeDesc031 string = "ASM", "RIMCAS Arrival/Departure Aircraft Separation Monitor()"
	msgTypeCode032, msgTypeDesc032 string = "IAVM", "RIMCAS ILS Area Violation Monitor ()"
	msgTypeCode033, msgTypeDesc033 string = "FTD", "Final Target Distance Indicator ()"
	msgTypeCode034, msgTypeDesc034 string = "ITD", "Initial Target Distance Indicator ()"
	msgTypeCode035, msgTypeDesc035 string = "IIA", "Wake Vortex Indicator Infringement Alert ()"
	msgTypeCode036, msgTypeDesc036 string = "SQW", "Sequence Warning ()"
	msgTypeCode037, msgTypeDesc037 string = "CUW", "Catch Up Warning ()"
	msgTypeCode038, msgTypeDesc038 string = "CATC", "Conflicting ATC Clearances ()"
	msgTypeCode039, msgTypeDesc039 string = "NOCLR", "No ATC Clearance ()"
	msgTypeCode040, msgTypeDesc040 string = "NOMOV", "Aircraft Not Moving despite ATC Clearance ()"
	msgTypeCode041, msgTypeDesc041 string = "NOH", "Aircraft leaving/entering the aerodrome area without proper handover ()"
	msgTypeCode042, msgTypeDesc042 string = "WRTY", "Wrong Runway or Taxiway Type ()"
	msgTypeCode043, msgTypeDesc043 string = "STOCC", "Stand Occupied ()"
	msgTypeCode044, msgTypeDesc044 string = "ONGOING", "Ongoing Alert ()"
	msgTypeCode097, msgTypeDesc097 string = "LTW", "Lost Track Warning ()"
	msgTypeCode098, msgTypeDesc098 string = "HVI", "Holding Volume Infringement ()"
	msgTypeCode099, msgTypeDesc099 string = "AIW", "Airspace Infringement Warning ()"
)

type MsgType struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

type Cat004Model struct {
	SacSic               *SourceIdentifier       `json:"sourceIdentifier,omitempty"`
	MessageType          MsgType                 `json:"messageType,omitempty"`
	SDPSIdentifier       []SourceIdentifier      `json:"sdpsIdentifier,omitempty"`
	TimeOfMessage        float64                 `json:"timeOfDay,omitempty"`
	AlertIdentifier      uint16                  `json:"alertIdentifier"`
	AlertStatus          uint8                   `json:"alertStatus"`
	TrackNumberOne       uint16                  `json:"trackNumberOne"`
	VerticalDeviation    int16                   `json:"verticalDeviation,omitempty"`
	TransversalDeviation float32                 `json:"transversalDeviation,omitempty"`
	AircraftOne          *AircraftIdentification `json:"aircraftOne,omitempty"`
}

// todo case 7
// todo case 10
// todo case 11
// todo case 13
// todo case 15
// todo case 16
// todo case 17
// todo case 18
// todo case 20
// todo case 21
func (data *Cat004Model) write(rec goasterix.Record) {
	for _, item := range rec.Items {
		switch item.Meta.FRN {
		case 1:
			// decode sac sic
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp, _ := sacSic(payload)
			data.SacSic = &tmp
		case 2:
			//decode messageTypeCat004
			var payload [1]byte
			copy(payload[:], item.Fixed.Data[:])
			data.MessageType = messageTypeCat004(payload)
		case 3:
			// Data Item I004/015 SDPS Identifier
			data.SDPSIdentifier = getSDPSIdentifier(*item.Repetitive)
		case 4:
			// decode timeOfMessage
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfMessage, _ = timeOfDay(payload)
		case 5:
			// I004/040 Alert Identifier
			data.AlertIdentifier = uint16(item.Fixed.Data[0])<<8 + uint16(item.Fixed.Data[1])
		case 6:
			// I004/045 Alert Status
			data.AlertStatus = item.Fixed.Data[0] & 0x0E >> 1
		case 8:
			// I004/030 Track Number 1
			data.TrackNumberOne = uint16(item.Fixed.Data[0])<<8 + uint16(item.Fixed.Data[1])
		case 9:
			// Aircraft Derived Data
			tmp := getAircraft(*item.Compound)
			data.AircraftOne = &tmp

		case 12:
			// I004/076, Vertical Deviation in ft, LSB = 25ft
			data.VerticalDeviation = (int16(item.Fixed.Data[0])<<8 + int16(item.Fixed.Data[1])) * 25
		case 14:
			// I004/075, Transversal Distance Deviation
			tmp := uint32(item.Fixed.Data[0])<<16 + uint32(item.Fixed.Data[1])<<8 + uint32(item.Fixed.Data[2])
			data.TransversalDeviation = float32(goasterix.TwoComplement32(24, tmp)) * 0.5
		}
	}
}

// MessageType returns a struct of strings of message type.
// Ref. 6.2.1 Data Item I004/000, Message Type
func messageTypeCat004(data [1]byte) MsgType {
	var msg MsgType
	msgType := data[0]
	switch msgType {
	case 1:
		msg.Code = msgTypeCode001
		msg.Desc = msgTypeDesc001
	case 2:
		msg.Code = msgTypeCode002
		msg.Desc = msgTypeDesc002
	case 3:
		msg.Code = msgTypeCode003
		msg.Desc = msgTypeDesc003
	case 4:
		msg.Code = msgTypeCode004
		msg.Desc = msgTypeDesc004
	case 5:
		msg.Code = msgTypeCode005
		msg.Desc = msgTypeDesc005
	case 6:
		msg.Code = msgTypeCode006
		msg.Desc = msgTypeDesc006
	case 7:
		msg.Code = msgTypeCode007
		msg.Desc = msgTypeDesc007
	case 8:
		msg.Code = msgTypeCode008
		msg.Desc = msgTypeDesc008
	case 9:
		msg.Code = msgTypeCode009
		msg.Desc = msgTypeDesc009
	case 10:
		msg.Code = msgTypeCode010
		msg.Desc = msgTypeDesc010
	case 11:
		msg.Code = msgTypeCode011
		msg.Desc = msgTypeDesc011
	case 12:
		msg.Code = msgTypeCode012
		msg.Desc = msgTypeDesc012
	case 13:
		msg.Code = msgTypeCode013
		msg.Desc = msgTypeDesc013
	case 14:
		msg.Code = msgTypeCode014
		msg.Desc = msgTypeDesc014
	case 15:
		msg.Code = msgTypeCode015
		msg.Desc = msgTypeDesc015
	case 16:
		msg.Code = msgTypeCode016
		msg.Desc = msgTypeDesc016
	case 17:
		msg.Code = msgTypeCode017
		msg.Desc = msgTypeDesc017
	case 18:
		msg.Code = msgTypeCode018
		msg.Desc = msgTypeDesc018
	case 19:
		msg.Code = msgTypeCode019
		msg.Desc = msgTypeDesc019
	case 20:
		msg.Code = msgTypeCode020
		msg.Desc = msgTypeDesc020
	case 21:
		msg.Code = msgTypeCode021
		msg.Desc = msgTypeDesc021
	case 22:
		msg.Code = msgTypeCode022
		msg.Desc = msgTypeDesc022
	case 23:
		msg.Code = msgTypeCode023
		msg.Desc = msgTypeDesc023
	case 24:
		msg.Code = msgTypeCode024
		msg.Desc = msgTypeDesc024
	case 25:
		msg.Code = msgTypeCode025
		msg.Desc = msgTypeDesc025
	case 26:
		msg.Code = msgTypeCode026
		msg.Desc = msgTypeDesc026
	case 27:
		msg.Code = msgTypeCode027
		msg.Desc = msgTypeDesc027
	case 28:
		msg.Code = msgTypeCode028
		msg.Desc = msgTypeDesc028
	case 29:
		msg.Code = msgTypeCode029
		msg.Desc = msgTypeDesc029
	case 30:
		msg.Code = msgTypeCode030
		msg.Desc = msgTypeDesc030
	case 31:
		msg.Code = msgTypeCode031
		msg.Desc = msgTypeDesc031
	case 32:
		msg.Code = msgTypeCode032
		msg.Desc = msgTypeDesc032
	case 33:
		msg.Code = msgTypeCode033
		msg.Desc = msgTypeDesc033
	case 34:
		msg.Code = msgTypeCode034
		msg.Desc = msgTypeDesc034
	case 35:
		msg.Code = msgTypeCode035
		msg.Desc = msgTypeDesc035
	case 36:
		msg.Code = msgTypeCode036
		msg.Desc = msgTypeDesc036
	case 37:
		msg.Code = msgTypeCode037
		msg.Desc = msgTypeDesc037
	case 38:
		msg.Code = msgTypeCode038
		msg.Desc = msgTypeDesc038
	case 39:
		msg.Code = msgTypeCode039
		msg.Desc = msgTypeDesc039
	case 40:
		msg.Code = msgTypeCode040
		msg.Desc = msgTypeDesc040
	case 41:
		msg.Code = msgTypeCode041
		msg.Desc = msgTypeDesc041
	case 42:
		msg.Code = msgTypeCode042
		msg.Desc = msgTypeDesc042
	case 43:
		msg.Code = msgTypeCode043
		msg.Desc = msgTypeDesc043
	case 44:
		msg.Code = msgTypeCode044
		msg.Desc = msgTypeDesc044
	case 97:
		msg.Code = msgTypeCode097
		msg.Desc = msgTypeDesc097
	case 98:
		msg.Code = msgTypeCode098
		msg.Desc = msgTypeDesc098
	case 99:
		msg.Code = msgTypeCode099
		msg.Desc = msgTypeDesc099

	default:
		msg.Code = msgTypeCode000
		msg.Desc = msgTypeDesc000
	}

	return msg
}

func getSDPSIdentifier(item goasterix.Repetitive) []SourceIdentifier {
	var sdps []SourceIdentifier
	data := item.Data
	for i := 0; i < int(item.Rep*2); i = i + 2 {
		var payload [2]byte
		copy(payload[:], data[i:i+2])
		tmp, _ := sacSic(payload)
		sdps = append(sdps, tmp)
	}
	return sdps
}

// AircraftIdentification
// Identification & Characteristics of Aircraft 1 Involved in the Conflict.
type AircraftIdentification struct {
	AircraftIdentifier                 string                     `json:"aircraftIdentifier,omitempty"`
	Mode3ACodeAircraft                 string                     `json:"mode3ACodeAircraft,omitempty"`
	PredictedConflictPositionWGS84     *ConflictPositionWGS84     `json:"predictedConflictPosition,omitempty"`
	PredictedConflictPositionCartesian *ConflictPositionCartesian `json:"predictedConflictPositionCartesian,omitempty"`
	TimeToThreshold                    float64                    `json:"timeToThreshold,omitempty"`
	DistanceToThreshold                float64                    `json:"DistanceToThreshold,omitempty"`
	ModeSIdentifier                    string                     `json:"modeSIdentifier,omitempty"`
	FlightPlanNumber                   uint32                     `json:"flightPlanNumber,omitempty"`
	ClearedFlightLevel                 float64                    `json:"clearedFlightLevel,omitempty"`
	AircraftCharacteristics            *Characteristics           `json:"aircraftCharacteristics,omitempty"`
}

type ConflictPositionWGS84 struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  int32   `json:"altitude"`
}

type ConflictPositionCartesian struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z int32   `json:"z"`
}

// Data Item I004/170, Aircraft Identification & Characteristics 1
func getAircraft(items goasterix.Compound) AircraftIdentification {
	var ai AircraftIdentification
	for _, item := range items.Secondary {
		switch item.Meta.FRN {
		case 1:
			ai.AircraftIdentifier = string(item.Fixed.Data)
		case 2:
			tmp := uint16(item.Fixed.Data[0])&0x000F<<8 + uint16(item.Fixed.Data[1])&0x00FF
			ai.Mode3ACodeAircraft = strconv.FormatUint(uint64(tmp), 8)
		case 3:
			var pos = new(ConflictPositionWGS84)
			lsb := 180 / math.Pow(2, 25)

			pos.Latitude = float64(int32(item.Fixed.Data[0])<<24+int32(item.Fixed.Data[1])<<16+int32(item.Fixed.Data[2])<<8+int32(item.Fixed.Data[3])) * lsb
			pos.Longitude = float64(int32(item.Fixed.Data[4])<<24+int32(item.Fixed.Data[5])<<16+int32(item.Fixed.Data[6])<<8+int32(item.Fixed.Data[7])) * lsb
			pos.Altitude = int32(int16(item.Fixed.Data[8])<<8+int16(item.Fixed.Data[9])) * 25

			ai.PredictedConflictPositionWGS84 = pos
		case 4:
			var pos = new(ConflictPositionCartesian)
			tmpX := uint32(item.Fixed.Data[0])<<16 + uint32(item.Fixed.Data[1])<<8 + uint32(item.Fixed.Data[2])
			pos.X = float64(goasterix.TwoComplement32(24, tmpX)) * 0.5

			tmpY := uint32(item.Fixed.Data[3])<<16 + uint32(item.Fixed.Data[4])<<8 + uint32(item.Fixed.Data[5])
			pos.Y = float64(goasterix.TwoComplement32(24, tmpY)) * 0.5

			pos.Z = int32(int16(item.Fixed.Data[6])<<8+int16(item.Fixed.Data[7])) * 25
			ai.PredictedConflictPositionCartesian = pos
		case 5:
			tmp := uint32(item.Fixed.Data[0])<<16 + uint32(item.Fixed.Data[1])<<8 + uint32(item.Fixed.Data[2])
			ai.TimeToThreshold = float64(goasterix.TwoComplement32(24, tmp)) / 128
		case 6:
			ai.DistanceToThreshold = float64(uint16(item.Fixed.Data[0])<<8+uint16(item.Fixed.Data[1])) * 0.5
		case 7:
			ai.AircraftCharacteristics = getCharacteristics(*item.Extended)
		case 8:
			var payload [6]byte
			copy(payload[:], item.Fixed.Data[:])
			ai.ModeSIdentifier, _ = modeSIdentification(payload)
		case 9:
			ai.FlightPlanNumber = uint32(item.Fixed.Data[0])<<24 + uint32(item.Fixed.Data[1])<<16 + uint32(item.Fixed.Data[2])<<8 + uint32(item.Fixed.Data[3])
		case 10:
			ai.ClearedFlightLevel = float64(int16(item.Fixed.Data[0])<<8+int16(item.Fixed.Data[1])) * 0.25
		}
	}
	return ai
}

type Characteristics struct {
	AT   string `json:"at"`
	FR   string `json:"fr"`
	RVSM string `json:"rvsm"`
	HPR  string `json:"hpr"`
	CDM  string `json:"cdm,omitempty"`
	PRI  string `json:"pri,omitempty"`
	GV   string `json:"gv,omitempty"`
}

func getCharacteristics(item goasterix.Extended) *Characteristics {
	var cha = new(Characteristics)
	tmp := item.Primary[0] & 0xc0 >> 6
	switch tmp {
	case 0:
		cha.AT = "unknown"
	case 1:
		cha.AT = "general_air_traffic"
	case 2:
		cha.AT = "operational_air_traffic"
	case 3:
		cha.AT = "not_applicable"
	}

	tmp = item.Primary[0] & 0x30 >> 4
	switch tmp {
	case 0:
		cha.FR = "instrument_flight_rules"
	case 1:
		cha.FR = "visual_flight_rules"
	case 2:
		cha.FR = "not_applicable"
	case 3:
		cha.FR = "controlled_visual_flight_rules"
	}

	tmp = item.Primary[0] & 0x0c >> 2
	switch tmp {
	case 0:
		cha.RVSM = "unknown"
	case 1:
		cha.RVSM = "approved"
	case 2:
		cha.RVSM = "exempt"
	case 3:
		cha.RVSM = "not_approved"
	}

	if item.Primary[0]&0x02 != 0 {
		cha.HPR = "high_priority_flight"
	} else {
		cha.HPR = "normal_priority_flight"
	}
	if item.Secondary != nil {
		tmp = item.Secondary[0] & 0xc0 >> 6
		switch tmp {
		case 0:
			cha.CDM = "maintaining"
		case 1:
			cha.CDM = "climbing"
		case 2:
			cha.CDM = "descending"
		case 3:
			cha.CDM = "invalid"
		}
		if item.Secondary[0]&0x20 != 0 {
			cha.PRI = "primary_target"
		} else {
			cha.PRI = "non_primary_target"
		}
		if item.Secondary[0]&0x10 != 0 {
			cha.GV = "ground_vehicle"
		} else {
			cha.GV = "default"
		}

	}
	return cha
}
