package transform

import (
	"github.com/mokhtarimokhtar/goasterix"
)

type Cat063Model struct {
	SacSic                SourceIdentifier `json:"dataSourceIdentifier"`
	TimeOfMessage         float64          `json:"timeOfMessage"`
	ServiceIdentification uint8            `json:"serviceIdentification,omitempty"`
	SensorIdentifier      SourceIdentifier `json:"sensorIdentifier"`
}

func (data *Cat063Model) write(rec goasterix.Record) {
	for _, item := range rec.Items {
		switch item.Meta.FRN {
		case 1:
			// decode sac sic
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.SacSic, _ = sacSic(payload)

			// TODO case 2
		case 2:
			// decode serviceIdentification
			//var payload [1]byte
			//copy(payload[:], item.Fixed.Data[:])
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

		}

	}
}
