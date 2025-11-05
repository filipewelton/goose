package generators

import (
	"fmt"
	"strings"

	"syreclabs.com/go/faker"
)

func GeneratePassword() string {
	return fmt.Sprint(
		strings.ToLower(faker.Lorem().Characters(3)),
		strings.ToUpper(faker.Lorem().Characters(3)),
		faker.Number().Number(1),
		"#",
	)
}
