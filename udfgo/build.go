package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"text/template"
)

type cmdBuild struct{}

type templateData struct {
	ModulePath    string
	Registrations []*udfRegistration
}

type udfRegistration struct {
	Name string
}

func (c *cmdBuild) Execute(args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	pkgs, err := parser.ParseDir(token.NewFileSet(), wd, nil, parser.AllErrors)
	if err != nil {
		return err
	}
	if len(pkgs) == 0 {
		return errors.New("no package detected")
	}
	if len(pkgs) > 1 {
		return errors.New("multiple packages detected")
	}

	var templateData templateData

	// find package import path
	ipkg, err := build.ImportDir(wd, build.FindOnly)
	if err != nil {
		return fmt.Errorf("can not find import path: %v", err)
	}
	if ipkg.ImportPath == "." {
		return errors.New("package is not located in GOPATH")
	}
	templateData.ModulePath = ipkg.ImportPath

	// find exported function declerations (UDF's)
	var name string
	var pkg *ast.Package
	for key, value := range pkgs {
		name = key
		pkg = value
	}

	if name == "main" {
		return errors.New("package cannot be main")
	}

	verbosef("detected package %s\n", name)

	for _, file := range pkg.Files {
		verbosef("scanning file %s\n", file.Name.String())
		for _, decl := range file.Decls {
			funcDecl, isfunc := decl.(*ast.FuncDecl)
			if !isfunc {
				continue
			}
			if !funcDecl.Name.IsExported() {
				continue
			}
			verbosef("found exported func %s\n", funcDecl.Name.String())
			templateData.Registrations = append(templateData.Registrations, &udfRegistration{
				Name: funcDecl.Name.String(),
			})
		}
	}

	tmplRegistry := template.Must(template.New("registry").Parse(boxTemplates.MustString("registry.go.tmpl")))
	err = tmplRegistry.Execute(os.Stdout, &templateData)
	if err != nil {
		return err
	}

	return nil
}
