package main

import (
	"os"

	msuptime "github.com/com-gft-tsbo-source/go-ms-uptime/msuptime"
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
