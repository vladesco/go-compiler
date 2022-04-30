package object

import (
	"bytes"
	"compiler/ast"
	"strings"
)

type Function struct {
	Parameters  []*ast.Identifier
	Body        *ast.BlockStatement
	Environment *Environment
}

func (functionObj *Function) GetObjectType() ObjectType { return FUNCTION_OBJ }
func (functionObj *Function) Inspect() string {
	var output bytes.Buffer

	parameters := []string{}

	for _, parameter := range functionObj.Parameters {
		parameters = append(parameters, parameter.ToString())
	}
	output.WriteString("fn")
	output.WriteString("(")
	output.WriteString(strings.Join(parameters, ", "))
	output.WriteString(") {\n")
	output.WriteString(functionObj.Body.ToString())
	output.WriteString("\n}")

	return output.String()
}
