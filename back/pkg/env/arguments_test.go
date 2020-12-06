package env

import (
	"os"
	"reflect"
	"testing"
)

func Test_NewArgs(t *testing.T) {
	myArgs := []string{"A", "B"}
	os.Args = myArgs

	args := NewArgs()

	if reflect.DeepEqual(args.arguments, &myArgs) {
		t.Error("args not equal", args, myArgs)
	}
}
