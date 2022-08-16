package msuptime

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/com-gft-tsbo-source/go-common/ms-framework/microservice"
	"github.com/com-gft-tsbo-source/go-ms-uptime/database"
)

// ###########################################################################
// ###########################################################################
// MsUptime
// ###########################################################################
// ###########################################################################

// MsUptime Encapsulates the ms-uptime data
type MsUptime struct {
	microservice.MicroService
	db        *database.Database
	starttime time.Time
}

// ###########################################################################

// InitMsUptimeFromArgs Constructor of a new ms-uptime
func InitFromArgs(ms *MsUptime, args []string, flagset *flag.FlagSet) *MsUptime {

	var cfg microservice.Configuration
	var db *database.Database

	if flagset == nil {
		flagset = flag.NewFlagSet("ms-uptime", flag.PanicOnError)
	}

	microservice.InitConfigurationFromArgs(&cfg, args, flagset)
	microservice.Init(&ms.MicroService, &cfg, nil)
	ms.starttime = time.Now()

	if len(ms.GetDBName()) > 0 {
		db = database.NewDatabase(ms.GetDBName(), ms.GetName())
		if db != nil {
			fmt.Printf("Got database configuration '%s'.\n", ms.GetDBName())
		} else {
			fmt.Printf("Bad database configuration '%s', ignoring it.\n", ms.GetDBName())
		}
	} else {
		fmt.Println("No database configured.")
	}

	ms.db = db
	ms.starttime = time.Now()
	uptimeHandler := ms.DefaultHandler()
	uptimeHandler.Get = ms.httpGetUptime
	ms.AddHandler("/uptime", uptimeHandler)

	if ms.db != nil {
		ms.db.Open()
		ms.GetLogger().Printf("Opened database '%s' with buckets 'uptime/%s.\n", ms.db.Path, ms.db.Instance)
	}

	return ms
}

// ---------------------------------------------------------------------------

func (ms *MsUptime) getStarttime() time.Time {
	return ms.starttime
}

// ---------------------------------------------------------------------------

func (ms *MsUptime) httpGetUptime(w http.ResponseWriter, r *http.Request) (status int, contentLen int, msg string) {

	clientID := r.Header.Get("cid")
	status = http.StatusOK
	response := NewUptimeResponse("OK", ms)
	ms.SetResponseHeaders("application/json; charset=utf-8", w, r)
	w.WriteHeader(status)
	contentLen = ms.Reply(w, response)
	if ms.db != nil {
		now := time.Now()
		ms.db.MarkUptime(ms.getStarttime(), now)
		msg = fmt.Sprintf("Up: %6ds | cid: %-8.8s | db: %s", response.Uptime, clientID, ms.GetName())
	} else {
		msg = fmt.Sprintf("Up: %6ds | cid: %-8.8s | db: ----", response.Uptime, clientID)
	}
	return status, contentLen, msg
}
