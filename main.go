package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"text/template"
)

const (
	templateFileName = ".gb-sublime-project.template"
	defaultTemplate  = `{
    "folders":
    [
        {
            "path": "{{ .ProjectDir }}"
        }
    ],
    "settings": {
        "GoSublime": {
            "env": {
                "GOPATH": "{{ .ProjectDir }}:{{ .ProjectDir }}/vendor"
            }
        }
    }
}`
)

var (
	writeDefaultTemplate = flag.Bool("write-default", false, "writes default template to ~/"+templateFileName)
)

func main() {
	flag.Parse()
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fullTemplatePath := path.Join(u.HomeDir, templateFileName)

	if *writeDefaultTemplate {
		out, err := os.Create(fullTemplatePath)
		if err != nil {
			panic(fmt.Errorf("Error opening template file for writing: %s", err.Error()))
		}
		out.WriteString(defaultTemplate)
		out.Close()
	}

	fmt.Printf("writing sublime-project to %s...\n", wd)

	f, err := os.Open(fullTemplatePath)
	if err != nil {
		panic(fmt.Sprintf("Error opening template file: %s", err.Error()))
	}

	tmpl := template.Must(template.New("project").Parse())

	projectName := filepath.Base(wd)

	out, err := os.Create(projectName + ".sublime-project")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = tmpl.Execute(out, struct{ ProjectDir string }{wd})
	if err != nil {
		panic(err)
	}

}
