package main

import (
	dwlog "github.com/agungdhewe/dwlog/src"
)

func main() {
	dlog := dwlog.New()

	dlog.Info("ini adalah info")
	dlog.Log("coba screen logging script")
	dlog.Warning("ini adalah warning")
	dlog.Error("ini log untuk error")
}
