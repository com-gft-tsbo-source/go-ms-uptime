package msuptime

import (
	"time"

	"com.gft.tsbo-training.src.go/common/ms-framework/microservice"
)

// ###########################################################################
// ###########################################################################
// MsUptime Response
// ###########################################################################
// ###########################################################################

// UptimeResponse Encapsulates the reploy of ms-uptime
type UptimeResponse struct {
	microservice.Response
	Uptime int `json:"uptime"`
}

// ###########################################################################

// NewUptimeResponse Constructor of a response of ms-uptime
func NewUptimeResponse(status string, ms *MsUptime) *UptimeResponse {
	var r UptimeResponse
	microservice.InitResponseFromMicroService(&r.Response, ms, status)
	now := time.Now()
	r.Uptime = int(now.Sub(ms.getStarttime()).Seconds())
	return &r
}
