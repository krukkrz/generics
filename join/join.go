package join

import (
	"fmt"
	"strings"
)

func Join[T fmt.Stringer](elements []T, separator string) string {
	var builder strings.Builder
	for i, element := range elements {
		if i > 0 {
			builder.WriteString(separator)
		}
		builder.WriteString(element.String())
	}
	return builder.String()
}
