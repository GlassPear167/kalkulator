// math_test.go
package count

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	Count("1+1")
	fmt.Println("buffer in rune", BufferUtama)
	fmt.Println("buffer in string", string(BufferUtama))

}
