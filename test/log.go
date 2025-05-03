package main

import (
	"fmt"
	"log"

	"github.com/seipan/trigger"
)

type Config struct {
	Host string
	Port int
}

func main() {

	makeErr := func(field string) func(*Config) error {
		return func(_ *Config) error {
			log.Printf("⚠️ %s is not set", field)
			return fmt.Errorf("%s is not set", field)
		}
	}

	flow := trigger.NewFlow[*Config]().AbortOnError(false)
	flow.Add("Step 1", func(c *Config) bool {
		return c.Host != ""
	}, func(c *Config) error {
		return nil
	})
}
