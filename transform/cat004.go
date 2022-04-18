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
	msgTypeCode007, msgTypeDesc007 string = "STCA", "short_term_conflict_alert"
	msgTypeCode008, msgTypeDesc008 string = "APM", "approach_path_monitor"
	msgTypeCode009, msgTypeDesc009 string = "ALM", "rimcas_arrival__Landing_monitor"
	msgTypeCode010, msgTypeDesc010 string = "WRA", "rimcas_arrival__departure_wrong_runway_alert"
	msgTypeCode011, msgTypeDesc011 string = "OTA", "rimcas_arrival__departure_opposite_traffic_alert"
	msgTypeCode012, msgTypeDesc012 string = "RDM", "rimcas_departure_monitor"
	msgTypeCode013, msgTypeDesc013 string = "RCM", "rimcas_runway__taxiway_crossing_monitor"
	msgTypeCode014, msgTypeDesc014 string = "TSM", "rimcas_taxiway_separation_monitor"
	msgTypeCode015, msgTypeDesc015 string = "UTMM", "rimcas_unauthorized_taxiway_movement_monitor"
	msgTypeCode016, msgTypeDesc016 string = "SBOA", "rimcas_stop_bar_overrun_alert"
	msgTypeCode017, msgTypeDesc017 string = "EOC", "end_of_conflict"
	msgTypeCode018, msgTypeDesc018 string = "ACASRA", "acas_resolution_advisory"
	msgTypeCode019, msgTypeDesc019 string = "NTCA", "near_term_conflict_alert"
	msgTypeCode020, msgTypeDesc020 string = "DBPSM", "downlinked_barometric_pressure_setting_monitor"
	msgTypeCode021, msgTypeDesc021 string = "SAM", "speed_adherence_monitor"
	msgTypeCode022, msgTypeDesc022 string = "OCAT", "outside_controlled_airspace_tool"
	msgTypeCode023, msgTypeDesc023 string = "VCD", "vertical_conflict_detection"
	msgTypeCode024, msgTypeDesc024 string = "VRAM", "vertical_rate_adherence_monitor"
	msgTypeCode025, msgTypeDesc025 string = "CHAM", "cleared_heading_adherence_monitor"
	msgTypeCode026, msgTypeDesc026 string = "DSAM", "downlinked_selected_altitude_monitor"
	msgTypeCode027, msgTypeDesc027 string = "HAM", "holding_adherence_monitor"
	msgTypeCode028, msgTypeDesc028 string = "VPM", "vertical_path_monitor"
	msgTypeCode029, msgTypeDesc029 string = "TTA", "rimcas_taxiway_traffic_alert"
	msgTypeCode030, msgTypeDesc030 string = "CRA", "rimcas_arrival__departure_close_runway_alert"
	msgTypeCode031, msgTypeDesc031 string = "ASM", "rimcas_arrival__departure_aircraft_separation_monitor"
	msgTypeCode032, msgTypeDesc032 string = "IAVM", "rimcas_ils_area_violation_monitor"
	msgTypeCode033, msgTypeDesc033 string = "FTD", "final_target_distance_indicator"
	msgTypeCode034, msgTypeDesc034 string = "ITD", "initial_target_distance_indicator"
	msgTypeCode035, msgTypeDesc035 string = "IIA", "wake_vortex_indicator_infringement_alert"
	msgTypeCode036, msgTypeDesc036 string = "SQW", "sequence_warning"
	msgTypeCode037, msgTypeDesc037 string = "CUW", "catch_up_warning"
	msgTypeCode038, msgTypeDesc038 string = "CATC", "conflicting_atc_clearances"
	msgTypeCode039, msgTypeDesc039 string = "NOCLR", "no_atc_clearance"
	msgTypeCode040, msgTypeDesc040 string = "NOMOV", "aircraft_not_moving_despite_atc_clearance"
	msgTypeCode041, msgTypeDesc041 string = "NOH", "aircraft_leaving__entering_the_aerodrome_area_without_proper_handover"
	msgTypeCode042, msgTypeDesc042 string = "WRTY", "wrong_runway_or_taxiway_type"
	msgTypeCode043, msgTypeDesc043 string = "STOCC", "stand_occupied"
	msgTypeCode044, msgTypeDesc044 string = "ONGOING", "ongoing_alert"
	msgTypeCode097, msgTypeDesc097 string = "LTW", "lost_track_warning"
	msgTypeCode098, msgTypeDesc098 string = "HVI", "holding_volume_infringement"
	msgTypeCode099, msgTypeDesc099 string = "AIW", "airspace_infringement_warning"
)

type MsgType struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

type Cat004Model struct {
	SacSic                   *SourceIdentifier         `json:"sourceIdentifier,omitempty"`
	MessageType              *MsgType                  `json:"messageType,omitempty"`
	SDPSIdentifier           []SourceIdentifier        `json:"sdpsIdentifier,omitempty"`
	TimeOfMessage            float64                   `json:"timeOfMessage,omitempty"`
	AlertIdentifier          uint16                    `json:"alertIdentifier"`
	AlertStatus              uint8                     `json:"alertStatus"`
	TrackNumberOne           uint16                    `json:"trackNumberOne,omitempty"`
	VerticalDeviation        int16                     `json:"verticalDeviation,omitempty"`
	AreaDefinition           *AreaDefinition           `json:"areaDefinition,omitempty"`
	TransversalDeviation     float32                   `json:"transversalDeviation,omitempty"`
	ConflictCharacteristics  *ConflictCharacteristics  `json:"conflictCharacteristics,omitempty"`
	ConflictTimingSeparation *ConflictTimingSeparation `json:"ConflictTimingSeparation,omitempty"`
	AircraftOne              *AircraftIdentification   `json:"aircraftOne,omitempty"`
	AircraftTwo              *AircraftIdentification   `json:"aircraftTwo,omitempty"`
	TrackNumberTwo           uint16                    `json:"trackNumberTwo,omitempty"`
}

// todo case 7
// todo case 13
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
			// I004/170, Aircraft Identification & Characteristics 1
			tmp := getAircraft(*item.Compound)
			data.AircraftOne = &tmp
		case 10:
			// I004/120, Conflict Characteristics
			data.ConflictCharacteristics = getConflictCharacteristics(*item.Compound)
		case 11:
			// I004/070, Conflict Timing and Separation
			data.ConflictTimingSeparation = getConflictTimingSeparation(*item.Compound)
		case 12:
			// I004/076, Vertical Deviation in ft, LSB = 25ft
			data.VerticalDeviation = (int16(item.Fixed.Data[0])<<8 + int16(item.Fixed.Data[1])) * 25
		case 14:
			// I004/075, Transversal Distance Deviation
			tmp := uint32(item.Fixed.Data[0])<<16 + uint32(item.Fixed.Data[1])<<8 + uint32(item.Fixed.Data[2])
			data.TransversalDeviation = float32(goasterix.TwoComplement32(24, tmp)) * 0.5
		case 15:
			// Data Item I004/100, Area Definition
			tmp := getAreaDefinition(*item.Compound)
			data.AreaDefinition = &tmp
		case 16:
			// I004/035 Track Number 2
			data.TrackNumberTwo = uint16(item.Fixed.Data[0])<<8 + uint16(item.Fixed.Data[1])
		case 17:
			// I004/171, Aircraft Identification & Characteristics 2
			tmp := getAircraft(*item.Compound)
			data.AircraftTwo = &tmp
		}
	}
}

type ConflictTimingSeparation struct {
	TimeToConflict              float64 `json:"timeToConflict,omitempty"`
	TimeToClosestApproach       float64 `json:"timeToClosestApproach,omitempty"`
	CurrentHorizontalSeparation float64 `json:"currentHorizontalSeparation,omitempty"`
	MinimumHorizontalSeparation float64 `json:"minimumHorizontalSeparation,omitempty"`
	CurrentVerticalSeparation   uint32  `json:"currentVerticalSeparation,omitempty"`
	MinimumVerticalSeparation   uint32  `json:"minimumVerticalSeparation,omitempty"`
}

func getConflictTimingSeparation(items goasterix.Compound) *ConflictTimingSeparation {
	var cts ConflictTimingSeparation
	for _, item := range items.Secondary {
		switch item.Meta.FRN {
		case 1:
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			cts.TimeToConflict, _ = timeOfDay(payload)
		case 2:
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			cts.TimeToClosestApproach, _ = timeOfDay(payload)
		case 3:
			cts.CurrentHorizontalSeparation = float64(uint32(item.Fixed.Data[0])<<16+uint32(item.Fixed.Data[1])<<8+uint32(item.Fixed.Data[2])) * 0.5
		case 4:
			cts.MinimumHorizontalSeparation = float64(uint16(item.Fixed.Data[0])<<8+uint16(item.Fixed.Data[1])) * 0.5
		case 5:
			cts.CurrentVerticalSeparation = uint32(uint16(item.Fixed.Data[0])<<8+uint16(item.Fixed.Data[1])) * 25
		case 6:
			cts.MinimumVerticalSeparation = uint32(uint16(item.Fixed.Data[0])<<8+uint16(item.Fixed.Data[1])) * 25
		}
	}

	return &cts
}

type ConflictCharacteristics struct {
	ConflictNature         *ConflictNature         `json:"conflictNature,omitempty"`
	ConflictClassification *ConflictClassification `json:"ConflictClassification,omitempty"`
	ConflictProbability    float32                 `json:"conflictProbability,omitempty"`
	ConflictDuration       float64                 `json:"conflictDuration,omitempty"`
}

type ConflictNature struct {
	MAS      string `json:"mas"`
	CAS      string `json:"cas"`
	FLD      string `json:"fld"`
	FVD      string `json:"fvd"`
	Type     string `json:"type"`
	Cross    string `json:"cross"`
	Div      string `json:"div"`
	RRC      string `json:"rrc,omitempty"`
	RTC      string `json:"rtc,omitempty"`
	MRVA     string `json:"mrva,omitempty"`
	VRAMCRM  string `json:"vramcrm,omitempty"`
	VRAMVRM  string `json:"vramvrm,omitempty"`
	VRAMVTM  string `json:"vramvtm,omitempty"`
	HAMHD    string `json:"hamhd,omitempty"`
	HAMRD    string `json:"hamrd,omitempty"`
	HAMVD    string `json:"hamvd,omitempty"`
	DBPSMARR string `json:"dbpsmarr,omitempty"`
	DBPSMDEP string `json:"dbpsmdep,omitempty"`
	DBPSMTL  string `json:"dbpsmtl,omitempty"`
	AIW      string `json:"aiw,omitempty"`
}

type ConflictClassification struct {
	TableId            uint8  `json:"tableId"`
	ConflictProperties uint8  `json:"conflictProperties"`
	CS                 string `json:"cs"`
}

func getConflictNature(item goasterix.Extended) *ConflictNature {
	var cn ConflictNature
	if item.Primary[0]&0x80 != 0 {
		cn.MAS = "conflict_predicted_to_occur_in_military_airspace"
	} else {
		cn.MAS = "conflict_not_predicted_to_occur_in_military_airspace"
	}
	if item.Primary[0]&0x40 != 0 {
		cn.CAS = "conflict_predicted_to_occur_in_civil_airspace"
	} else {
		cn.CAS = "conflict_not_predicted_to_occur_in_civil_airspace"
	}
	if item.Primary[0]&0x20 != 0 {
		cn.FLD = "aircraft_are_fast_diverging_laterally_at_current_time"
	} else {
		cn.FLD = "aircraft_are_not_fast_diverging_laterally_at_current_time"
	}
	if item.Primary[0]&0x10 != 0 {
		cn.FVD = "aircraft_are_fast_diverging_vertically_at_current_time"
	} else {
		cn.FVD = "aircraft_are_not_fast_diverging_vertically_at_current_time"
	}
	if item.Primary[0]&0x08 != 0 {
		cn.Type = "major_separation_infringement"
	} else {
		cn.Type = "minor_separation_infringement"
	}
	if item.Primary[0]&0x04 != 0 {
		cn.Cross = "aircraft_have_crossed_at_starting_time_of_conflict"
	} else {
		cn.Cross = "aircraft_have_not_crossed_at_starting_time_of_conflict"
	}
	if item.Primary[0]&0x02 != 0 {
		cn.Div = "aircraft_are_diverging_at_starting_time_of_conflict"
	} else {
		cn.Div = "aircraft_are_not_diverging_at_starting_time_of_conflict"
	}
	if item.Secondary != nil {
		if item.Secondary[0]&0x80 != 0 {
			cn.RRC = "runway_runway_crossing"
		} else {
			cn.RRC = "default"
		}
		if item.Secondary[0]&0x40 != 0 {
			cn.RTC = "runway_taxiway_crossing"
		} else {
			cn.RTC = "default"
		}
		if item.Secondary[0]&0x20 != 0 {
			cn.MRVA = "msg_type_4_indicates_mrva"
		} else {
			cn.MRVA = "default"
		}
		if item.Secondary[0]&0x10 != 0 {
			cn.VRAMCRM = "msg_type_25_indicates_crm"
		} else {
			cn.VRAMCRM = "default"
		}
		if item.Secondary[0]&0x08 != 0 {
			cn.VRAMVRM = "msg_type_25_indicates_vrm"
		} else {
			cn.VRAMVRM = "default"
		}
		if item.Secondary[0]&0x04 != 0 {
			cn.VRAMVTM = "msg_type_25_indicates_vtm"
		} else {
			cn.VRAMVTM = "default"
		}
		if item.Secondary[0]&0x02 != 0 {
			cn.HAMHD = "msg_type_29_indicates_hd"
		} else {
			cn.HAMHD = "default"
		}

		if item.Secondary[0]&0x01 != 0 {
			if item.Secondary[1]&0x80 != 0 {
				cn.HAMRD = "msg_type_29_indicates_rd"
			} else {
				cn.HAMRD = "default"
			}
			if item.Secondary[1]&0x40 != 0 {
				cn.HAMVD = "msg_type_29_indicates_vd"
			} else {
				cn.HAMVD = "default"
			}
			if item.Secondary[1]&0x20 != 0 {
				cn.DBPSMARR = "msg_type_20_indicates_arr"
			} else {
				cn.DBPSMARR = "default"
			}
			if item.Secondary[1]&0x10 != 0 {
				cn.DBPSMDEP = "msg_type_20_indicates_dep"
			} else {
				cn.DBPSMDEP = "default"
			}
			if item.Secondary[1]&0x08 != 0 {
				cn.DBPSMTL = "msg_type_20_indicates_above_tl"
			} else {
				cn.DBPSMTL = "default"
			}
			if item.Secondary[1]&0x04 != 0 {
				cn.AIW = "msg_type_99_indicates_paiw_alert"
			} else {
				cn.AIW = "default"
			}
		}
	}

	return &cn
}

func getConflictClassification(item byte) *ConflictClassification {
	var cc ConflictClassification
	cc.TableId = item & 0xf0 >> 4
	cc.ConflictProperties = item & 0x0e >> 1
	if item&0x01 != 0 {
		cc.CS = "high"
	} else {
		cc.CS = "low"
	}
	return &cc
}

func getConflictCharacteristics(items goasterix.Compound) *ConflictCharacteristics {
	var cc ConflictCharacteristics
	for _, item := range items.Secondary {
		switch item.Meta.FRN {
		case 1:
			cc.ConflictNature = getConflictNature(*item.Extended)
		case 2:
			var payload = item.Fixed.Data[0]
			cc.ConflictClassification = getConflictClassification(payload)
		case 3:
			cc.ConflictProbability = float32(item.Fixed.Data[0]) * 0.5
		case 4:
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			cc.ConflictDuration, _ = timeOfDay(payload)
		}
	}
	return &cc
}

type AreaDefinition struct {
	AreaName            string `json:"areaName,omitempty"`
	CrossingAreaName    string `json:"crossingAreaName,omitempty"`
	RunwayDesignatorOne string `json:"runwayDesignatorOne,omitempty"`
	RunwayDesignatorTwo string `json:"runwayDesignatorTwo,omitempty"`
	StopBarDesignator   string `json:"stopBarDesignator,omitempty"`
	GateDesignator      string `json:"gateDesignator,omitempty"`
}

func getAreaDefinition(items goasterix.Compound) AreaDefinition {
	var ad AreaDefinition
	for _, item := range items.Secondary {
		switch item.Meta.FRN {
		case 1:
			var payload [6]byte
			copy(payload[:], item.Fixed.Data[:])
			ad.AreaName, _ = modeSIdentification(payload)
		case 2:
			ad.CrossingAreaName = string(item.Fixed.Data)
		case 3:
			ad.RunwayDesignatorOne = string(item.Fixed.Data)
		case 4:
			ad.RunwayDesignatorTwo = string(item.Fixed.Data)
		case 5:
			ad.StopBarDesignator = string(item.Fixed.Data)
		case 6:
			ad.GateDesignator = string(item.Fixed.Data)
		}
	}
	return ad
}

// MessageType returns a struct of strings of message type.
// Ref. 6.2.1 Data Item I004/000, Message Type
func messageTypeCat004(data [1]byte) *MsgType {
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

	return &msg
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
