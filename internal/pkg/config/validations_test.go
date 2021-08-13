package config

import (
	"testing"

	"kube-monkey/internal/pkg/config/param"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestValidateConfigs(t *testing.T) {
	SetDefaults()

	assert.Nil(t, ValidateConfigs())

	viper.Set(param.RunInterval, 601)
	assert.EqualError(t, ValidateConfigs(), "RunInterval: "+param.RunInterval+" is outside valid range of ]0,600]")
	viper.Set(param.RunInterval, 23)

	viper.Set(param.StartHour, 24)
	assert.EqualError(t, ValidateConfigs(), "StartHour: "+param.StartHour+" is outside valid range of [0,23]")
	viper.Set(param.StartHour, 23)

	viper.Set(param.EndHour, 24)
	assert.EqualError(t, ValidateConfigs(), "EndHour: "+param.EndHour+" is outside valid range of [0,23]")
	viper.Set(param.EndHour, 23)

	viper.Set(param.StartHour, 23)
	assert.EqualError(t, ValidateConfigs(), "StartHour: "+param.StartHour+" must be less than "+param.EndHour)
	viper.Set(param.StartHour, 22)

}

func TestIsValidHour(t *testing.T) {
	for i := 0; i <= 23; i++ {
		assert.True(t, IsValidHour(i))
	}
	assert.False(t, IsValidHour(24))
}

func TestIsValidHeader(t *testing.T) {
	header := "header1Key:header1Value"
	assert.True(t, isValidHeader(header))

	header = "header1/Key:header1/Value"
	assert.True(t, isValidHeader(header))

	header = "header1:{$env:VARIABLE_NAME}"
	assert.True(t, isValidHeader(header))

	header = "header1Key"
	assert.False(t, isValidHeader(header))

	header = "header1Key:"
	assert.False(t, isValidHeader(header))
}

func TestIsValidInterval(t *testing.T) {
	assert.False(t, IsValidInterval(0))
	assert.False(t, IsValidInterval(601))
	assert.True(t, IsValidInterval(42))
}
