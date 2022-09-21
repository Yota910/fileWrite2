package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("Yota")
    want := "Hi, Yota. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}
