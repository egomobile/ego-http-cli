package main

import (
	cmd "github.com/egomobile/ego-http-cli/cmd"

	"github.com/thatisuday/commando"
)

// Top level configuration of the ego-http-cli tool.
const description = "The ego-http is an invaluable Command Line Interface (CLI) tool that offers a wide range of useful functionalities, including but not limited to scaffolding endpoints, generating tests, and creating OpenAPI documentation, thereby streamlining the development process and enhancing the overall efficiency of the workflow."
const name = "ego-http"
const version = "0.1.0"

func main() {
	// setup commando
	commando.
		SetExecutableName(name).
		SetVersion(version).
		SetDescription(description)

	cmd.Execute()

	// parse CLI args
	commando.Parse(nil)
}
