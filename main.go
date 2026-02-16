package main

import (
	"os"
	"github.com/ZeljkoBenovic/tsbc/cmd"
)

func main() {
	// Set the environment variable
    os.Setenv("DOCKER_API_VERSION", "1.44")
	cmd.Execute()
}
