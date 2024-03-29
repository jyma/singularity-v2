package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"
	"text/template"
	"unicode"

	"github.com/data-preservation-programs/singularity/cmd/datasource"
	"github.com/urfave/cli/v2"
)

const header = `
//lint:file-ignore U1000 Ignore all unused code, it's generated
// Code generated. DO NOT EDIT.
package datasource

import (
	"github.com/data-preservation-programs/singularity/model"
)
`

const structTemplate = `
type {{.Name}} struct {
    SourcePath string ` + "`validate:\"required\" json:\"sourcePath\"`" + `// The path of the source to scan items
    DeleteAfterExport bool ` + "`validate:\"required\" json:\"deleteAfterExport\"`" + `// Delete the source after exporting to CAR files
    RescanInterval string ` + "`validate:\"required\" json:\"rescanInterval\"`" + `// Automatically rescan the source directory when this interval has passed from last successful scan
    ScanningState model.WorkState ` + "`validate:\"required\" json:\"scanningState\"`" + `// Starting state for scanning
		{{- range .Fields }}
    {{.Name}} {{.Type}} {{.Tag}} // {{.Description}}
	{{- end }}
}
`

const allStructTemplate = `
type {{.Name}} struct {
		SourcePath string ` + "`validate:\"required\" json:\"sourcePath\"`" + `// The path of the source to scan items
    DeleteAfterExport bool ` + "`validate:\"optional\" json:\"deleteAfterExport\"`" + `// Delete the source after exporting to CAR files
    RescanInterval string ` + "`validate:\"optional\" json:\"rescanInterval\"`" + `// Automatically rescan the source directory when this interval has passed from last successful scan
    ScanningState model.WorkState ` + "`validate:\"optional\" json:\"scanningState\"`" + `// Starting state for scanning
		{{- range .Fields }}
    {{.Name}} {{.Type}} {{.Tag}} // {{.Description}}
	{{- end }}
}
`

const handlerTemplate = `
// @Summary Add {{.Name}} source for a dataset
// @Tags Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {object} model.Source
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body {{.StructName}} true "Request body"
// @Router /source/{{.Name}}/dataset/{datasetName} [post]
func {{.FuncName}}() {}

`

type Func struct {
	FuncName   string
	Name       string
	StructName string
}

type Field struct {
	Name        string
	Type        string
	Tag         string
	Description string
}

type Type struct {
	Name   string
	Fields []Field
}

func main() {
	command := datasource.AddCmd

	f := bytes.NewBuffer(nil)

	// Write package lines
	f.WriteString(header)

	tmpl, err := template.New("handler").Parse(handlerTemplate)
	if err != nil {
		panic(err)
	}

	structTemplate, err := template.New("struct").Parse(structTemplate)
	if err != nil {
		panic(err)
	}

	allStructTemplate, err := template.New("struct").Parse(allStructTemplate)
	if err != nil {
		panic(err)
	}

	all := Type{
		Name:   "AllConfig",
		Fields: []Field{},
	}

	for _, cmd := range command.Subcommands {
		var fields []Field
		for _, flag := range cmd.Flags {
			flagName := strings.SplitN(flag.Names()[0], "-", 2)[1]
			name := argNameToCamel(flagName)
			allName := cmd.Name + name
			snake := lowerFirst(name)
			allSnake := lowerFirst(allName)
			stringFlag, ok := flag.(*cli.StringFlag)
			if !ok {
				continue
			}
			tag := fmt.Sprintf("`json:\"%s\"`", snake)
			allTag := fmt.Sprintf("`json:\"%s\"`", allSnake)
			if stringFlag.Value != "" {
				tag = fmt.Sprintf("`json:\"%s\" default:\"%s\"`", snake, stringFlag.Value)
				allTag = fmt.Sprintf("`json:\"%s\" default:\"%s\"`", allSnake, stringFlag.Value)
			}
			if name == "2fa" {
				name = "TwoFA"
			}
			fields = append(fields, Field{
				Name:        name,
				Type:        "string",
				Tag:         tag,
				Description: stringFlag.Usage,
			})
			all.Fields = append(all.Fields, Field{
				Name:        capitalizeFirst(allName),
				Type:        "string",
				Tag:         allTag,
				Description: stringFlag.Usage,
			})
		}
		t := Type{
			Name:   capitalizeFirst(cmd.Name) + "Request",
			Fields: fields,
		}
		err = structTemplate.Execute(f, t)
		if err != nil {
			panic(err)
		}

		fobj := Func{
			FuncName:   "handle" + capitalizeFirst(cmd.Name),
			Name:       cmd.Name,
			StructName: capitalizeFirst(cmd.Name) + "Request",
		}

		err = tmpl.Execute(f, fobj)
		if err != nil {
			panic(err)
		}
	}

	err = allStructTemplate.Execute(f, all)
	if err != nil {
		panic(err)
	}

	formatted, err := format.Source(f.Bytes())
	if err != nil {
		panic(err)
	}
	// Create generated file in the same directory
	err = os.WriteFile("handler/datasource/add_gen.go", formatted, 0644)
	if err != nil {
		panic(err)
	}
}

func argNameToCamel(s string) string {
	parts := strings.Split(s, "-")
	for i := 0; i < len(parts); i++ {
		r := []rune(parts[i])
		r[0] = unicode.ToUpper(r[0])
		parts[i] = string(r)
	}
	return strings.Join(parts, "")
}
func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}
