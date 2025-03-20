package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "expecterlint",
	Doc:  "check if mock.On(...) could be replaced with mock.EXPECT()",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	inspect := func(node ast.Node) bool {
		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok || selectorExpr.Sel.Name != "On" {
			return true
		}

		methodName := firstArg(callExpr)
		if methodName == "" {
			return false
		}

		ident, ok := selectorExpr.X.(*ast.Ident)
		if !ok {
			return false
		}

		pointer, ok := pass.TypesInfo.ObjectOf(ident).Type().(*types.Pointer)
		if !ok {
			return false
		}

		named, ok := pointer.Elem().(*types.Named)
		if !ok {
			return false
		}

		if !hasExpect(named, methodName) {
			return false
		}

		pass.Report(analysis.Diagnostic{
			Pos:     callExpr.Pos(),
			End:     callExpr.End(),
			Message: fmt.Sprintf("mock.On(%[1]q, ...) could be replaced with mock.EXPECT().%[1]s(...)", methodName),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "Replace mock.On with mock.EXPECT",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     callExpr.Pos(),
							End:     callExpr.End(),
							NewText: replacement(pass.Fset, ident.Name, methodName, callExpr.Args),
						},
					},
				},
			},
		})

		return false
	}

	for _, file := range pass.Files {
		filename := pass.Fset.File(file.Pos()).Name()
		if strings.HasSuffix(filename, "_test.go") {
			ast.Inspect(file, inspect)
		}
	}

	return nil, nil
}

func replacement(fSet *token.FileSet, mockName string, methodName string, mockArgs []ast.Expr) []byte {
	buf := bytes.NewBufferString(mockName + ".EXPECT()." + methodName + "(")

	for i, arg := range mockArgs[1:] {
		if i > 0 {
			buf.WriteString(", ")
		}
		_ = printer.Fprint(buf, fSet, arg)
	}
	buf.WriteString(")")
	return buf.Bytes()
}

func firstArg(expr *ast.CallExpr) string {
	arg1, ok := expr.Args[0].(*ast.BasicLit)
	if !ok || arg1.Kind != token.STRING || len(arg1.Value) < 3 {
		return ""
	}
	return arg1.Value[1 : len(arg1.Value)-1]
}

// hasExpect checks if instead of .On("MethodName", ...) there is callable .EXPECT().MethodName(...)
func hasExpect(named *types.Named, methodName string) bool {
	for i := range named.NumMethods() {
		if named.Method(i).Name() == "EXPECT" && expectHasMethod(named.Method(i), methodName) {
			return true
		}
	}
	return false
}

func expectHasMethod(method *types.Func, methodName string) bool {
	pointer, ok := method.Signature().Results().At(0).Type().(*types.Pointer)
	if !ok {
		return false
	}

	named, ok := pointer.Elem().(*types.Named)
	if !ok {
		return false
	}

	for i := range named.NumMethods() {
		if named.Method(i).Name() == methodName {
			return true
		}
	}

	return false
}
