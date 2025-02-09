package parser

import (
	"go/ast"
	"go/token"
)

func AddImport(file *ast.File, packageName string, filename string) {
	if file.Name.Name != "main" {
		return
	}
	print("AddImport ", filename, "\n")
	noImport := true
	toInsert := &ast.ImportSpec{
		Name: &ast.Ident{
			Name: "_",
		},
		Path: &ast.BasicLit{
			ValuePos: 0,
			Kind:     token.STRING,
			Value:    packageName,
		},
		EndPos: 0,
	}
	for _, decl := range file.Decls {
		fd, ok := decl.(*ast.GenDecl)
		if ok && fd.Tok == token.IMPORT {
			imports := make([]ast.Spec, 0, len(fd.Specs)+1)
			imports = append(imports, toInsert)
			imports = append(imports, fd.Specs...)
			fd.Specs = imports
			noImport = false
		}
	}
	if noImport {
		decls := make([]ast.Decl, 0, len(file.Decls)+1)
		imports := make([]ast.Spec, 0, 1)
		imports = append(imports, toInsert)
		decl := &ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: imports,
		}
		decls = append(decls, decl)
		decls = append(decls, file.Decls...)
		file.Decls = decls
	}
	// ast.Print(fset, file)

	//var cfg printer.Config
	//var buf bytes.Buffer
	//
	//cfg.Fprint(&buf, fset, file)
	//fmt.Printf(buf.String())
	// print("add import here")
	// for _, decl := range file.Decls {
	// 	fd, ok := decl.(*ast.FuncDecl)
	// 	if ok && fd.Name.Name == "main" {
	// 		for _, _stmt := range fd.Body.List{
	// 			_gostmt, ok := _stmt.(*ast.GoStmt)
	// 			if ok {
	// 				print("go stmt here",_gostmt)
	// 			}
	// 		}
	// 	}
	// }
}
