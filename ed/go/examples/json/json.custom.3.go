package models

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"common/api"
	"common/guid"
)

type ShortLink struct {
	ID                string    `datastore:"-" goon:"id" json:"id"`
	Target            string    `datastore:"target" json:"target"`
	TargetHash        string    `datastore:"target_hash" json:"target_hash"`
	Source            string    `datastore:"source" json:"source"`
	SourceDescription string    `datastore:"source_description" json:"source_description"`
	DomainID          string    `datastore:"domain_id" json:"domain_id"`
	Permanent         bool      `datastore:"permanent" json:"permanent"`
	CreatedDate       time.Time `datastore:"created_date" json:"created_date"`
	UpdatedDate       time.Time `datastore:"updated_date" json:"updated_date"`
	LastAccessedDate  time.Time `datastore:"last_accessed_date" json:"last_accessed_date"`
	Metadata          []byte    `datastore:"metadata" json:"metadata"`
}

func (s ShortLink) GetMetadata() (interface{}, error) {
	var data interface{}
	err := json.Unmarshal(s.Metadata, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal metadata, error: %s", err)
	}

	return data, nil
}

func (s ShortLink) MarshalJSON() ([]byte, error) {
	m, err := s.GetMetadata()
	if err != nil {
		return nil, err
	}

	return json.Marshal(map[string]interface{}{
		"id":                 s.ID,
		"target":             s.Target,
		"target_hash":        s.TargetHash,
		"source":             s.Source,
		"source_description": s.SourceDescription,
		"domain_id":          s.DomainID,
		"permanent":          s.Permanent,
		"created_date":       s.CreatedDate,
		"updated_date":       s.UpdatedDate,
		"last_accessed_date": s.LastAccessedDate,
		"metadata":           m,
	})
}

func PutShortLink(cctx *api.CliqueContext, input ShortLinkInput) (ShortLink, error) {
	m, err := json.Marshal(input.Metadata)
	if err != nil {
		return ShortLink{}, fmt.Errorf("failed to marshal metadata, error: %s", err)
	}

	hash := sha1.New()
	_, er := io.WriteString(hash, input.Target)
	if er != nil {
		return ShortLink{}, fmt.Errorf("failed to make hast out of target, error: %s", er)
	}
	h := fmt.Sprintf("%x", hash.Sum(nil))

	entity := ShortLink{
		ID:          guid.New(),
		Target:      input.Target,
		TargetHash:  h,
		Source:      input.Source,
		DomainID:    input.DomainID,
		Permanent:   input.Permanent,
		CreatedDate: time.Now().UTC(),
		Metadata:    m,
	}

	cctx.MustPut(&entity)

	return entity, nil
}
