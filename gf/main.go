package main

import (
	_ "gf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"gf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
