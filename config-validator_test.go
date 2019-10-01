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

func TestProbe_IsValid(t *testing.T) {
	t.Run("valid probe", func(t *testing.T) {
		validProbe := Probe{Url: "/internal/health"}

		result := validProbe.IsValid()

		assert.Nil(t, result)
	})
	t.Run("invalid probe", func(t *testing.T) {
		validProbe := Probe{}

		result := validProbe.IsValid()

		assert.NotNil(t, result)
	})
}

func TestConfigElement_IsValid(t *testing.T) {
	t.Run("valid config element", func(t *testing.T) {
		validConfigElement := ConfigElement{
			Name:    "some-backend",
			Context: []string{"/context"},
			BackendConfiguration: BackendConfiguration{
				Backends: []Backend{
					{
						Host: "some-backend.some-host.com",
						Port: 443,
					},
				},
				ConnectTimeout:      1,
				FirstByteTimeout:    2,
				BetweenBytesTimeout: 3,
			},
			Probe: Probe{Url: "/health"},
			Local: false,
		}

		assert.Nil(t, validConfigElement.IsValid(true))
	})

	t.Run("config element with invalid probe", func(t *testing.T) {
		invalidConfigElement := ConfigElement{
			Name:    "some-backend",
			Context: []string{"/context"},
			BackendConfiguration: BackendConfiguration{
				Backends: []Backend{
					{
						Host: "some-backend.some-host.com",
						Port: 443,
					},
				},
				ConnectTimeout:      1,
				FirstByteTimeout:    2,
				BetweenBytesTimeout: 3,
			},
			Probe: Probe{},
			Local: false,
		}

		assert.NotNil(t, invalidConfigElement.IsValid(true))
	})

	t.Run("config element name missing", func(t *testing.T) {
		validConfigElement := ConfigElement{
			Name:    "",
			Context: []string{"/context"},
			BackendConfiguration: BackendConfiguration{
				Backends: []Backend{
					{
						Host: "some-backend.some-host.com",
						Port: 443,
					},
				},
				ConnectTimeout:      1,
				FirstByteTimeout:    2,
				BetweenBytesTimeout: 3,
			},
			Probe: Probe{Url: "/health"},
			Local: false,
		}

		assert.NotNil(t, validConfigElement.IsValid(true))
	})

	t.Run("config element context missing", func(t *testing.T) {
		validConfigElement := ConfigElement{
			Name:    "name",
			Context: []string{},
			BackendConfiguration: BackendConfiguration{
				Backends: []Backend{
					{
						Host: "some-backend.some-host.com",
						Port: 443,
					},
				},
				ConnectTimeout:      1,
				FirstByteTimeout:    2,
				BetweenBytesTimeout: 3,
			},
			Probe: Probe{Url: "/health"},
			Local: false,
		}

		assert.NotNil(t, validConfigElement.IsValid(true))
	})
}
