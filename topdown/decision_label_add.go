package topdown

import (
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/topdown/builtins"
)

// add init() in here
// In init() register the function
// http.go Line 201 as example
func init() {
	RegisterBuiltinFunc(ast.DecisionLabelAdd.Name, builtinDecisionLabelAdd)
}

// Operands (http.go Line 128) reference the Decl field from the Builtin Struct and are insanely useful in here
func builtinDecisionLabelAdd(bctx BuiltinContext, operands []*ast.Term, _ func(*ast.Term) error) error {

	// validate the operands as Strings
	//   key as String
	//   value as String representation of JSON
	key, err := validateKeyStringOperand(operands[0], 1)
	if err != nil {
		return handleBuiltinErr(ast.DecisionLabelAdd.Name, bctx.Location, err)
	} // end if
	value, err := validateValueStringOperand(operands[1], 2)
	if err != nil {
		return handleBuiltinErr(ast.DecisionLabelAdd.Name, bctx.Location, err)
	} // end if

	// use validated operands to assign the DecisionLabel field of the BuiltinContext
	return assignOperandsToDecisionLabel(bctx, key, value)

}

func validateKeyStringOperand(term *ast.Term, pos int) (ast.String, error) {

	keyStr, err := builtins.StringOperand(term.Value, pos)
	if err != nil {
		return "", err // nil cannot be used as return value for ast.String; used empty String
	}

	return keyStr, nil

}

func validateValueStringOperand(term *ast.Term, pos int) (ast.String, error) {

	valueStr, err := builtins.StringOperand(term.Value, pos)
	if err != nil {
		return "", err // nil cannot be used as return value for ast.String; used empty String
	}

	return valueStr, nil

}

func assignOperandsToDecisionLabel(bctx BuiltinContext, key, value ast.String) error {

	bctx.DecisionLabel.Add(key, value)

	if _, ok := bctx.DecisionLabel.Get(key); !ok {
		return ast.NewError(InternalErr, bctx.Location, "Entry for %s was not added", key.String())
	}

	return nil

}
