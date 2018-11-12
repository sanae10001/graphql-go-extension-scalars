package scalars

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateTime_UnmarshalGraphQL(t *testing.T) {
	oriNow := time.Now()
	oriNowStr := oriNow.Format(time.RFC3339Nano)
	utcNow := oriNow.UTC()
	noNanoUTCNow := time.Date(utcNow.Year(), utcNow.Month(), utcNow.Day(), utcNow.Hour(), utcNow.Minute(), utcNow.Second(), 0, utcNow.Location())

	// String
	dt := DateTime{}
	err := dt.UnmarshalGraphQL(oriNowStr)
	assert.NoError(t, err)
	assert.True(t, dt.Equal(noNanoUTCNow))
}

func TestDateTime_MarshalJSON(t *testing.T) {
	oriNow := time.Now()
	utcNow := oriNow.UTC()
	noNanoUTCNow := time.Date(utcNow.Year(), utcNow.Month(), utcNow.Day(), utcNow.Hour(), utcNow.Minute(), utcNow.Second(), 0, utcNow.Location())
	noNanoUTCNowStr := noNanoUTCNow.Format(time.RFC3339)

	// Marshal RFC3339
	s := struct {
		D DateTime `json:"d"`
	}{DateTime{oriNow}}
	d, err := json.Marshal(s)
	assert.NoError(t, err)
	assert.Equal(t, string(d), fmt.Sprintf(`{"d":"%s"}`, noNanoUTCNowStr))

	// Unmarshal RFC3339
	data := fmt.Sprintf(`{"d": "%s"}`, noNanoUTCNowStr)
	s = struct {
		D DateTime `json:"d"`
	}{}
	err = json.Unmarshal([]byte(data), &s)
	assert.NoError(t, err)
	assert.True(t, s.D.Equal(noNanoUTCNow))

	// Unmarshal RFC3339Nano
	data = fmt.Sprintf(`{"d": "%s"}`, oriNow.Format(time.RFC3339Nano))
	s = struct {
		D DateTime `json:"d"`
	}{}
	err = json.Unmarshal([]byte(data), &s)
	assert.NoError(t, err)
	assert.True(t, s.D.Equal(oriNow))
}
