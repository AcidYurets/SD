package index

import (
	"encoding/json"
	"time"
)

const IndexedAtField = "@indexed_at"

type IndexedAtData struct {
	Data      interface{}
	IndexedAt time.Time
}

func (d IndexedAtData) MarshalJSON() ([]byte, error) {
	if d.IndexedAt.IsZero() {
		d.IndexedAt = time.Now()
	}

	pJSON, err := json.Marshal(map[string]interface{}{
		IndexedAtField: d.IndexedAt,
	})
	if err != nil {
		return nil, err
	}

	eJSON, err := json.Marshal(d.Data)
	if err != nil {
		return nil, err
	}

	eJSON[0] = ','
	return append(pJSON[:len(pJSON)-1], eJSON...), nil
}
