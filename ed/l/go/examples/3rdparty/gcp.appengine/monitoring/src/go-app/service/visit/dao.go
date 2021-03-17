package visit

import (
	"time"
)

type Visit struct {
	TimeStamp     time.Time `datastore:"timeStamp,noindex"`
	Path          string    `datastore:"path,noindex"`
	RemoteAddress string    `datastore:"remoteAddress,noindex"`
}
