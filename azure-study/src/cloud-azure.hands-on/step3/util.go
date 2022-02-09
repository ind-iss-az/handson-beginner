package main

import (
	"os"
	"time"
)

func GetDatetime() string {
	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)
	const layout = "2006-01-02 15:04:05"
	return nowJST.Format(layout)
}

func GetHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}