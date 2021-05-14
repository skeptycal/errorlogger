package osargsutils

import (
	"os"
	"path"
	"path/filepath"

	"github.com/skeptycal/errorlogger"
)

var log = errorlogger.Log
var Err = log.Err

// Arg0 returns the absolute path for the executable that started
// the current process. There is no guarantee that the path is
// still pointing to the correct executable. Symlinks are evaluated
// if necessary.
//
// Executable returns an absolute path unless an error occurred.
//
// The main use case is finding resources located relative to an executable.
func Arg0() (string, error) {
	// As of Go 1.8 (Released February 2017) the recommended
	// way of doing this is with os.Executable:
	ex, err := os.Executable()
	if err != nil {
		return "", Err(err)
	}

	return filepath.EvalSymlinks(ex)
}

// HereMe returns the basename (me) and folder (here) of
// the executable that started the current process.
func HereMe() (string, string, error) {
	// As of Go 1.8 (Released February 2017) the recommended
	// way of doing this is with os.Executable:
	ex, err := os.Executable()
	if err != nil {
		return "", "", Err(err)
	}

	zero, err := filepath.EvalSymlinks(ex)
	if err != nil {
		return "", "", Err(err)
	}

	return filepath.Dir(zero), filepath.Base(zero), nil
}

// hereMe2 returns the folder (here) and basename (me) of
// the executable that started the current process.
func hereMe2() (string, string, error) {
	zero, err := zeroOsExecutable()
	if err != nil {
		return "", "", Err(err)
	}

	// TODO - using path.Split() returns dir ending
	// with a slash, where Dir() would not
	dir, base := path.Split(zero)
	return dir, base, nil
}

func zeroOsArgs() (string, error) {
	// Prior to Go 1.8, you could use os.Args[0]
	ex, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", Err(err)
	}

	return filepath.EvalSymlinks(ex)
}

func zeroOsExecutable() (string, error) {
	// As of Go 1.8 (Released February 2017) the recommended
	// way of doing this is with os.Executable:
	ex, err := os.Executable()
	if err != nil {
		return "", Err(err)
	}

	return filepath.EvalSymlinks(ex)
}

func rawOsArgsZero() (string, error) {
	return os.Args[0], nil
}
