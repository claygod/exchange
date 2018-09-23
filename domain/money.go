package domain

// Exchange
// Money
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Money - 32-bit uint

	- ID (24-bit)
	- depth (8-bit)
*/

import (
	"unsafe"
)

type Money uint32

func NewMoney(amount uint32, depth uint8) Money {
	return Money(amount<<8) + Money(depth)
}

func (m Money) Amount() uint32 {
	return uint32(m >> 8)
}
func (m Money) Depth() uint8 {
	return uint8(m)
}

/*
TwoMoney - 64-bit uint

	- MoneyA (32-bit)
	- MoneyB (32-bit)
*/
type TwoMoney uint64

func NewTwoMoney(a Money, b Money) TwoMoney {
	return (TwoMoney(a) << 32) + TwoMoney(b)
}

func (t TwoMoney) GetMoneyA() Money {
	return Money(t >> 32)
}
func (t TwoMoney) GetMoneyB() Money {
	return Money(t)
}
func (t *TwoMoney) SetMoneyA(m Money) {
	*(*Money)(unsafe.Pointer(t)) = m
	//uintptr(unsafe.Pointer((*iface)(unsafe.Pointer(&item)).data))
	// return *(*string) (unsafe.Pointer(uintptr(unsafe.Pointer((*iface)(unsafe.Pointer(&item)).data)) + s.offsetId))
	t2 := TwoMoney((uint64(m) << 32) + uint64(uint32(*t)))
	t = &t2
}
func (t *TwoMoney) SetMoneyB(m Money) {
	*(*Money)(unsafe.Pointer(uintptr(unsafe.Pointer(t)) + 32)) = m
	t2 := TwoMoney((uint64(m) << 32) + uint64(uint32(*t)))
	t = &t2
}

// --

type FinData64 uint64

func (f FinData64) GetData() uint64 {
	return uint64(f >> 8)
}
func (f FinData64) GetDepth() uint8 {
	return uint8(f)
}
func (f FinData64) GetPartA() uint32 {
	return uint32(f >> 32)
}
func (f *FinData64) GetPartAunsafe() uint32 {
	return (uint32)(uintptr(unsafe.Pointer(f)))
	// return uint32(f >> 32)
}
func (f FinData64) GetPartB() uint32 {
	return uint32(f)
}
func (f FinData64) GetPartAData() uint32 {
	return uint32(f>>32 + 8)
}
func (f FinData64) GetPartBData() uint32 {
	return uint32(f) >> 8
}
func (f FinData64) GetPartADepth() uint8 {
	return uint8(uint32(f >> 32))
}
func (f FinData64) GetPartBDepth() uint8 {
	return uint8(f)
}

type FinData32 uint32

func (f FinData32) GetData() uint32 {
	return uint32(f >> 8)
}
func (f FinData32) GetDepth() uint8 {
	return uint8(f)
}
func (f FinData32) GetPartA() uint16 {
	return uint16(f >> 16)
}
func (f FinData32) GetPartB() uint16 {
	return uint16(f)
}
