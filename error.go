package gyro

import (
	"fmt"
	"runtime"
)

type GyroErr struct {
	info string
}

func NewGyroErr(info any) *GyroErr {
	callstack := make([]byte, 1024)
	stackLen := runtime.Stack(callstack, false)
	callstack = callstack[2:stackLen]
	return &GyroErr{info: fmt.Sprintf("%v %v", info, callstack)}
}

func (e *GyroErr) Error() string {
	return e.info
}

type GyroOverflowErr struct {
	*GyroErr
}

func NewGyroOverflowErr(info any) *GyroOverflowErr {
	return &GyroOverflowErr{GyroErr: NewGyroErr(info)}
}

type GyroInvalidErr struct {
	*GyroErr
}

func NewGyroInvalidErr(info any) *GyroOverflowErr {
	return &GyroOverflowErr{GyroErr: NewGyroErr(info)}
}

type GyroFractionalTooLongErr struct {
	*GyroErr
}

func NewGyroFractionalTooLongErr(info any) *GyroFractionalTooLongErr {
	return &GyroFractionalTooLongErr{GyroErr: NewGyroErr(info)}
}

type GyroExponentNANErr struct {
	*GyroErr
}

func NewGyroExponentNANErr(info any) *GyroExponentNANErr {
	return &GyroExponentNANErr{GyroErr: NewGyroErr(info)}
}

type GyroTooManyDecimals struct {
	*GyroErr
}

func NewGyroTooManyDecimals(info any) *GyroTooManyDecimals {
	return &GyroTooManyDecimals{GyroErr: NewGyroErr(info)}
}
