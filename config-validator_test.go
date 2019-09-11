package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("valid config results in no error", func(t *testing.T) {
		input, _ := ioutil.ReadFile("test-resources/valid.config.yaml")

		result := Validate(input, false)

		assert.Equal(t, nil, result)
	})
	t.Run("empty file results in an error", func(t *testing.T) {
		input, _ := ioutil.ReadFile("test-resources/empty.config.yaml")

		result := Validate(input, false)

		assert.Equal(t, errors.New("configuration is empty"), result)
	})
	t.Run("invalid port results in an error", func(t *testing.T) {
		input, _ := ioutil.ReadFile("test-resources/invalid.port.yaml")

		result := Validate(input, false)

		assert.Equal(t, errors.New("backend not valid"), result)
	})
	t.Run("missing connect_timeout results in an error", func(t *testing.T) {
		input, _ := ioutil.ReadFile("test-resources/missing.connect_timeout.yaml")

		result := Validate(input, false)

		assert.Equal(t, errors.New("backend configuration not valid"), result)
	})
	t.Run("unknown parameter should result in error", func(t *testing.T) {
		input, _ := ioutil.ReadFile("test-resources/unknown.parameter.yaml")

		result := Validate(input, false)

		assert.Equal(t, errors.New("field bar is not supported"), result)
	})
	t.Run("nested unknown parameter should result in error", func(t *testing.T) {
		input, _ := ioutil.ReadFile("test-resources/nested.unknown.parameter.yaml")

		result := Validate(input, false)

		assert.Equal(t, errors.New("field foo is not supported"), result)
	})
	t.Run("specifying local in non-local mode should return an error", func(t *testing.T) {
		input, _ := ioutil.ReadFile("test-resources/valid.config.yaml")

		result := Validate(input, true)

		assert.Equal(t, errors.New("local: true is not allowed in non-local environments"), result)
	})

}
