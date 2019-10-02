package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Russell")
	if s != "Welcome my dear, Russell" {
		t.Error("got", s, "expected", "Welcome my dear, Russell")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Russell"))
	// Output:
	// Welcome my dear, Russell
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Russell")
	}
}
