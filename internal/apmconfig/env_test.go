package apmconfig_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/apm-agent-go/internal/apmconfig"
)

func TestParseDurationEnv(t *testing.T) {
	const envKey = "ELASTIC_APM_TEST_DURATION"
	os.Setenv(envKey, "")

	d, err := apmconfig.ParseDurationEnv(envKey, 42*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, 42*time.Second, d)

	os.Setenv(envKey, "5s")
	d, err = apmconfig.ParseDurationEnv(envKey, 42*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, 5*time.Second, d)

	os.Setenv(envKey, "5ms")
	d, err = apmconfig.ParseDurationEnv(envKey, 42*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, 5*time.Millisecond, d)

	os.Setenv(envKey, "5m")
	d, err = apmconfig.ParseDurationEnv(envKey, 42*time.Minute)
	assert.NoError(t, err)
	assert.Equal(t, 5*time.Minute, d)

	os.Setenv(envKey, "5 h")
	_, err = apmconfig.ParseDurationEnv(envKey, 42*time.Second)
	assert.EqualError(t, err, "failed to parse ELASTIC_APM_TEST_DURATION: invalid character ' ' in duration 5 h")

	os.Setenv(envKey, "5h")
	_, err = apmconfig.ParseDurationEnv(envKey, 42*time.Second)
	assert.EqualError(t, err, "failed to parse ELASTIC_APM_TEST_DURATION: invalid unit in duration 5h (allowed units: ms, s, m)")

	os.Setenv(envKey, "5")
	_, err = apmconfig.ParseDurationEnv(envKey, 42*time.Second)
	assert.EqualError(t, err, "failed to parse ELASTIC_APM_TEST_DURATION: missing unit in duration 5 (allowed units: ms, s, m)")

	os.Setenv(envKey, "blah")
	_, err = apmconfig.ParseDurationEnv(envKey, 42*time.Second)
	assert.EqualError(t, err, "failed to parse ELASTIC_APM_TEST_DURATION: invalid duration blah")
}

func TestParseBoolEnv(t *testing.T) {
	const envKey = "ELASTIC_APM_TEST_BOOL"
	os.Setenv(envKey, "")

	b, err := apmconfig.ParseBoolEnv(envKey, true)
	assert.NoError(t, err)
	assert.True(t, b)

	os.Setenv(envKey, "true")
	b, err = apmconfig.ParseBoolEnv(envKey, false)
	assert.NoError(t, err)
	assert.True(t, b)

	os.Setenv(envKey, "false")
	b, err = apmconfig.ParseBoolEnv(envKey, true)
	assert.NoError(t, err)
	assert.False(t, b)

	os.Setenv(envKey, "falsk")
	_, err = apmconfig.ParseBoolEnv(envKey, true)
	assert.EqualError(t, err, `failed to parse ELASTIC_APM_TEST_BOOL: strconv.ParseBool: parsing "falsk": invalid syntax`)
}
