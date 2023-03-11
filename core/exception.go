package core

import (
	"fmt"
	"runtime/debug"

	"github.com/GoFeGroup/gordp/glog"
)

func TryCatch(f func(), catch func(e any)) {
	defer func() {
		if r := recover(); r != nil {
			catch(r)
		}
	}()
	f()
}

func toError(e any) error {
	if e == nil {
		return nil
	}
	if err, ok := e.(error); ok {
		return err
	}
	return fmt.Errorf("%v", e)
}

func Try(f func()) error {
	var ret error
	TryCatch(f, func(e any) {
		ret = toError(e)
	})
	return ret
}

func Throw(e any) {
	glog.Debugf("%v", e)
	glog.Debugf("%s", debug.Stack())
	panic(e)
}

func Throwf(format string, args ...any) {
	str := fmt.Sprintf(format, args)
	Throw(str)
}

func ThrowNil(e any) {
	if e != nil {
		Throw(e)
	}
}

func ThrowError(e any) {
	ThrowNil(e)
}

func ThrowIf(cond bool, e any) {
	if cond {
		Throw(e)
	}
}
