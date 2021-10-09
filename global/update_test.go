package global

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestName(t *testing.T) {
	err := Update()
	if err != nil {
		log.Errorln(err.Error())
		return
	}
}
