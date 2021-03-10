package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"

	"gopkg.in/yaml.v2"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

func printhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello server I am client, %s!", r.URL.Path[1:])
}
func main() {
	log.Println("Starting echo server on : ", runtime.GOOS)
	http.HandleFunc("/", printhello)
	http.ListenAndServe(":8080", nil) // nolint
}

// ControllerOptions configuration
type ControllerOptions struct {
	Services bool `yaml:"services"`
}

// Config struct contains operator controller configuration
type Config struct {
	ControllerOptions ControllerOptions `yaml:"controller"`
}

// Load loads configuration from config file
func (c *Config) Load() error {
	fmt.Println("in load")
	files, err := ioutil.ReadDir("/tmpdir")
	if err != nil {
		logf.Log.WithName("configmap").Error(err, "error while reading")
		return err
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
	fmt.Println("list over")
	file, err := os.Open("/tmpdir/aporeto-config.yaml")
	if err != nil {
		logf.Log.WithName("configmap").Error(err, " could not open the file")
		return err
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		logf.Log.WithName("configmap").Error(err, "error while reading")
		return err
	}
	fmt.Println("COnfig: ", string(b))
	if len(b) != 0 {
		return yaml.Unmarshal(b, c)
	}
	return nil
}

// NewConfigMap creates new config object
func NewConfigMap() (*Config, error) {
	c := &Config{}
	if err := c.Load(); err != nil {
		return c, err
	}

	return c, nil
}
