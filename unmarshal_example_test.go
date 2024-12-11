package env_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"rodusek.dev/pkg/env"
)

func ExampleUnmarshal() {
	type Environment struct {
		ProjectName string        `env:"PROJECT_NAME,required"`
		Timeout     time.Duration `env:"TIMEOUT"`
		Path        []string      `env:"PATH,required,sep=;"`
	}

	os.Setenv("PROJECT_NAME", "example")
	environment := Environment{
		Timeout: 5 * time.Second, // Setting a default for if not specified
	}
	if err := env.Unmarshal(&environment); err != nil {
		log.Fatalf("failed to unmarshal environment: %v", err)
	}
	fmt.Println(environment.ProjectName, environment.Timeout)
	// Output: example 5s
}
