package math

import (
	"testing"
)

func TestMod_Mul(t *testing.T) {
	const (
		a = 111111111
		b = 123456789
		c = 987654321
	)
	if r := ModMul(ModMul(a, b), c); r != 769682799 {
		t.Fatal(r)
	}
}

func TestMod_Add(t *testing.T) {
	const (
		a = 1e9 + 7
		b = 10
	)
	if r := ModAdd(a, b); r != 10 {
		t.Fatal(r)
	}
}

func TestMod_Sub(t *testing.T) {
	const (
		a = 2000000020
		b = 20
	)
	if r := ModSub(a, b); r != 999999993 {
		t.Fatal(r)
	}
}

func TestMod_Inv(t *testing.T) {
	const (
		a = 12345678900000
		b = 100000
	)
	if r := ModDiv(a, b); r != 123456789 {
		t.Fatal(r)
	}
}

func TestMod_Pow(t *testing.T) {
	const (
		a = 2
		b = 10 * 5
	)
	if r := ModPow(a, b); r != 123456789 {
		t.Fatal(r)
	}
}
