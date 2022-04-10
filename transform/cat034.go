package transform

import (
	"encoding/hex"
	"errors"
	"github.com/mokhtarimokhtar/goasterix"
	"strconv"
)

const (
	sysIn          string = "system_inhibited"
	sysOp          string = "system_operational"
	antenna1       string = "antenna_1"
	antenna2       string = "antenna_2"
	noChanSelected string = "no_channel_selected"
	chanASelected  string = "channel_a_only_selected"
	chanBSelected  string = "channel_b_only_selected"
	chanABSelected string = "channel_a_and_b_selected"
	chanIllCombi   string = "illegal_combination"
	overload       string = "overload"
	noOverload     string = "no_overload"
	mscC           string = "monitoring_system_connected"
	mscD           string = "monitoring_system_disconnected"
	chanAuse       string = "channel_a_in_use"
	chanBuse       string = "channel_b_in_use"
)

var ErrTypeUnknown = errors.New("[ASTERIX Error CAT034] Message TYPE Unknown")

type collimationError struct {
	RangeError   float64 `json:"rangeError"`
	AzimuthError float64 `json:"azimuthError"`
}

type GenericPolarWindow struct {
	RhoStart   float64 `json:"rhoStart"`
	RhoEnd     float64 `json:"rhoEnd"`
	ThetaStart float64 `json:"thetaStart"`
	ThetaEnd   float64 `json:"thetaEnd"`
}

type MessageCounter struct {
	Type    string `json:"type"`
	Counter uint16 `json:"counter"`
}

type ComSysConf struct {
	Nogo   string `json:"nogo"`
	Rdpc   string `json:"rdpc"`
	Rdpr   string `json:"rdpr"`
	Ovlrdp string `json:"ovlrdp"`
	Ovlxmt string `json:"ovlxmt"`
	Msc    string `json:"msc"`
	Tsv    string `json:"tsv"`
}
type PsrSsrSysConf struct {
	Ant  string `json:"ant"`
	ChAB string `json:"chAB"`
	Ovl  string `json:"ovl"`
	Msc  string `json:"msc"`
}
type MdsSysConf struct {
	Ant    string `json:"ant"`
	ChAB   string `json:"chAB"`
	Ovlsur string `json:"ovlsur"`
	Msc    string `json:"msc"`
	Scf    string `json:"scf"`
	Dlf    string `json:"dlf"`
	Ovlscf string `json:"ovlscf"`
	Ovldlf string `json:"ovldlf"`
}
type SysConf struct {
	Com *ComSysConf    `json:"com,omitempty"`
	Psr *PsrSsrSysConf `json:"psr,omitempty"`
	Ssr *PsrSsrSysConf `json:"ssr,omitempty"`
	Mds *MdsSysConf    `json:"mds,omitempty"`
}

type Pos3D struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
	Height    uint16  `json:"height,omitempty"`
}

type Cat034Model struct {
	SacSic                 *SourceIdentifier   `json:"sourceIdentifier,omitempty"`
	MessageType            string              `json:"messageType,omitempty"`
	TimeOfDay              float64             `json:"timeOfDay,omitempty"`
	SectorNumber           float64             `json:"sectorNumber,omitempty"`
	AntennaRotationSpeed   float64             `json:"antennaRotationSpeed,omitempty"`
	SystemConfiguration    *SysConf            `json:"systemConfiguration,omitempty"`
	SystemProcessingMode   *SysProcess         `json:"systemProcessingMode,omitempty"`
	MessageCountValues     []MessageCounter    `json:"messageCountValues,omitempty"`
	GenericPolarWindow     *GenericPolarWindow `json:"genericPolarWindow,omitempty"`
	DataFilter             string              `json:"dataFilter,omitempty"`
	Position3DofDataSource *Pos3D              `json:"position3DofDataSource,omitempty"`
	CollimationError       *collimationError   `json:"collimationError,omitempty"`
	REDataItem             string              `json:"reDataItem,omitempty"`
	SPDataItem             string              `json:"spDataItem,omitempty"`
}

func (data *Cat034Model) write(rec goasterix.Record) {
	for _, item := range rec.Items {
		switch item.Meta.FRN {
		case 1:
			// decode sac sic
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp, _ := sacSic(payload)
			data.SacSic = &tmp
		case 2:
			//decode messageTypeCat034
			var payload [1]byte
			copy(payload[:], item.Fixed.Data[:])
			data.MessageType = messageTypeCat034(payload)
		case 3:
			// decode timeOfDay
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfDay, _ = timeOfDay(payload)
		case 4:
			// decode sector number
			// SectorNumber returns a float.
			// Ref: 5.2.3 Records Item I034/020.
			// Eight most significant bits of the antenna azimuth defining a particular azimuth sector.
			data.SectorNumber = float64(item.Fixed.Data[0]) * 1.40625
		case 5:
			// AntennaRotationSpeed returns a float in second.
			// Antenna rotation period as measured between two consecutive
			// North crossings or as averaged during a period of time.
			// Ref: 5.2.3 Records Item I034/041.
			data.AntennaRotationSpeed = float64(uint16(item.Fixed.Data[0])<<8+uint16(item.Fixed.Data[1])) / 128
		case 6:
			tmp := systemConfiguration(*item.Compound)
			data.SystemConfiguration = &tmp
		case 7:
			tmp := systemProcessingMode(*item.Compound)
			data.SystemProcessingMode = &tmp
		case 8:
			// todo fix
			tmp, _ := messageCountValues(item.Fixed.Data)
			data.MessageCountValues = tmp
		case 9:
			var payload [8]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := genericPolarWindow(payload)
			data.GenericPolarWindow = &tmp
		case 10:
			var payload [1]byte
			copy(payload[:], item.Fixed.Data[:])
			data.DataFilter, _ = dataFilter(payload)
		case 11:
			var payload [8]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := position3DofDataSource(payload)
			data.Position3DofDataSource = &tmp
		case 12:
			// collimationError returns an array float64.
			// RANGE ERROR and AZIMUTH ERROR
			// Ref: 5.2.9 Records Item I034/090.
			tmp := new(collimationError)
			tmp.RangeError = float64(int8(item.Fixed.Data[0])) / 128
			tmp.AzimuthError = float64(int8(item.Fixed.Data[1])) * 0.021972656
			data.CollimationError = tmp
		case 13:
			data.REDataItem = hex.EncodeToString(item.Fixed.Data)
		case 14:
			data.SPDataItem = hex.EncodeToString(item.Fixed.Data)
		}
	}
}

// MessageType returns a string of message type.
// Ref. 5.2.1 Data Item I034/000, Message Type
func messageTypeCat034(data [1]byte) string {
	var msg string
	msgType := data[0]

	switch msgType {
	case 1:
		msg = "north_marker_message"
	case 2:
		msg = "sector_crossing_message"
	case 3:
		msg = "geographical_filtering_message"
	case 4:
		msg = "jamming_strobe_message"
	case 5:
		msg = "solar_storm_message"
	case 6:
		msg = "ssr_jamming_strobe_message"
	case 7:
		msg = "mode_s_jamming_strobe_message"
	default:
		msg = "undefined_message_type"
	}

	return msg
}

// systemConfiguration returns map of map string.
// Ref: 5.2.6 Data Item I034/050, System Configuration and Status
func systemConfiguration(cp goasterix.Compound) SysConf {
	var sysConf SysConf

	for _, item := range cp.Secondary {
		switch item.Meta.FRN {
		case 1:
			com := new(ComSysConf)
			tmp := item.Fixed.Data[0]
			if tmp&0x80 == 0 {
				com.Nogo = sysIn
			} else {
				com.Nogo = sysOp
			}
			if tmp&0x40 == 0 {
				com.Rdpc = "radar_data_processor_chain1"
			} else {
				com.Rdpc = "radar_data_processor_chain2"
			}
			if tmp&0x20 == 0 {
				com.Rdpr = "default_situation"
			} else {
				com.Rdpr = "reset_of_rdpc"
			}
			if tmp&0x10 == 0 {
				com.Ovlrdp = noOverload
			} else {
				com.Ovlrdp = overload
			}
			if tmp&0x08 == 0 {
				com.Ovlxmt = noOverload
			} else {
				com.Ovlxmt = overload
			}
			if tmp&0x04 == 0 {
				com.Msc = mscC
			} else {
				com.Msc = mscD
			}
			if tmp&0x02 == 0 {
				com.Tsv = "time_source_valid"
			} else {
				com.Tsv = "time_source_invalid"
			}
			sysConf.Com = com
		case 4:
			psr := new(PsrSsrSysConf)
			tmp := item.Fixed.Data[0]
			if tmp&0x80 == 0 {
				psr.Ant = antenna1
			} else {
				psr.Ant = antenna2
			}

			tmpChAB := tmp & 0x60 >> 5
			if tmpChAB == 0 {
				psr.ChAB = noChanSelected
			} else if tmpChAB == 1 {
				psr.ChAB = chanASelected
			} else if tmpChAB == 2 {
				psr.ChAB = chanBSelected
			} else if tmpChAB == 3 {
				psr.ChAB = chanABSelected
			}
			if tmp&0x10 == 0 {
				psr.Ovl = noOverload
			} else {
				psr.Ovl = overload
			}
			if tmp&0x08 == 0 {
				psr.Msc = mscC
			} else {
				psr.Msc = mscD
			}
			sysConf.Psr = psr
		case 5:
			ssr := new(PsrSsrSysConf)
			tmp := item.Fixed.Data[0]

			if tmp&0x80 == 0 {
				ssr.Ant = antenna1
			} else {
				ssr.Ant = antenna2
			}

			tmpChAB := tmp & 0x60 >> 5
			if tmpChAB == 0 {
				ssr.ChAB = noChanSelected
			} else if tmpChAB == 1 {
				ssr.ChAB = chanASelected
			} else if tmpChAB == 2 {
				ssr.ChAB = chanBSelected
			} else if tmpChAB == 3 {
				ssr.ChAB = chanABSelected
			}
			if tmp&0x10 == 0 {
				ssr.Ovl = noOverload
			} else {
				ssr.Ovl = overload
			}
			if tmp&0x08 == 0 {
				ssr.Msc = mscC
			} else {
				ssr.Msc = mscD
			}
			sysConf.Ssr = ssr
		case 6:
			mds := new(MdsSysConf)
			tmp := item.Fixed.Data[0]

			if tmp&0x80 == 0 {
				mds.Ant = antenna1
			} else {
				mds.Ant = antenna2
			}

			tmpChAB := tmp & 0x60 >> 5
			if tmpChAB == 0 {
				mds.ChAB = noChanSelected
			} else if tmpChAB == 1 {
				mds.ChAB = chanASelected
			} else if tmpChAB == 2 {
				mds.ChAB = chanBSelected
			} else if tmpChAB == 3 {
				mds.ChAB = chanIllCombi
			}

			if tmp&0x10 == 0 {
				mds.Ovlsur = noOverload
			} else {
				mds.Ovlsur = overload
			}
			if tmp&0x08 == 0 {
				mds.Msc = mscC
			} else {
				mds.Msc = mscD
			}
			if tmp&0x04 == 0 {
				mds.Scf = chanAuse
			} else {
				mds.Scf = chanBuse
			}
			if tmp&0x02 == 0 {
				mds.Dlf = chanAuse
			} else {
				mds.Dlf = chanBuse
			}
			if tmp&0x01 == 0 {
				mds.Ovlscf = noOverload
			} else {
				mds.Ovlscf = overload
			}
			if item.Fixed.Data[1]&0x80 == 0 {
				mds.Ovldlf = noOverload
			} else {
				mds.Ovldlf = overload
			}
			sysConf.Mds = mds
		}
	}
	return sysConf
}

type ComSysPro struct {
	Redrdp string `json:"redrdp"`
	Redxmt string `json:"redxmt"`
}
type PsrSysPro struct {
	Pol    string `json:"pol"`
	Redrad string `json:"redrad"`
	Stc    string `json:"stc"`
}
type SsrSysPro struct {
	Redrad string `json:"redrad"`
}
type MdsSysPro struct {
	Redrad string `json:"redrad"`
	Clu    string `json:"clu"`
}
type SysProcess struct {
	ComSysPro *ComSysPro `json:"com,omitempty"`
	Psr       *PsrSysPro `json:"psr,omitempty"`
	Ssr       *SsrSysPro `json:"ssr,omitempty"`
	Mds       *MdsSysPro `json:"mds,omitempty"`
}

// systemProcessingMode returns map of map string.
// Ref: Data Item I034/060, System Processing Mode
func systemProcessingMode(cp goasterix.Compound) SysProcess {
	var sysProc SysProcess
	for _, item := range cp.Secondary {
		switch item.Meta.FRN {
		case 1:
			tmp := new(ComSysPro)

			tmpRedrdp := item.Fixed.Data[0] & 0x70 >> 4
			if tmpRedrdp == 0 {
				tmp.Redrdp = "no_reduction_active"
			} else {
				tmp.Redrdp = "reduction_step_" + strconv.Itoa(int(tmpRedrdp)) + "_active"
			}

			tmpRedxmt := item.Fixed.Data[0] & 0x0E >> 1
			if tmpRedxmt == 0 {
				tmp.Redxmt = "no_reduction_active"
			} else {
				tmp.Redxmt = "reduction_step_" + strconv.Itoa(int(tmpRedxmt)) + "_active"
			}
			sysProc.ComSysPro = tmp
		case 4:
			tmp := new(PsrSysPro)
			if item.Fixed.Data[0]&0x80 == 0 {
				tmp.Pol = "linear_polarization"
			} else {
				tmp.Pol = "circular_polarization"
			}

			tmpRedrdp := item.Fixed.Data[0] & 0x70 >> 4
			if tmpRedrdp == 0 {
				tmp.Redrad = "no_reduction_active"
			} else {
				tmp.Redrad = "reduction_step_" + strconv.Itoa(int(tmpRedrdp)) + "_active"
			}

			tmpStc := item.Fixed.Data[0] & 0x0C >> 2
			tmp.Stc = "stcMap_" + strconv.Itoa(int(tmpStc)+1)
			sysProc.Psr = tmp
		case 5:
			tmp := new(SsrSysPro)
			tmpRedrad := item.Fixed.Data[0] & 0xE0 >> 5
			if tmpRedrad == 0 {
				tmp.Redrad = "no_reduction_active"
			} else {
				tmp.Redrad = "reduction_step_" + strconv.Itoa(int(tmpRedrad)) + "_active"
			}
			sysProc.Ssr = tmp
		case 6:
			tmp := new(MdsSysPro)

			tmpRedrad := item.Fixed.Data[0] & 0xE0 >> 5
			if tmpRedrad == 0 {
				tmp.Redrad = "no_reduction_active"
			} else {
				tmp.Redrad = "reduction_step_" + strconv.Itoa(int(tmpRedrad)) + "_active"
			}

			if item.Fixed.Data[0]&0x10 == 0 {
				tmp.Clu = "autonomous"
			} else {
				tmp.Clu = "no_autonomous"
			}
			sysProc.Mds = tmp
		}
	}
	return sysProc
}

// messageCountValues
// Message Count values, according the various types of messages, for the last completed antenna revolution,
// counted between two North crossings.
// Ref. 5.2.8 Data Item I034/070, Message Count Values
func messageCountValues(data []byte) ([]MessageCounter, error) {
	var mcv []MessageCounter
	var err error
	rep := data[0]

	for i := 0; i < int(rep*2); i = i + 2 {
		m := MessageCounter{}
		m.Counter = uint16(data[i+1]&0x07)<<8 + uint16(data[i+2])

		typeMCtmp := uint16(data[i+1] & 0xF8 >> 3)
		switch typeMCtmp {
		case 0:
			m.Type = "no_detection"
		case 1:
			m.Type = "single_psr_target_reports"
		case 2:
			m.Type = "single_ssr_target_reports"
		case 3:
			m.Type = "ssr_psr_target_reports"
		case 4:
			m.Type = "single_all_call_target_reports"
		case 5:
			m.Type = "single_roll_call_target_reports"
		case 6:
			m.Type = "all_call_psr_target_reports"
		case 7:
			m.Type = "roll_call_psr_target_reports"
		case 8:
			m.Type = "filter_for_weather_data"
		case 9:
			m.Type = "filter_for_jamming_strobe"
		case 10:
			m.Type = "filter_for_psr_data"
		case 11:
			m.Type = "filter_for_ssr_mode_s_data"
		case 12:
			m.Type = "filter_for_ssr_mode_s_psr_data"
		case 13:
			m.Type = "filter_for_enhanced_surveillance_data"
		case 14:
			m.Type = "filter_for_psr_enhanced_surveillance"
		case 15:
			m.Type = "filter_for_psr_enhanced_surveillance_ssr_mode_s_data_not_in_area"
		case 16:
			m.Type = "filter_for_psr_enhanced_surveillance__all_ssr_mode_s_data"
		case 17:
			m.Type = "re_interrogations_per_sector"
		case 18:
			m.Type = "bds_swap_and_wrong_df_replies_per_sector"
		case 19:
			m.Type = "mode_ac_fruit_per_sector"
		case 20:
			m.Type = "mode_s_fruit_per_sector"
		default:
			m.Type = "unknown"
			err = ErrTypeUnknown
		}
		mcv = append(mcv, m)
	}
	return mcv, err
}

// genericPolarWindow returns a map of float64.
// rhoStart and rhoEnd (NM),  thetaStart and thetaEnd degrees
// Ref: 5.2.10 Data Item I034/100
func genericPolarWindow(data [8]byte) GenericPolarWindow {
	var g GenericPolarWindow
	g.RhoStart = float64(uint16(data[0])<<8+uint16(data[1])) / 256
	g.RhoEnd = float64(uint16(data[2])<<8+uint16(data[3])) / 256
	g.ThetaStart = float64(uint16(data[4])<<8+uint16(data[5])) * 0.0055
	g.ThetaEnd = float64(uint16(data[6])<<8+uint16(data[7])) * 0.0055
	return g
}

// dataFilter returns an integer.
// Ref: 5.2.11 Data Item I034/110
func dataFilter(data [1]byte) (df string, err error) {
	tmp := data[0]
	switch tmp {
	case 0:
		df = "invalid_value"
	case 1:
		df = "filter_weather_data"
	case 2:
		df = "filter_jamming_strobe"
	case 3:
		df = "filter_psr_data"
	case 4:
		df = "filter_ssr_modes_data"
	case 5:
		df = "filter_ssr_modes_psr_data"
	case 6:
		df = "enhanced_surveillance_data"
	case 7:
		df = "filter_psr_enhanced_surveillance_data"
	case 8:
		df = "filter_psr_enhanced_surveillance_ssr_modes_data_not_in_area_prime_interest"
	case 9:
		df = "filter_psr_enhanced_surveillance_all_ssr_modes_data"
	default:
		df = "error_undefined"
	}

	return df, nil
}

// position3DofDataSource returns a map of float with
// height in metre, latitude and longitude WGS84.
// 3D-Position of Data Source in WGS 84 Co-ordinates.
// Ref: 5.2.12 Data Item I034/120 3D-Position Of Data Source
func position3DofDataSource(data [8]byte) Pos3D {
	var pos Pos3D
	pos.Height = uint16(data[0])<<8 + uint16(data[1])

	tmpLatitude := uint32(data[2])<<16 + uint32(data[3])<<8 + uint32(data[4])
	pos.Latitude = float32(goasterix.TwoComplement32(24, tmpLatitude)) * 0.000021458

	tmpLongitude := uint32(data[5])<<16 + uint32(data[6])<<8 + uint32(data[7])
	pos.Longitude = float32(goasterix.TwoComplement32(24, tmpLongitude)) * 0.000021458

	return pos
}
