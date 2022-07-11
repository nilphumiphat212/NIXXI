package interpreter

import (
	"os"
	"testing"
)

func TestNewInterpreter(t *testing.T) {
	os.Args = []string{"main.go", "../../test.nixxi"}
	err := NewInterpreter()
	if err != nil {
		t.Error("Can not get command line arguments")
	}
	os.Remove("language_token.json")

	// TODO: can not mock Stdin

	// os.Args = []string{}
	// tempFile, err := ioutil.TempFile("test_dir", "test")
	// if err != nil {
	// 	t.Error("Can not mock stdin")
	// }

	// input := []byte("exit\n")

	// write, err := tempFile.Write(input)
	// if err != nil {
	// 	t.Error("Can not mock stdin (Can not write temp file status : " + strconv.Itoa(write))
	// }

	// seek, err := tempFile.Seek(0, 0)
	// if err != nil {
	// 	t.Error("Can not mock stdin (Can not write seek status : " + strconv.Itoa(int(seek)))
	// }

	// os.Stdin = tempFile

	// err2 := NewInterpreter()
	// if err2 != nil {
	// 	t.Error("Can not open interactive interpreter when args empty")
	// }
}
