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

		result := Validate(input)

		assert.Equal(t, nil, result)
	})
	t.Run("empty file results in an error", func(t *testing.T) {
		input, _ := ioutil.ReadFile("test-resources/empty.config.yaml")

		result := Validate(input)

		assert.Equal(t, errors.New("configuration is empty"), result)
	})
	t.Run("invalid port results in an error", func(t *testing.T) {
		input, _ := ioutil.ReadFile("test-resources/invalid.port.yaml")

		result := Validate(input)

		assert.Equal(t, errors.New("backend not valid"), result)
	})
	t.Run("missing connect_timeout results in an error", func(t *testing.T) {
		input, _ := ioutil.ReadFile("test-resources/missing.connect_timeout.yaml")

		result := Validate(input)

		assert.Equal(t, errors.New("backend configuration not valid"), result)
	})

}
