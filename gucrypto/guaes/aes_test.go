package guaes

import (
	"testing"
)

func TestEncryptDir(t *testing.T) {

	// The key length can be 32, 24, 16  bytes // in bits: 128, 192 or 256)
	EncryptDir("/tmp/brol", "/tmp/out.zip", "/tmp/out.zip.encrypted", "01234567890123456789012345678901")

}
