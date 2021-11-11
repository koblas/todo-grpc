package logger

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/sirupsen/logrus"
// )

// type SumoLogicHook struct {
// 	client      *http.Client
// 	endPointUrl string
// 	tags        []string
// 	host        string
// 	levels      []logrus.Level
// 	sender      chan *logrus.Entry
// }

// type SumoLogicMesssage struct {
// 	Tags  []string    `json:"tags,omitempty"`
// 	Msg   string      `json:"msg,omitempty"`
// 	Host  string      `json:"host"`
// 	Level string      `json:"level"`
// 	Data  interface{} `json:"data"`
// }

// func NewSumoLogicHook(endPointUrl string, host string, level logrus.Level, tags ...string) *SumoLogicHook {
// 	levels := []logrus.Level{}
// 	for _, l := range logrus.AllLevels {
// 		if l <= level {
// 			levels = append(levels, l)
// 		}
// 	}

// 	sender := make(chan *logrus.Entry, 10)

// 	// Copy of DefaultTransport
// 	/*
// 		transport := http.Transport{
// 			Proxy: http.ProxyFromEnvironment,
// 			DialContext: (&net.Dialer{
// 				Timeout:   30 * time.Second,
// 				KeepAlive: 30 * time.Second,
// 				DualStack: true,
// 			}).DialContext,
// 			MaxIdleConns:          100,
// 			IdleConnTimeout:       90 * time.Second,
// 			TLSHandshakeTimeout:   10 * time.Second,
// 			ExpectContinueTimeout: 1 * time.Second,
// 		}

// 		transport.IdleConnTimeout = 5 * time.Second
// 		transport.MaxIdleConns = 5
// 	*/

// 	hook := SumoLogicHook{
// 		client: &http.Client{
// 			Timeout: time.Duration(30 * time.Second),
// 			// Transport: &transport,
// 		},
// 		host:        host,
// 		tags:        tags,
// 		endPointUrl: endPointUrl,
// 		levels:      levels,
// 		sender:      sender,
// 	}

// 	// go hook.messageLoop(sender)

// 	return &hook
// }

// func (hook *SumoLogicHook) postMessage(entry *logrus.Entry) error {
// 	rawData := make(logrus.Fields, len(entry.Data))
// 	for k, v := range entry.Data {
// 		switch v := v.(type) {
// 		case error:
// 			// Otherwise errors are ignored by `encoding/json`
// 			// https://github.com/Sirupsen/logrus/issues/137
// 			rawData[k] = v.Error()
// 		default:
// 			rawData[k] = v
// 		}
// 	}

// 	message := SumoLogicMesssage{
// 		Host:  hook.host,
// 		Level: strings.ToUpper(entry.Level.String()),
// 		Data:  rawData,
// 		Tags:  hook.tags,
// 		Msg:   entry.Message,
// 	}

// 	payload, err := json.Marshal(message)
// 	if err != nil {
// 		return err
// 	}
// 	req, err := http.NewRequest(
// 		"POST",
// 		hook.endPointUrl,
// 		bytes.NewBuffer(payload),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	if resp, err := hook.client.Do(req); err == nil {
// 		defer resp.Body.Close()
// 		if _, err := io.Copy(ioutil.Discard, resp.Body); err != nil {
// 			return err
// 		}
// 	} else {
// 		return err
// 	}

// 	return nil
// }

// func (hook *SumoLogicHook) postMessageWithPrint(entry *logrus.Entry) {
// 	if err := hook.postMessage(entry); err != nil {
// 		fmt.Println("Error sending to Sumo", err)
// 	}
// }

// func (hook *SumoLogicHook) messageLoop(sender chan *logrus.Entry) {
// 	for {
// 		hook.postMessageWithPrint(<-sender)
// 	}
// }

// func (hook *SumoLogicHook) Fire(entry *logrus.Entry) error {
// 	// go hook.postMessageWithPrint(entry)
// 	// return nil

// 	// hook.sender <- entry
// 	// return nil

// 	return hook.postMessage(entry)
// }

// func (hook *SumoLogicHook) Levels() []logrus.Level {
// 	return hook.levels

// }
