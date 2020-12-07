package env

import (
	"os"
	"reflect"
	"testing"
)

func Test_NewArgs(t *testing.T) {
	myArgs := []string{"SouthWest", "NorthEast"}
	os.Args = myArgs

	args := NewArgs()

	if reflect.DeepEqual(args.arguments, &myArgs) {
		t.Error("args not equal", args, myArgs)
	}
}
