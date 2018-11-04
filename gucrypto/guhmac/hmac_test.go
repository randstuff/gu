package guhmac

import (
	"fmt"
	"testing"
)

func TestGernerateSHA2_HMAC(t *testing.T) {

	fmt.Println(" :: ", GernerateSHA2_HMAC("secret", "my data"))
}
