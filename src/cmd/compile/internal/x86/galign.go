// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x86

import (
	"cmd/compile/internal/gc"
	"cmd/internal/obj"
	"cmd/internal/obj/x86"
	"fmt"
	"os"
)

func betypeinit() {
}

func Main() {
	gc.Thearch.LinkArch = &x86.Link386
	gc.Thearch.REGSP = x86.REGSP
	gc.Thearch.REGCTXT = x86.REGCTXT
	gc.Thearch.REGCALLX = x86.REG_BX
	gc.Thearch.REGCALLX2 = x86.REG_AX
	gc.Thearch.REGRETURN = x86.REG_AX
	gc.Thearch.REGMIN = x86.REG_AX
	gc.Thearch.REGMAX = x86.REG_DI
	switch v := obj.GO386; v {
	case "387":
		gc.Thearch.FREGMIN = x86.REG_F0
		gc.Thearch.FREGMAX = x86.REG_F7
		gc.Thearch.Use387 = true
	case "sse2":
		gc.Thearch.FREGMIN = x86.REG_X0
		gc.Thearch.FREGMAX = x86.REG_X7
	default:
		fmt.Fprintf(os.Stderr, "unsupported setting GO386=%s\n", v)
		gc.Exit(1)
	}
	gc.Thearch.MAXWIDTH = (1 << 32) - 1
	gc.Thearch.ReservedRegs = resvd

	gc.Thearch.Betypeinit = betypeinit
	gc.Thearch.Defframe = defframe
	gc.Thearch.Gins = gins
	gc.Thearch.Proginfo = proginfo

	gc.Thearch.SSARegToReg = ssaRegToReg
	gc.Thearch.SSAMarkMoves = ssaMarkMoves
	gc.Thearch.SSAGenValue = ssaGenValue
	gc.Thearch.SSAGenBlock = ssaGenBlock

	gc.Main()
	gc.Exit(0)
}
