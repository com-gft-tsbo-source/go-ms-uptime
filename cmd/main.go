package main

import (
	"os"

	msuptime "com.gft.tsbo-training.src.go/ms-uptime/msuptime"
)

// ###########################################################################
// ###########################################################################
// MAIN
// ###########################################################################
// ###########################################################################

func main() {

	var ms msuptime.MsUptime
	msuptime.InitFromArgs(&ms, os.Args, nil)

	// defer ms.db.Close()
	ms.Run()
}
