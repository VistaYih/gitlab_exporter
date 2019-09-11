package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

// Project holds information about a GitLab project
type Project struct {
	Name string
	Refs string
}

// Wildcard is a specific handler to dynamically search projects
type Wildcard struct {
	Search string
	Owner  struct {
		Name string
		Kind string
	}
	Refs string
}

// Config represents what can be defined as a yaml config file
type Config struct {
	URL           string `yaml:"url"`             // The URL of the GitLab server/api
	Token         string `yaml:"token"`           // Token to use to authenticate against the API
	SkipTLSVerify bool   `yaml:"skip_tls_verify"` // Whether to validate TLS certificates or not

	//ProjectsPollingIntervalSeconds     int        `yaml:"projects_polling_interval_seconds"`      // Interval in seconds at which to poll projects from wildcards
	//RefsPollingIntervalSeconds         int        `yaml:"refs_polling_interval_seconds"`          // Interval in seconds to fetch refs from projects
	//PipelinesPollingIntervalSeconds    int        `yaml:"pipelines_polling_interval_seconds"`     // Interval in seconds to get new pipelines from refs (exponentially backing of to maximum value)
	//PipelinesMaxPollingIntervalSeconds int        `yaml:"pipelines_max_polling_interval_seconds"` // Maximum interval in seconds to fetch new pipelines from refs
	//DefaultRefsRegexp                  string     `yaml:"default_refs"`                           // Default regular expression
	//Projects                           []Project  `yaml:"projects"`                               // List of projects to poll
	//Wildcards                          []Wildcard `yaml:"wildcards"`                              // List of wildcards to search projects from
}

// Default values
const (
	defaultProjectsPollingIntervalSeconds     = 1800
	defaultRefsPollingIntervalSeconds         = 300
	defaultPipelinesPollingIntervalSeconds    = 30
	defaultPipelinesMaxPollingIntervalSeconds = 3600
)

// Parse loads a yaml file into a Config structure
func (cfg *Config) Parse(path string) error {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Couldn't open config file : %s", err.Error())
	}

	fmt.Println(reflect.TypeOf(cfg))
	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		return fmt.Errorf("Unable to parse config file: %s", err.Error())
	}
	/*
		if len(cfg.Projects) < 1 && len(cfg.Wildcards) < 1 {
			return fmt.Errorf("You need to configure at least one project/wildcard to poll, none given")
		}

		// Defining defaults polling intervals
		if cfg.ProjectsPollingIntervalSeconds == 0 {
			cfg.ProjectsPollingIntervalSeconds = defaultProjectsPollingIntervalSeconds
		}

		if cfg.RefsPollingIntervalSeconds == 0 {
			cfg.RefsPollingIntervalSeconds = defaultRefsPollingIntervalSeconds
		}

		if cfg.PipelinesPollingIntervalSeconds == 0 {
			cfg.PipelinesPollingIntervalSeconds = defaultPipelinesPollingIntervalSeconds
		}

		if cfg.PipelinesMaxPollingIntervalSeconds == 0 {
			cfg.PipelinesMaxPollingIntervalSeconds = defaultPipelinesMaxPollingIntervalSeconds
		}*/

	return nil
}
