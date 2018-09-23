package domain

// Exchange
// FinData bench
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"testing"
)

func BenchmarkUint64partA(b *testing.B) {
	b.StopTimer()

	var i2 uint64
	var i3 uint32

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		i2 = uint64(i)
		i3 = uint32(i2)
	}
	_ = i2
	_ = i3
}

func BenchmarkFinData64partA(b *testing.B) {
	b.StopTimer()

	var i2 FinData64
	var i3 FinData32

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		i2 = FinData64(i)
		i3 = FinData32(i2)
	}
	_ = i2
	_ = i3
}

func BenchmarkFinData64GetPartA(b *testing.B) {
	b.StopTimer()

	var i2 FinData64
	var i3 uint32

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		i2 = FinData64(i)
		i3 = i2.GetPartA()
	}
	_ = i2
	_ = i3
}

func BenchmarkFinData64GetPartAunsafe(b *testing.B) {
	b.StopTimer()

	var i2 FinData64
	var i3 uint32

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		i2 = FinData64(i)
		i3 = i2.GetPartAunsafe()
	}
	_ = i2
	_ = i3
}

func BenchmarkFinData64GetData(b *testing.B) {
	b.StopTimer()

	var i2 FinData64
	var i3 uint64

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		i2 = FinData64(i)
		i3 = i2.GetData()
	}
	_ = i2
	_ = i3
}
