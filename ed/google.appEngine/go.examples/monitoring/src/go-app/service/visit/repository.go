package visit

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"

	"go-app/config/taxonomy"
)

func TrackVisit(ctx context.Context, r *http.Request) (datastore.Key, error) {
	v := Visit{
		TimeStamp:     time.Now().UTC(),
		Path:          r.URL.Path,
		RemoteAddress: r.RemoteAddr,
	}
	key := datastore.NewIncompleteKey(ctx, taxonomy.DataStoreKindVisit, nil)

	k, err := datastore.Put(ctx, key, &v)
	if err != nil {
		return datastore.Key{}, errors.New("[20180703-006] Failed to store visit, error: " + err.Error())
	}

	return *k, nil
}

func GetCount(ctx context.Context) (int, error) {
	q := datastore.NewQuery(taxonomy.DataStoreKindVisit)
	count, err := q.Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("[20180703-007] failed to get count: %v", err)
	}

	return count, nil
}
