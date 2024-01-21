package utils

import (
	log "github.com/sirupsen/logrus" // Import logrus package
)

func LogError(action, message string, err error) {
	log.WithFields(log.Fields{
		"action":  action,
		"message": message,
		"error":   err,
	}).Error("Error occurred")
}
