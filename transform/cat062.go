package transform

import (
	"encoding/hex"
	"math"
	"strconv"
	"strings"

	"github.com/mokhtarimokhtar/goasterix"
)

type TrackVelocity struct {
	Vx float32 `json:"vx,omitempty"`
	Vy float32 `json:"vy,omitempty"`
}
type Acceleration struct {
	Ax float32 `json:"ax,omitempty"`
	Ay float32 `json:"ay,omitempty"`
}
type PositionWGS84 struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type TrackMode3A struct {
	V      string `json:"v"`
	G      string `json:"g"`
	CH     string `json:"ch"`
	Squawk string `json:"squawk"`
}
type TargetIdent struct {
	Target string `json:"target,omitempty"`
	STI    string `json:"sti,omitempty"`
}
type BarometricAltitude struct {
	QNH      string  `json:"qnh,omitempty"`
	Altitude float64 `json:"altitude,omitempty"`
}

type IAS struct {
	IM       string  `json:"im"`
	AirSpeed float64 `json:"airSpeed"`
}

type SelectedAltitude struct {
	SAS      string  `json:"sas"`
	Source   string  `json:"source"`
	Altitude float64 `json:"altitude"`
}

type StateSelectedAltitude struct {
	MV       string  `json:"mv"`
	AH       string  `json:"ah"`
	AM       string  `json:"am"`
	Altitude float64 `json:"altitude"`
}

type DerivedData struct {
	TargetAddress         string                 `json:"targetAddress,omitempty"`
	TargetIdentification  string                 `json:"targetIdentification,omitempty"`
	MagneticHeading       float64                `json:"magneticHeading,omitempty"`
	IndicatedAirspeedOld  *IAS                   `json:"indicatedAirspeedOld,omitempty"`
	AirSpeed              uint16                 `json:"airSpeed,omitempty"`
	SelectedAltitude      *SelectedAltitude      `json:"selectedAltitude,omitempty"`
	StateSelectedAltitude *StateSelectedAltitude `json:"stateSelectedAltitude,omitempty"`
	MachNumber            float64                `json:"machNumber,omitempty"`
	IndicatedAirSpeed     float64                `json:"indicatedAirSpeed,omitempty"`
}

type ModeMov struct {
	TRANS string `json:"trans"`
	LONG  string `json:"long"`
	VERT  string `json:"vert"`
	ADF   string `json:"adf"`
}

type Cat062Model struct {
	SacSic                *SourceIdentifier    `json:"sourceIdentifier,omitempty"`
	ServiceIdentification uint8                `json:"serviceIdentification,omitempty"`
	TimeOfDay             float64              `json:"timeOfDay,omitempty"`
	TrackPositionWGS84    *PositionWGS84       `json:"trackPositionWGS84"`
	CartesianXY           *CartesianXYPosition `json:"cartesianXY,omitempty"`
	TrackVelocity         *TrackVelocity       `json:"trackVelocity,omitempty"`
	Acceleration          *Acceleration        `json:"acceleration,omitempty"`
	Mode3ACode            *TrackMode3A         `json:"mode3ACode,omitempty"`
	TargetIdentification  *TargetIdent         `json:"targetIdentification,omitempty"`
	AircraftDerivedData   *DerivedData         `json:"aircraftDerivedData,omitempty"`
	TrackNumber           uint16               `json:"trackNumber,omitempty"`
	TrackStatus           *TrackStatus         `json:"trackStatus,omitempty"`
	ModeOfMovement        *ModeMov             `json:"modeOfmovement,omitempty"`
	FlightLevel           float32              `json:"flightLevel,omitempty"`
	GeometricAltitude     float32              `json:"geometricAltitude,omitempty"`
	BarometricAltitude    *BarometricAltitude  `json:"barometricAltitude,omitempty"`
	RateOfClimbDescent    float32              `json:"rateOfClimbDescent,omitempty"`
}

// todo case 14
// todo case 16
// todo case 21
// todo case 22
// todo case 23
// todo case 24
// todo case 25
// todo case 26
// todo case 27
// todo case 28
// todo case 34
// todo case 35
// Write writes a single ASTERIX Record to Cat062Model.
// CompoundItems is a slice of CompoundItems DataField.
func (data *Cat062Model) write(rec goasterix.Record) {
	for _, item := range rec.Items {
		switch item.Meta.FRN {
		case 1:
			// decode sac sic
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp, _ := sacSic(payload)
			data.SacSic = &tmp
		// case 2 is spare
		case 3:
			// Service Identification
			data.ServiceIdentification = item.Fixed.Data[0]
		case 4:
			// Time Of Track Information
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfDay, _ = timeOfDay(payload)
		case 5:
			// Calculated Track Position (WGS-84)
			var payload [8]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := calculatedTrackPositionWGS84(payload)
			data.TrackPositionWGS84 = &tmp
		case 6:
			// Calculated Track Position. (Cartesian)
			var payload [6]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := calculatedTrackPositionCartesian(payload)
			data.CartesianXY = &tmp
		case 7:
			// Calculated Track Velocity (Cartesian)
			var payload [4]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := calculatedTrackVelocityCartesian(payload)
			data.TrackVelocity = &tmp
		case 8:
			// Calculated Acceleration (Cartesian)
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := calculatedAccelerationCartesian(payload)
			data.Acceleration = &tmp
		case 9:
			// Track Mode 3/A Code
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := mode3ACode(payload)
			data.Mode3ACode = &tmp
		case 10:
			// Target Identification
			var payload [7]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := targetIdentification(payload)
			data.TargetIdentification = &tmp
		case 11:
			// Aircraft Derived Data
			tmp := extractDerivedData(*item.Compound)
			data.AircraftDerivedData = &tmp
		case 12:
			// Track Number
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TrackNumber = trackNumber(payload)
		case 13:
			// Track Status
			tmp := extractTrackStatus(*item.Extended)
			data.TrackStatus = &tmp
		case 15:
			// Mode of Movement
			var payload [1]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := extractModeOfMovement(payload)
			data.ModeOfMovement = &tmp
		case 17:
			// Measured Flight Level
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.FlightLevel = measuredFlightLevel(payload)
		case 18:
			// Calculated Track Geometric Altitude
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.GeometricAltitude = trackGeometricAltitude(payload)
		case 19:
			// Calculated Track Barometric Altitude
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := trackBarometricAltitude(payload)
			data.BarometricAltitude = &tmp
		case 20:
			// Calculated Rate Of Climb/Descent
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.RateOfClimbDescent = rateOfClimbDescent(payload)

		}
	}
}

// extractModeOfMovement returns Calculated Mode of Movement of a target
func extractModeOfMovement(data [1]byte) ModeMov {
	var mov ModeMov
	tmp := data[0] & 0xc0 >> 6
	switch tmp {
	case 0:
		mov.TRANS = "constant_course"
	case 1:
		mov.TRANS = "right_turn"
	case 2:
		mov.TRANS = "left_turn"
	case 3:
		mov.TRANS = "undetermined"
	}

	tmp = data[0] & 0x30 >> 4
	switch tmp {
	case 0:
		mov.LONG = "constant_groundspeed"
	case 1:
		mov.LONG = "increasing_groundspeed"
	case 2:
		mov.LONG = "decreasing_groundspeed"
	case 3:
		mov.LONG = "undetermined"
	}

	tmp = data[0] & 0x0c >> 2
	switch tmp {
	case 0:
		mov.VERT = "level"
	case 1:
		mov.VERT = "climb"
	case 2:
		mov.VERT = "descent"
	case 3:
		mov.VERT = "undetermined"
	}

	if data[0]&0x02 != 0 {
		mov.ADF = "altitude_discrepancy"
	} else {
		mov.ADF = "no_altitude_discrepancy"
	}

	return mov
}

type TrackStatus struct {
	MON string `json:"mon"`
	SPI string `json:"spi"`
	MRH string `json:"mrh"`
	SRC string `json:"src"`
	CNF string `json:"cnf"`
	TrackStatusFirstExtent
	TrackStatusSecondExtent
	TrackStatusThirdExtent
	TrackStatusFourthExtent
	TrackStatusFifthExtent
}
type TrackStatusFirstExtent struct {
	SIM string `json:"sim,omitempty"`
	TSE string `json:"tse,omitempty"`
	TSB string `json:"tsb,omitempty"`
	FPC string `json:"fpc,omitempty"`
	AFF string `json:"aff,omitempty"`
	STP string `json:"stp,omitempty"`
	KOS string `json:"kos,omitempty"`
}
type TrackStatusSecondExtent struct {
	AMA string `json:"ama,omitempty"`
	MD4 string `json:"md4,omitempty"`
	ME  string `json:"me,omitempty"`
	MI  string `json:"mi,omitempty"`
	MD5 string `json:"md5,omitempty"`
}
type TrackStatusThirdExtent struct {
	CST string `json:"cst,omitempty"`
	PSR string `json:"psr,omitempty"`
	SSR string `json:"ssr,omitempty"`
	MDS string `json:"mds,omitempty"`
	ADS string `json:"ads,omitempty"`
	SUC string `json:"suc,omitempty"`
	AAC string `json:"aac,omitempty"`
}
type TrackStatusFourthExtent struct {
	SDS  string `json:"sds,omitempty"`
	EMS  string `json:"ems,omitempty"`
	PFT  string `json:"pft,omitempty"`
	FPLT string `json:"fplt,omitempty"`
}
type TrackStatusFifthExtent struct {
	DUPT string `json:"dupt,omitempty"`
	DUPF string `json:"dupf,omitempty"`
	DUPM string `json:"dupm,omitempty"`
	SFC  string `json:"sfc,omitempty"`
	IDD  string `json:"idd,omitempty"`
	IEC  string `json:"iec,omitempty"`
}

// extractTrackStatus returns Status of a track.
func extractTrackStatus(item goasterix.Extended) TrackStatus {
	var ts TrackStatus

	if item.Primary[0]&0x80 != 0 {
		ts.MON = "multisensor"
	} else {
		ts.MON = "monosensor"
	}
	if item.Primary[0]&0x40 != 0 {
		ts.SPI = "last_report_received"
	} else {
		ts.SPI = "default_value"
	}
	if item.Primary[0]&0x20 != 0 {
		ts.MRH = "geometric_altitude_reliable"
	} else {
		ts.MRH = "barometric_altitude_reliable"
	}
	tmp := item.Primary[0] & 0x1c >> 2
	switch tmp {
	case 0:
		ts.SRC = "no_source"
	case 1:
		ts.SRC = "gnss"
	case 2:
		ts.SRC = "3d_radar"
	case 3:
		ts.SRC = "triangulation"
	case 4:
		ts.SRC = "height_coverage"
	case 5:
		ts.SRC = "speed_look_up_table"
	case 6:
		ts.SRC = "default_height"
	case 7:
		ts.SRC = "multilateration"
	}

	if item.Primary[0]&0x02 != 0 {
		ts.CNF = "tentative_track"
	} else {
		ts.CNF = "confirmed_track"
	}

	if item.Secondary != nil {
		if item.Secondary[0]&0x80 != 0 {
			ts.SIM = "simulated_track"
		} else {
			ts.SIM = "actual_track"
		}
		if item.Secondary[0]&0x40 != 0 {
			ts.TSE = "last_message_transmitted"
		} else {
			ts.TSE = "default_value"
		}
		if item.Secondary[0]&0x20 != 0 {
			ts.TSB = "first_message_transmitted"
		} else {
			ts.TSB = "default_value"
		}
		if item.Secondary[0]&0x10 != 0 {
			ts.FPC = "flight_plan_correlated"
		} else {
			ts.FPC = "not_flight_plan_correlated"
		}
		if item.Secondary[0]&0x08 != 0 {
			ts.AFF = "ads_b_data_inconsistent"
		} else {
			ts.AFF = "default_value"
		}
		if item.Secondary[0]&0x04 != 0 {
			ts.STP = "slave_track_promotion"
		} else {
			ts.STP = "default_value"
		}
		if item.Secondary[0]&0x02 != 0 {
			ts.KOS = "background_service_used"
		} else {
			ts.KOS = "complementary_service_used"
		}

		if item.Secondary[0]&0x01 != 0 {
			if item.Secondary[1]&0x80 != 0 {
				ts.AMA = "track_resulting_amalgamation_process"
			} else {
				ts.AMA = "track_not_resulting_amalgamation_process"
			}
			tmp := item.Secondary[1] & 0x60 >> 5
			switch tmp {
			case 0:
				ts.MD4 = "no_mode_4_interrogation"
			case 1:
				ts.MD4 = "friendly_target"
			case 2:
				ts.MD4 = "unknown_target"
			case 3:
				ts.MD4 = "no_reply"
			}
			if item.Secondary[1]&0x10 != 0 {
				ts.ME = "military_emergency_last_report_received"
			} else {
				ts.ME = "default_value"
			}
			if item.Secondary[1]&0x08 != 0 {
				ts.MI = "military_identification_last_report_received"
			} else {
				ts.MI = "default_value"
			}
			tmp = item.Secondary[1] & 0x06 >> 1
			switch tmp {
			case 0:
				ts.MD5 = "no_mode_5_interrogation"
			case 1:
				ts.MD5 = "friendly_target"
			case 2:
				ts.MD5 = "unknown_target"
			case 3:
				ts.MD5 = "no_reply"
			}

			if item.Secondary[1]&0x01 != 0 {
				if item.Secondary[2]&0x80 != 0 {
					ts.CST = "age_last_track_higher_than_system_dependent_threshold"
				} else {
					ts.CST = "default_value"
				}
				if item.Secondary[2]&0x40 != 0 {
					ts.PSR = "age_last_psr_track_higher_than_system_dependent_threshold"
				} else {
					ts.PSR = "default_value"
				}
				if item.Secondary[2]&0x20 != 0 {
					ts.SSR = "age_last_ssr_track_higher_than_system_dependent_threshold"
				} else {
					ts.SSR = "default_value"
				}
				if item.Secondary[2]&0x10 != 0 {
					ts.MDS = "age_last_mode_s_track_higher_than_system_dependent_threshold"
				} else {
					ts.MDS = "default_value"
				}
				if item.Secondary[2]&0x08 != 0 {
					ts.ADS = "age_last_ads_b_track_higher_than_system_dependent_threshold"
				} else {
					ts.ADS = "default_value"
				}
				if item.Secondary[2]&0x04 != 0 {
					ts.SUC = "special_used_code"
				} else {
					ts.SUC = "default_value"
				}
				if item.Secondary[2]&0x02 != 0 {
					ts.AAC = "assigned_mode_a_code_conflict"
				} else {
					ts.AAC = "default_value"
				}

				if item.Secondary[2]&0x01 != 0 {
					tmp := item.Secondary[3] & 0xc0 >> 6
					switch tmp {
					case 0:
						ts.SDS = "combined"
					case 1:
						ts.SDS = "cooperative_only"
					case 2:
						ts.SDS = "non_cooperative_only"
					case 3:
						ts.SDS = "not_defined"
					}

					tmp = item.Secondary[3] & 0x38 >> 3
					switch tmp {
					case 0:
						ts.EMS = "no_emergency"
					case 1:
						ts.EMS = "general_emergency"
					case 2:
						ts.EMS = "lifeguard_medical"
					case 3:
						ts.EMS = "minimum_fuel"
					case 4:
						ts.EMS = "no_communications"
					case 5:
						ts.EMS = "unlawful_interference"
					case 6:
						ts.EMS = "downed_aircraft"
					case 7:
						ts.EMS = "undefined"
					}

					if item.Secondary[3]&0x04 != 0 {
						ts.PFT = "potential_false_track_indication"
					} else {
						ts.PFT = "no_indication"
					}
					if item.Secondary[3]&0x02 != 0 {
						ts.FPLT = "track_created_updated_fpl_data"
					} else {
						ts.FPLT = "default_value"
					}

					if item.Secondary[3]&0x01 != 0 {
						if item.Secondary[4]&0x80 != 0 {
							ts.DUPT = "duplicate_mode_3a_code"
						} else {
							ts.DUPT = "default_value"
						}
						if item.Secondary[4]&0x40 != 0 {
							ts.DUPF = "duplicate_flight_plan"
						} else {
							ts.DUPF = "default_value"
						}
						if item.Secondary[4]&0x20 != 0 {
							ts.DUPM = "duplicate_flight_plan_manual_correlation"
						} else {
							ts.DUPM = "default_value"
						}
						if item.Secondary[4]&0x10 != 0 {
							ts.SFC = "surface_target"
						} else {
							ts.SFC = "default_value"
						}
						if item.Secondary[4]&0x08 != 0 {
							ts.IDD = "duplicate_flight_id"
						} else {
							ts.IDD = "no_indication"
						}
						if item.Secondary[4]&0x04 != 0 {
							ts.IEC = "inconsistent_emergency_code"
						} else {
							ts.IEC = "default_value"
						}
					}
				}
			}
		}
	}

	return ts
}

// extractDerivedData returns Data derived directly by the aircraft.
func extractDerivedData(cp goasterix.Compound) DerivedData {
	var dd DerivedData
	for _, item := range cp.Secondary {
		switch item.Meta.FRN {
		case 1:
			dd.TargetAddress = strings.ToUpper(hex.EncodeToString(item.Fixed.Data))
		case 2:
			var payload [6]byte
			copy(payload[:], item.Fixed.Data[:])
			dd.TargetIdentification, _ = modeSIdentification(payload)
		case 3:
			dd.MagneticHeading = float64(uint16(item.Fixed.Data[0])<<8+uint16(item.Fixed.Data[1])) * 0.0055
		case 4:
			tmp := new(IAS)
			var lsb float64
			if item.Fixed.Data[0]&0x80 != 0 {
				tmp.IM = "mach"
				lsb = 0.001
			} else {
				tmp.IM = "ias"
				lsb = 0.000061035
			}
			tmp.AirSpeed = float64(uint16(item.Fixed.Data[0]&0x7f)<<8+uint16(item.Fixed.Data[1])) * lsb
			dd.IndicatedAirspeedOld = tmp
		case 5:
			dd.AirSpeed = uint16(item.Fixed.Data[0])<<8 + uint16(item.Fixed.Data[1])
		case 6:
			tmp := new(SelectedAltitude)
			if item.Fixed.Data[0]&0x80 != 0 {
				tmp.SAS = "source_information_provided"
			} else {
				tmp.SAS = "no_source_information_provided"
			}
			source := item.Fixed.Data[0] & 0x60 >> 5
			switch source {
			case 0:
				tmp.Source = "unknown"
			case 1:
				tmp.Source = "aircraft_altitude"
			case 2:
				tmp.Source = "fcu_mcp_selected_altitude"
			case 3:
				tmp.Source = "fms_selected_altitude"
			}
			data := uint16(item.Fixed.Data[0]&0x1f)<<8 + uint16(item.Fixed.Data[1])
			altitude := goasterix.TwoComplement16(13, data)
			tmp.Altitude = float64(altitude) * 25
			dd.SelectedAltitude = tmp
		case 7:
			tmp := new(StateSelectedAltitude)
			if item.Fixed.Data[0]&0x80 != 0 {
				tmp.MV = "manage_vertical_mode_active"
			} else {
				tmp.MV = "manage_vertical_mode_not_active"
			}
			if item.Fixed.Data[0]&0x40 != 0 {
				tmp.AH = "altitude_hold_active"
			} else {
				tmp.AH = "altitude_hold_not_active"
			}
			if item.Fixed.Data[0]&0x20 != 0 {
				tmp.AM = "approach_mode_active"
			} else {
				tmp.AM = "approach_mode_not_active"
			}

			data := uint16(item.Fixed.Data[0]&0x1f)<<8 + uint16(item.Fixed.Data[1])
			altitude := goasterix.TwoComplement16(13, data)
			tmp.Altitude = float64(altitude) * 25
			dd.StateSelectedAltitude = tmp
		case 26:
			dd.IndicatedAirSpeed = float64(uint16(item.Fixed.Data[0])<<8 + uint16(item.Fixed.Data[1]))
		case 27:
			dd.MachNumber = float64(uint16(item.Fixed.Data[0])<<8+uint16(item.Fixed.Data[1])) * 0.008
		}
	}
	return dd
}

// calculatedTrackPositionWGS84 returns Latitude and Longitude.
// Calculated Position in WGS-84 Co-ordinates with a resolution of 180/2^25 degrees
func calculatedTrackPositionWGS84(data [8]byte) PositionWGS84 {
	var pos PositionWGS84
	lsb := 180 / math.Pow(2, 25)
	pos.Latitude = float64(int32(data[0])<<24+int32(data[1])<<16+int32(data[2])<<8+int32(data[3])) * lsb
	pos.Longitude = float64(int32(data[4])<<24+int32(data[5])<<16+int32(data[6])<<8+int32(data[7])) * lsb
	return pos
}

// calculatedTrackPositionCartesian returns X and Y float64 in m
// Calculated position in Cartesian co-ordinates with a resolution of 0.5m
// LSB = 0.5
func calculatedTrackPositionCartesian(data [6]byte) CartesianXYPosition {
	var pos CartesianXYPosition

	tmpX := uint32(data[0])<<16 + uint32(data[1])<<8 + uint32(data[2])
	pos.X = float64(goasterix.TwoComplement32(24, tmpX)) * 0.5

	tmpY := uint32(data[3])<<16 + uint32(data[4])<<8 + uint32(data[5])
	pos.Y = float64(goasterix.TwoComplement32(24, tmpY)) * 0.5

	return pos
}

// calculatedTrackVelocityCartesian returns Vx and Vy float32 in m/s
// Calculated track velocity expressed in Cartesian co-ordinates
func calculatedTrackVelocityCartesian(data [4]byte) TrackVelocity {
	var vel TrackVelocity
	vel.Vx = float32(int16(data[0])<<8+int16(data[1])) / 4
	vel.Vy = float32(int16(data[2])<<8+int16(data[3])) / 4
	return vel
}

// calculatedAccelerationCartesian returns Ax and Ay float32 in m/s^2.
// Calculated Acceleration of the target expressed in Cartesian co-ordinates
// LSB = 0.25 m/s^2
func calculatedAccelerationCartesian(data [2]byte) Acceleration {
	var acc Acceleration
	acc.Ax = float32(int8(data[0])) / 4
	acc.Ay = float32(int8(data[1])) / 4
	return acc
}

// mode3ACode returns the squawk.
// Mode-3/A code converted into octal representation
func mode3ACode(data [2]byte) TrackMode3A {
	var mode3A TrackMode3A
	if data[0]&0x80 != 0 {
		mode3A.V = "code_not_validated"
	} else {
		mode3A.V = "code_validated"
	}

	if data[0]&0x40 != 0 {
		mode3A.G = "garbled_code"
	} else {
		mode3A.G = "default"
	}

	if data[0]&0x20 != 0 {
		mode3A.CH = "changed"
	} else {
		mode3A.CH = "no_change"
	}

	tmp := uint16(data[0])&0x000F<<8 + uint16(data[1])&0x00FF
	mode3A.Squawk = strconv.FormatUint(uint64(tmp), 8)

	return mode3A
}

// Target (aircraft or vehicle) identification in 8 characters
func targetIdentification(data [7]byte) TargetIdent {
	var target TargetIdent
	tmp := data[0] & 0xc0 >> 6

	switch tmp {
	case 0:
		target.STI = "downlinked_target"
	case 1:
		target.STI = "callsign_not_downlinked_target"
	case 2:
		target.STI = "registration_not_downlinked_target"
	case 3:
		target.STI = "invalid"
	}
	tmpData := [6]byte{data[1], data[2], data[3], data[4], data[5], data[6]}
	target.Target, _ = modeSIdentification(tmpData)

	return target
}

// measuredFlightLevel returns level in 100's ft
func measuredFlightLevel(data [2]byte) float32 {
	fl := float32(int16(data[0])<<8+int16(data[1])) / 4
	return fl
}

// trackGeometricAltitude returns a float32 in ft
// Vertical distance between the target and the projection of its position
// on the earth’s ellipsoid, as defined by WGS84
func trackGeometricAltitude(data [2]byte) float32 {
	fl := float32(int16(data[0])<<8+int16(data[1])) * 6.25
	return fl
}

// trackBarometricAltitude returns Altitude in FL
// Calculated Barometric Altitude of the track
func trackBarometricAltitude(data [2]byte) BarometricAltitude {
	var ba BarometricAltitude
	if data[0]&0x80 != 0 {
		ba.QNH = "qnh_correction_applied"
	} else {
		ba.QNH = "no_qnh_correction_applied"
	}
	ba.Altitude = float64((uint16(data[0])<<8+uint16(data[1]))&0x7FFF) / 4
	return ba
}

// rateOfClimbDescent returns a float32 in feet/minute
// Calculated rate of Climb/Descent of an aircraft in feet/minute
// A positive value indicates a climb, whereas a negative value indicates a descent.
func rateOfClimbDescent(data [2]byte) float32 {
	rate := float32(int16(data[0])<<8+int16(data[1])) * 6.25
	return rate
}
