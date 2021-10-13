package constants

import "time"

const (
	ReadDeadLine          = time.Second * time.Duration(30)
	WriteDeadLine         = time.Second * time.Duration(30)
	DefaultBackupInterval = time.Second * time.Duration(30)
	SoftwareName          = "cachy"
	ListenPort            = 8080
)
