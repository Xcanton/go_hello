package main

// 产量生成器iota，它创建一系列相关值， 而不是逐个值显式地写出。
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type flags uint

const (
	FlagUp flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *flags)     { *v &^= FlagUp }
func SetBroadcast(v *flags) { *v |= FlagBroadcast }
func IsCast(v flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

const (
	_ = 1 << (10 * iota)
	kib
	mib
	gib
	tib
	pib
	eib
	zib
	yib
)

func main() {

}
