package javadoc

import (
	"fmt"
	"strings"

	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
)

type Func struct {
	ReturnType string
	Parameters []string
}

func NewFunc(ret string, params ...string) Func {
	p := make([]string, len(params))
	for i, par := range params {
		p[i] = strings.TrimSpace(par)
	}
	return Func{
		ReturnType: strings.TrimSpace(ret),
		Parameters: p,
	}
}

func parseFuncSigantureDirective(input string) (Func, error) {
	parts := strings.Split(input, ",")
	if len(parts) < 1 {
		return Func{}, fmt.Errorf("signature directive needs to contain at least one type")
	}
	return NewFunc(parts[0], parts[1:]...), nil
}

func getFuncSignatureString(input []string) (string, error) {
	directive := strings.Join(input, ",")
	sign, err := parseFuncSigantureDirective(directive)
	if err == nil {
		return sign.String(), nil
	}
	return "", err
}

func (fs Func) String() string {
	pars := fs.Parameters
	for i, p := range pars {
		pars[i] = "var " + p
	}
	return "func " + fs.ReturnType + " [](" + strings.Join(pars, ",") + ")"

}

func (fs Func) EqualSym(other symbol.Function) bool {
	if !strings.EqualFold(fs.ReturnType, other.ReturnType) {
		return false
	}
	if len(fs.Parameters) != len(other.Parameters) {
		return false
	}
	for index, p := range fs.Parameters {
		if !strings.EqualFold(p, other.Parameters[index].Type) {
			return false
		}
	}
	return true
}
func (fs Func) Equal(other Func) bool {
	if !strings.EqualFold(fs.ReturnType, other.ReturnType) {
		return false
	}
	if len(fs.Parameters) != len(other.Parameters) {
		return false
	}
	for index, p := range fs.Parameters {
		if !strings.EqualFold(p, other.Parameters[index]) {
			return false
		}
	}
	return true
}
