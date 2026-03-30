package types

import (
	"fmt"
	"strings"
)

// AddIndent adds indentation to the code.
func AddIndent(code string, numOfIndentations int) string {
	indent := strings.Repeat("\t", numOfIndentations)
	return indent + strings.ReplaceAll(code, "\n", "\n"+indent)
}

// Var generates unique variable name with provided prefix.
func Var(prefix string, varIndex *uint64) string {
	*varIndex++
	return fmt.Sprintf("%s%d", prefix, *varIndex)
}
