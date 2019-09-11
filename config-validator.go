package main

import (
	"errors"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Backend struct {
	Host string `yaml:"host"`
	Port int64  `yaml:"port"`
}

func (b *Backend) IsValid() error {
	if !(len(b.Host) > 0 && b.Port >= 80 && b.Port < 65536) {
		return errors.New("backend not valid")
	}
	return nil
}

type BackendConfiguration struct {
	Backends            []Backend `yaml:"backends"`
	ConnectTimeout      int64     `yaml:"connect_timeout"`
	FirstByteTimeout    int64     `yaml:"first_byte_timeout"`
	BetweenBytesTimeout int64     `yaml:"between_bytes_timeout"`
}

func (b *BackendConfiguration) IsValid() error {
	for _, backend := range b.Backends {
		err := backend.IsValid()
		if err != nil {
			return err
		}
	}
	if !(b.ConnectTimeout > 0 && b.FirstByteTimeout > 0 && b.BetweenBytesTimeout > 0) {
		return errors.New("backend configuration not valid")
	}
	return nil
}

type Probe struct {
	Url string `yaml:"url"`
}

func (p *Probe) IsValid() error {
	if !(len(p.Url) > 0) {
		return errors.New("url missing")
	}
	return nil
}

type ConfigElement struct {
	Name                 string               `yaml:"name"`
	Context              []string             `yaml:"context"`
	BackendConfiguration BackendConfiguration `yaml:"backend_configuration"`
	Probe                Probe                `yaml:"probe"`
}

func (c *ConfigElement) IsValid() error {
	backendValidationErr := c.BackendConfiguration.IsValid()
	if backendValidationErr != nil {
		return backendValidationErr
	}
	probeValidationErr := c.Probe.IsValid()
	if probeValidationErr != nil {
		return probeValidationErr
	}

	if !(len(c.Name) > 0 && len(c.Context) > 0) {
		return errors.New("name and context must not be empty")
	}

	return nil
}

func Validate(input []byte) error {
	var config []ConfigElement

	err := yaml.UnmarshalStrict(input, &config)
	if err != nil {
		if strings.Contains(err.Error(), "not found in type") {
			r := regexp.MustCompile(`.+field (.*?) not found in type.+`)
			field := r.FindStringSubmatch(err.Error())[1]
			return errors.New("field " + field + " is not supported")
		}

		return err
	}

	if !(len(config) > 0) {
		return errors.New("configuration is empty")
	}

	for _, c := range config {
		err = c.IsValid()
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "mosby"
	app.Usage = "I validate YOUR config. Yey"
	app.HideVersion = true

	var path string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "path to config-file",
			Destination: &path,
		},
	}

	app.Action = func(c *cli.Context) error {

		var data []byte
		var err error

		data, err = ioutil.ReadFile(path)

		if err != nil {
			log.Fatal(err)
		}
		return Validate(data)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
