package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello, World!"
	temp := Hello{"Hello", "World", "!"}
	if got := temp.String(); got != want {
		t.Errorf("Hello.String() = %q, want = %q", got, want)
	}
}

func TestGoodByeWorld(t *testing.T) {
	want := "GoodBye, World!"
	temp := Hello{"GoodBye", "World", "!"}
	if got := temp.String(); got != want {
		t.Errorf("Hello.String() = %q, want = %q", got, want)
	}
}
