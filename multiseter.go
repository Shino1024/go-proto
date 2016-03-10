package multiseter

import (
	"fmt"
)

type Lesser interface {
	Less(other Lesser) bool
}

type Multiseter struct {
	data []Lesser
}
