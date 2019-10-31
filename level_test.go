package logoped

import (
	"encoding/json"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CreateLogLevel(t *testing.T) {

	level := CreateLogLevel(`asd`, 123)

	assert.IsType(t, &LogLevel{}, level)
}

func Test_Name(t *testing.T) {

	name, err := CreateRandomString([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"), 16)

	assert.Nil(t, err)

	level := CreateLogLevel(name, 123)

	assert.Equal(t, name, level.Name())
}

func Test_Weight(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	weight := rand.Intn(1000)

	level := CreateLogLevel(`qwe`, weight)

	assert.Equal(t, weight, level.Weight())
}

func Test_MarshalJSON(t *testing.T) {

	levels := []ILogLevel{
		LogLevelInfo,
		LogLevelWarn,
		LogLevelError,
		LogLevelFatal,
	}
	for _, level := range levels {
		bytesGot, errGot := json.Marshal(level)
		assert.Nil(t, errGot)

		bytesExpected, errExpexted := json.Marshal(level.Name())
		assert.Nil(t, errExpexted)

		assert.Equal(t, bytesExpected, bytesGot)
	}
}
