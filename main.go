package main

import (
	"runtime"
	"somegame/objects"
)

func main() {
	runtime.LockOSThread()
	objects.CreateGame()
}
