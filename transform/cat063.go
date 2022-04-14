package transform

import (
	"github.com/mokhtarimokhtar/goasterix"
)

type Cat063Model struct {
	SacSic                SourceIdentifier `json:"dataSourceIdentifier"`
	TimeOfMessage         float64          `json:"timeOfMessage"`
	ServiceIdentification uint8            `json:"serviceIdentification,omitempty"`
	SensorIdentifier      SourceIdentifier `json:"sensorIdentifier"`
	SensorConfigStatus    *SensorStatus    `json:"sensorConfigStatus,omitempty"`
	TimeStampingBias      int16            `json:"timeStampingBias"`
	ModeSRangeGainAndBias *ModeSRange      `json:"modeSRangeGainAndBias,omitempty"`
	SSRModeSAzimuthBias   float64          `json:"ssrModeSAzimuthBias,omitempty"`
	PSRRangeGainAndBias   *PSRRange        `json:"psrRangeGainAndBias,omitempty"`
	PSRAzimuthBias        float64          `json:"psrAzimuthBias,omitempty"`
	PSRElevationBias      float64          `json:"psrElevationBias,omitempty"`
}

type ModeSRange struct {
	SRG float64 `json:"srg"`
	SRB float64 `json:"srb"`
}
type PSRRange struct {
	PRG float64 `json:"prg"`
	PRB float64 `json:"prb"`
}

func (data *Cat063Model) write(rec goasterix.Record) {
	for _, item := range rec.Items {
		switch item.Meta.FRN {
		case 1:
			// decode sac sic
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.SacSic, _ = sacSic(payload)

		case 2:
			// decode serviceIdentification
			data.ServiceIdentification = item.Fixed.Data[0]

		case 3:
			// decode timeOfDay
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfMessage, _ = timeOfDay(payload)
		case 4:
			//decode Sensor Identifier sac sic
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.SensorIdentifier, _ = sacSic(payload)
		case 5:
			// decode Sensor Configuration and Status
			tmp := extractSensorStatus(*item.Extended)
			data.SensorConfigStatus = &tmp
		case 6:
			//Time Stamping Bias
			data.TimeStampingBias = int16(item.Fixed.Data[0])<<8 + int16(item.Fixed.Data[1]) //008f => 8f00 shift to the left
		case 7:
			//Data Item I063/080, SSR / Mode S Range Gain and Bias
			tmp := new(ModeSRange)
			tmp.SRG = float64(int16(item.Fixed.Data[0])<<8+int16(item.Fixed.Data[1])) * 0.00001
			tmp.SRB = float64(int16(item.Fixed.Data[2])<<8+int16(item.Fixed.Data[3])) / 128
			data.ModeSRangeGainAndBias = tmp
		case 8:
			// I063/081 SSR/Mode S Azimuth Bias
			data.SSRModeSAzimuthBias = float64(int16(item.Fixed.Data[0])<<8+int16(item.Fixed.Data[1])) * 0.0055
		case 9:
			//Data Item I063/090, PSR Range Gain and Bias
			tmp := new(PSRRange)
			tmp.PRG = float64(int16(item.Fixed.Data[0])<<8+int16(item.Fixed.Data[1])) * 0.00001
			tmp.PRB = float64(int16(item.Fixed.Data[2])<<8+int16(item.Fixed.Data[3])) / 128
			data.PSRRangeGainAndBias = tmp
		case 10:
			//Data Item I063/091, PSR Azimuth Bias
			data.PSRAzimuthBias = float64(int16(item.Fixed.Data[0])<<8+int16(item.Fixed.Data[1])) * 0.0055
		case 11:
			//Data Item I063/092, PSR Elevation Bias
			data.PSRElevationBias = float64(int16(item.Fixed.Data[0])<<8+int16(item.Fixed.Data[1])) * 0.0055
		}

	}
}

type SensorStatus struct {
	CON string `json:"con"`
	PSR string `json:"psr"`
	SSR string `json:"ssr"`
	MDS string `json:"mds"`
	ADS string `json:"ads"`
	MLT string `json:"mlt"`
	OPS string `json:"ops,omitempty"`
	ODP string `json:"odp,omitempty"`
	OXT string `json:"oxt,omitempty"`
	MSC string `json:"msc,omitempty"`
	TSV string `json:"tsv,omitempty"`
	NPW string `json:"npw,omitempty"`
}

func extractSensorStatus(item goasterix.Extended) SensorStatus {
	var sr SensorStatus

	tmp := item.Primary[0] & 0xc0 >> 6
	switch tmp {
	case 0:
		sr.CON = "operational"
	case 1:
		sr.CON = "degraded"
	case 2:
		sr.CON = "initialization"
	case 3:
		sr.CON = "not_currently_connected"

	}

	if item.Primary[0]&0x20 != 0 {
		sr.PSR = "psr_nogo"
	} else {
		sr.PSR = "psr_go"
	}
	if item.Primary[0]&0x10 != 0 {
		sr.SSR = "ssr_nogo"
	} else {
		sr.SSR = "ssr_go"
	}
	if item.Primary[0]&0x08 != 0 {
		sr.MDS = "mode_s_nogo"
	} else {
		sr.MDS = "mode_s_go"
	}
	if item.Primary[0]&0x04 != 0 {
		sr.ADS = "ads_nogo"
	} else {
		sr.ADS = "ads_go"
	}
	if item.Primary[0]&0x02 != 0 {
		sr.MLT = "mlt_nogo"
	} else {
		sr.MLT = "mlt_go"
	}

	if item.Secondary != nil {
		if item.Secondary[0]&0x80 != 0 {
			sr.OPS = "operationnal_use_of_system_inhibited"
		} else {
			sr.OPS = "system_released_for_operationnal_use"
		}
		if item.Secondary[0]&0x40 != 0 {
			sr.ODP = "overload_in_dp"
		} else {
			sr.ODP = "default_no_overload"
		}
		if item.Secondary[0]&0x20 != 0 {
			sr.OXT = "overload_in_transmission_subsystem"
		} else {
			sr.OXT = "default_no_overload"
		}
		if item.Secondary[0]&0x10 != 0 {
			sr.MSC = "monitoring_system_disconnected"
		} else {
			sr.MSC = "monitoring_system_connected"
		}
		if item.Secondary[0]&0x08 != 0 {
			sr.TSV = "invalid"
		} else {
			sr.TSV = "valid"
		}
		if item.Secondary[0]&0x04 != 0 {
			sr.NPW = "no_plot_being_received"
		} else {
			sr.NPW = "default"
		}
	}

	return sr

}
