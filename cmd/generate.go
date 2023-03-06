// This file is part of the ego-cli distribution.
// Copyright (c) Next.e.GO Mobile SE, Aachen, Germany (https://e-go-mobile.com/)
//
// ego-cli is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, version 3.
//
// ego-cli is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"text/template"

	types "github.com/egomobile/ego-http-cli/cmd/types"
	utils "github.com/egomobile/ego-http-cli/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/thatisuday/commando"
)

// templates directory hold files for the new controller's following parts:
//   - index.ts (file to contain implementation of endpoints)
//   - index.spec.ts (file to contain test specifications of endpoints)
//   - index.openapi.ts (file to contain OpenAPI documentation of endpoints)
//
//go:embed templates/*
var templates embed.FS

// Configuration of the command.
const description = "generates variety of files such as endpoints, repositories, etc"
const name = "generate"
const shortDescription = "generates files based on a schematic"

func Setup_generate_command() {
	commando.
		Register(name).
		SetShortDescription(shortDescription).
		SetDescription(description).
		AddArgument("schematic", "schematic to generate", "").
		AddFlag("name, n", `name of the generated schematic`, commando.String, "index").
		AddFlag("path,p", "run \"audit fix\" after \"install\" or \"update\"", commando.String, "").
		SetAction(run_generate_command)
}

func run_generate_command(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	// get arguments
	schematic := args["schematic"].Value

	// get flags
	name := utils.GetStringFlag(flags, "name", "index")
	path := utils.GetStringFlag(flags, "path", "")

	switch schematic {
	case "controller":
		generate_controller(name, path)
	default:
		utils.ExitForInvalidArg("Unknown schematic", schematic)
	}
}

func generate_controller(name string, path string) {
	controller_base_name := cases.Title(language.English, cases.Compact).String(name)
	controller_directory_path := "./src/app/controllers"
	controller_formatted_name := fmt.Sprintf("%vController", controller_base_name)
	controller_full_path := fmt.Sprintf("%v/%v", controller_directory_path, path)

	if _, err := os.Stat(controller_full_path); !os.IsNotExist(err) {
		log.Fatalf("directory %v already exists!", controller_full_path)
	}

	err := os.MkdirAll(controller_full_path, os.ModePerm)
	if err != nil {
		log.Fatalf("directory %v could not be created! Error = %q", controller_full_path, err)
	}

	createAndWriteImplementationFile(controller_full_path, controller_formatted_name, path)
	createAndWriteSpecFile(controller_full_path, controller_formatted_name, path)
	createAndWriteOpenApiFile(controller_full_path, controller_formatted_name, path)
}

func createAndWriteImplementationFile(controller_full_path string, controller_formatted_name string, path string) {
	f, err := os.Create(fmt.Sprintf("%v/index.ts", controller_full_path))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	templateContent, _ := templates.ReadFile("templates/index.ts")
	templateContentString := string(templateContent)

	content := types.Controller{
		Name:     controller_formatted_name,
		BasePath: path,
		Status:   "200",
	}

	tmpl, err := template.New("controller_template").Parse(templateContentString)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, content)
	if err != nil {
		panic(err)
	}

	f.Write(tpl.Bytes())
}

func createAndWriteSpecFile(controller_full_path string, controller_formatted_name string, path string) {
	f, err := os.Create(fmt.Sprintf("%v/index.spec.ts", controller_full_path))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	templateContent, _ := templates.ReadFile("templates/index.spec.ts")
	templateContentString := string(templateContent)

	content := types.Controller{
		Name:     controller_formatted_name,
		BasePath: path,
		Status:   "200",
	}

	tmpl, err := template.New("controller_template").Parse(templateContentString)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, content)
	if err != nil {
		panic(err)
	}

	f.Write(tpl.Bytes())
}

func createAndWriteOpenApiFile(controller_full_path string, controller_formatted_name string, path string) {
	f, err := os.Create(fmt.Sprintf("%v/index.openapi.ts", controller_full_path))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	templateContent, _ := templates.ReadFile("templates/index.openapi.ts")
	templateContentString := string(templateContent)

	content := types.Controller{
		Name:     controller_formatted_name,
		BasePath: path,
		Status:   "200",
	}

	tmpl, err := template.New("controller_template").Parse(templateContentString)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, content)
	if err != nil {
		panic(err)
	}

	f.Write(tpl.Bytes())
}
