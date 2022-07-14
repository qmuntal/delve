package winutil

import (
	"fmt"

	"github.com/go-delve/delve/pkg/dwarf/op"
	"github.com/go-delve/delve/pkg/proc"
)

type CONTEXT interface {
	SetReg(regNum uint64, reg *op.DwarfRegister) error
	SetFlags(flags uint32)
	SetTrap(trap bool)
	SetPC(pc uint64)
}

type Registers interface {
	proc.Registers
	Ctx() CONTEXT
}

func NewCONTEXT(goarch string) CONTEXT {
	switch goarch {
	case "amd64":
		return NewAMD64CONTEXT()
	case "arm64":
		return NewARM64CONTEXT()
	default:
		panic(fmt.Errorf("unknown goarch %s", goarch))
	}
}

func NewRegisters(ctx CONTEXT, TebBaseAddress uint64) proc.Registers {
	switch ctx := ctx.(type) {
	case *AMD64CONTEXT:
		return NewAMD64Registers(ctx, TebBaseAddress)
	case *ARM64CONTEXT:
		return NewARM64Registers(ctx, TebBaseAddress)
	default:
		panic(fmt.Errorf("unknown context type %T", ctx))
	}
}
