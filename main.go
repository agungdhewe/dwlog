package main

import (
	dwlog "github.com/agungdhewe/dwlog/src"
)

func main() {
	dwlog.New()

	dwlog.Info("ini adalah info")
	dwlog.Log("coba screen logging script")
	dwlog.Warning("ini adalah warning")
	dwlog.Error("ini log untuk error")
}
