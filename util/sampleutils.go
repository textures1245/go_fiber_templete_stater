package util

import log "github.com/sirupsen/logrus"

func SomeUtilMethods(data string) string {
	/** Define log component **/
	log := log.WithFields(log.Fields{
		"component": "Utils",
		"funciton":  "SomeUtilMethods",
	})
	log.Info("-= Some Utils Code =-")
	return "RESULT"
}
