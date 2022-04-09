package logrus

import "log"

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	log.Print(args...)
}
