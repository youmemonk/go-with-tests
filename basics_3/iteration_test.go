package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeat != expected {
		t.Errorf("expected %q but got %q", expected, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeat := Repeat("b", 3)
	fmt.Println(repeat)
	//output "bbb"
}
