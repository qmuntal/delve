package native

const (
	_CONTEXT_ARM64               = 0x00400000
	_CONTEXT_CONTROL             = (_CONTEXT_ARM64 | 0x1)
	_CONTEXT_INTEGER             = (_CONTEXT_ARM64 | 0x2)
	_CONTEXT_FLOATING_POINT      = (_CONTEXT_ARM64 | 0x4)
	_CONTEXT_DEBUG_REGISTERS     = (_CONTEXT_ARM64 | 0x8)
	_CONTEXT_ARM64_X18           = (_CONTEXT_ARM64 | 0x10)
	_CONTEXT_FULL                = (_CONTEXT_CONTROL | _CONTEXT_INTEGER | _CONTEXT_FLOATING_POINT)
	_CONTEXT_ALL                 = (_CONTEXT_CONTROL | _CONTEXT_INTEGER | _CONTEXT_FLOATING_POINT | _CONTEXT_DEBUG_REGISTERS | _CONTEXT_ARM64_X18)
	_CONTEXT_EXCEPTION_ACTIVE    = 0x8000000
	_CONTEXT_SERVICE_ACTIVE      = 0x10000000
	_CONTEXT_EXCEPTION_REQUEST   = 0x40000000
	_CONTEXT_EXCEPTION_REPORTING = 0x80000000
)
