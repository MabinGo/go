// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !math_big_pure_go

package big

// implemented in arith_$GOARCH.s
func mulWW(x, y Word) (z1, z0 Word)
func divWW(x1, x0, y Word) (q, r Word)
func addVV(z, x, y []Word) (c Word)
func subVV(z, x, y []Word) (c Word)
func addVW(z, x []Word, y Word) (c Word)
func subVW(z, x []Word, y Word) (c Word)
func shlVU(z, x []Word, s uint) (c Word)
func shrVU(z, x []Word, s uint) (c Word)
func mulAddVWW(z, x []Word, y, r Word) (c Word)
func addMulVVW(z, x []Word, y Word) (c Word)
func divWVW(z []Word, xn Word, x []Word, y Word) (r Word)

//TODO:implemented divWWByInv and divWVWByInv in arith_$GOARCH.s
func divWWByInv(x1, x0, y Word, inv uint, shift uint) (q, r Word) {
	return divWWByInv_g(x1, x0, y, inv, shift)
}
func divWVWByInv(z []Word, xn Word, x []Word, y Word) (r Word) {
	return divWVW_g(z, xn, x, y)
}
