package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func about() {
	fmt.Printf("This is burry in version %s\n", VERSION)
}

func datadirs() {
	stateurl := strings.Join([]string{"http://", endpoint, INFRA_SVC_EXHIBITOR}, "")
	econfig := &ExhibitorState{}
	if err := get(stateurl, econfig); err != nil {
		log.WithFields(log.Fields{"func": "datadirs"}).Error(fmt.Sprintf("Can't parse response from endpoint: %s", err))
	} else {
		log.WithFields(log.Fields{"func": "datadirs"}).Info(fmt.Sprintf("Config %#v", econfig))
	}
}

func get(url string, payload interface{}) error {
	c := &http.Client{Timeout: 2 * time.Second}
	r, err := c.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(payload)
}

func lookupst(name string) int {
	switch strings.ToLower(name) {
	case "tty":
		return 0
	case "local":
		return 1
	default:
		return -1
	}
}

func store(path string, val string) {
	log.WithFields(log.Fields{"func": "store"}).Info(fmt.Sprintf("Stored %s", path))
}