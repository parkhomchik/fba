// go build gen/* && ./codegen.exe pack/unpack.go  pack/marshaller.go
// go run pack/*

/*

Что нужно сделать

-	генерация работы с БД		+
-	генерация работы с http		-
- 	main.go REST 				-
*/

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"text/template"
)

type tpl struct {
	Name string
}

var tmpl = template.Must(template.New("tmpl").Parse(`
package db

import "fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) {{.Name}}Create(c model.{{.Name}}) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) {{.Name}}Update(c model.{{.Name}}) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) {{.Name}}Delete(c model.{{.Name}}) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) {{.Name}}Get(size, page int) (citys []model.{{.Name}}, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) {{.Name}}GetById(id uuid.UUID) (city model.{{.Name}}, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) {{.Name}}Count(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.{{.Name}}{}).Count(&count).Error
	return
}
`))

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range node.Decls {
		g, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range g.Specs {
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			_, ok = currType.Type.(*ast.StructType)
			if ok {
				out, _ := os.Create(os.Args[2] + "/" + currType.Name.Name + "Resource.go")
				tmpl.Execute(out, tpl{currType.Name.Name})
				continue
			}

		}
	}
}

// go build gen/* && ./codegen.exe pack/unpack.go  pack/marshaller.go
// go run pack/*
