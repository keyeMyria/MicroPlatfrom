/*
@Time : 2019/4/23 15:48 
@Author : kenny zhu
@File : common.go
@Software: GoLand
@Others:
*/
package monitor

import (
	"github.com/micro/go-micro/registry"
	"time"
)


type Heartbeat struct {
	id string// The ID of the heartbeat. Will generally be the Node ID.
	service *registry.Service // The service sending the heartbeat, support multi-service
	timestamp int64 // Unix time at which this was sent
	ttl int64  // The time to live for this heartbeat,  for out-dated.
	weights int // node weight.
}

const (
	// default ttl value
    HeartBeatTTL = time.Second * 60

	// state define here.
	ServiceStatus = "ServiceStatus"
	NormalState = "normal"
	UpThresh = "upThresh"
	DownThresh = "downThresh"

	// for Bucket algorithm .
	maxCalls  = 2000
	upLimits = int(maxCalls * 0.9)
	downLimits = int(maxCalls * 0.7)
)