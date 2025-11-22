package utils

import (
	"fmt"
	"strings"

	"syreclabs.com/go/faker"
)

func GeneratePassword() string {
	return fmt.Sprint(
		faker.Letterify("??"),
		strings.ToLower(faker.Letterify("??")),
		faker.Number().Number(2),
		"@#",
	)
}
