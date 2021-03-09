package web

import (
	"context"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"server/config"
	"strings"
	"time"
)

var running = false

// CallWebService ...
type CallWebService struct {
	ctx context.Context
}

// NewCallWebService creates a new call web service
func NewCallWebService(ctx context.Context) *CallWebService {
	return &CallWebService{
		ctx: ctx,
	}
}

// GetCall ...
func (svc *CallWebService) GetCall(ctx context.Context) error {
	if running {
		return errors.New("Call already running")
	}
	go handleMakeCall()
	return nil
}

func handleMakeCall() {
	running = true

	cfg := config.Get()
	var trueResult = cfg.TRUEResult
	var falseResult = cfg.FALSEResult
	var url string
	var path = "/output/result.json"
	var iserr = false

	// delete result.json if exists
	// var _, errfs = os.Stat(path)
	// if os.IsExist(errfs) {
	errfs := os.Remove(path)
	if errfs != nil {
		log.Printf("Can't delete file /output/result.json: %v", errfs)
	}
	// }

	log.Printf("Starting voip...")
	cmd := exec.Command("/entry.sh")
	err := cmd.Start()
	if err != nil {
		log.Printf("Cat't start voip: %v", err)
		iserr = true
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with: %v", err)

	if err != nil {
		log.Print("Voip returned not 0")
		iserr = true
	}

	// Анализ JSON
	if iserr {
		url = falseResult
	} else {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("Can't read file /output/result.json: %v", err)
			url = falseResult
		} else {
			s := string(b)
			if strings.Contains(s, "FAIL") {
				url = falseResult
			} else {
				url = trueResult
			}
		}
	}

	// Отправка результата
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(url)
	if err != nil {
		log.Printf("Connection failed: %v", err)
	} else {
		log.Printf("Responce status code: %v", res.StatusCode)
	}

	running = false
}
