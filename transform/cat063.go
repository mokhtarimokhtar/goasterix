package transform

import (
	"github.com/mokhtarimokhtar/goasterix"
)


type Cat063Model struct {
	SacSic                 *SourceIdentifier   `json:"dataSourceIdentifier"`
	TimeOfMessage          float64             `json:"timeOfMessage"`
}


func (data *Cat063Model) write(rec goasterix.Record) {
	for _, item := range rec.Items {
		switch item.Meta.FRN {
		case 1:
			// decode sac sic
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp, _ := sacSic(payload)
			data.SacSic = &tmp
		case 3:
			// decode timeOfDay
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfMessage, _ = timeOfDay(payload)
		}
	}
}