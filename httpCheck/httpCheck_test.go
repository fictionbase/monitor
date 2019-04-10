package httpcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResponseAndTime(t *testing.T) {
	resp, elapsed, err := GetResponseAndTime("https://portfolio.seike460.com/")
	assert.NotNil(t, resp)
	assert.Equal(t, resp.StatusCode, 200)
	assert.NotEqual(t, elapsed, float64(0))
	assert.NotNil(t, elapsed)
	assert.Nil(t, err)
}

func TestGetResponseAndTimeNotFound(t *testing.T) {
	resp, elapsed, err := GetResponseAndTime("https://NotFound.seike460.com/")
	assert.Nil(t, resp)
	assert.Equal(t, elapsed, float64(0))
	assert.NotNil(t, err)
}
