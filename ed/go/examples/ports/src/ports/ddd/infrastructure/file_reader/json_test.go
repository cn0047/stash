package file_reader

import (
	"strings"
	"testing"

	"ports/ddd/domain/model"
)

func TestRead(test *testing.T) {
	test.Run("one port", func(t *testing.T) {
		j := `{
		  "AEAJM": {
		    "name": "Ajman",
		    "city": "Ajman",
		    "country": "United Arab Emirates",
		    "alias": [],
		    "regions": [],
		    "coordinates": [
		      55.5136433,
		      25.4052165
		    ],
		    "province": "Ajman",
		    "timezone": "Asia/Dubai",
		    "unlocs": [
		      "AEAJM"
		    ],
		    "code": "52000"
		  }
		}`
		err := Read(strings.NewReader(j), func(port model.PortEntity) error {
			if port.ID != "AEAJM" {
				t.Errorf("Got: %v, want: AEAJM", port.ID)
			}
			return nil
		})
		if err != nil {
			t.Errorf("Got: %v, want: nil", err)
		}
	})
}
