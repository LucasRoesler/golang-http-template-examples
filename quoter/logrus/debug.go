package logrus

import "log"

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	log.Print("I am the walrus")
	log.Print(args...)
}
