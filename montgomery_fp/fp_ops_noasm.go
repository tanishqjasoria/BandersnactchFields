//go:build !amd64
// +build !amd64

// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by consensys/gnark-crypto DO NOT EDIT

package main

// MulBy3 x *= 3 (mod q)
func MulBy3(x *fp) {
	_x := *x
	x.Double(x).Add(x, &_x)
}

// MulBy5 x *= 5 (mod q)
func MulBy5(x *fp) {
	_x := *x
	x.Double(x).Double(x).Add(x, &_x)
}

// MulBy13 x *= 13 (mod q)
func MulBy13(x *fp) {
	var y = fp{
		120259084260,
		15510977298029211676,
		7326335280343703402,
		5909200893219589146,
	}
	x.Mul(x, &y)
}

// Butterfly sets
//
//	a = a + b (mod q)
//	b = a - b (mod q)
func Butterfly(a, b *fp) {
	_butterflyGeneric(a, b)
}
func mul(z, x, y *fp) {
	_mulGeneric(z, x, y)
}

func fromMont(z *fp) {
	_fromMontGeneric(z)
}

func reduce(z *fp) {
	_reduceGeneric(z)
}
