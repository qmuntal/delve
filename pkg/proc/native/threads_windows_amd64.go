package native

import (
	"errors"

	"github.com/go-delve/delve/pkg/proc/amd64util"
	"github.com/go-delve/delve/pkg/proc/winutil"
)

func (t *nativeThread) withDebugRegisters(f func(*amd64util.DebugRegisters) error) error {
	if !enableHardwareBreakpoints {
		return errors.New("hardware breakpoints not supported")
	}

	context := winutil.NewCONTEXT()
	context.ContextFlags = _CONTEXT_DEBUG_REGISTERS

	err := _GetThreadContext(t.os.hThread, context)
	if err != nil {
		return err
	}

	drs := amd64util.NewDebugRegisters(&context.Dr0, &context.Dr1, &context.Dr2, &context.Dr3, &context.Dr6, &context.Dr7)

	err = f(drs)
	if err != nil {
		return err
	}

	if drs.Dirty {
		return _SetThreadContext(t.os.hThread, context)
	}

	return nil
}
