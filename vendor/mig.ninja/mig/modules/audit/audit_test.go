// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Contributor: Arun Sori   <arunsori94@gmail.com>

package audit /* import "mig.ninja/mig/modules/audit" */

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/mozilla/libaudit-go"
	"mig.ninja/mig/testutil"
)

func TestRegistration(t *testing.T) {
	testutil.CheckModuleRegistration(t, "audit")
}

func TestReadAuditMessages(t *testing.T) {
	// read test messages
	f, err := os.OpenFile("/tmp/messages", os.O_RDONLY, 0644)
	if err != nil {
		t.Errorf("open file failed %v", err)
	}
	defer f.Close()
	var (
		message string
		scanner *bufio.Scanner
	)
	// write packed events
	scanner = bufio.NewScanner(f)
	w, err := os.OpenFile("/tmp/packedlog", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		t.Errorf("open file failed %v", err)
	}
	defer w.Close()

	//go writeJSON(w)
	go dispatchEventMozdef("http://localhost:4443")

	for scanner.Scan() {
		message = scanner.Text()
		index := strings.Index(message, " ")
		if index == -1 {
			t.Errorf("message malformed")
		}
		// strip type=* from the message
		msgType := message[5:index]
		// strip ' msg='
		message = message[index+5:]
		headerType, ok := libaudit.MsgTypeTab[msgType]
		if !ok {
			t.Errorf("msg type not found")
		}
		event, err := libaudit.ParseAuditEvent(message, headerType, true)
		if err != nil {
			t.Errorf("parse failed %v", err)
		}
		messageHandler(event, err)
	}
	for len(jsonBuffChan) > 0 {
		time.Sleep(time.Millisecond * 10)
	}
}

func TestPeriodReadAuditMessages(t *testing.T) {
	// read test messages
	f, err := os.OpenFile("/tmp/messages", os.O_RDONLY, 0644)
	if err != nil {
		t.Errorf("open file failed %v", err)
	}
	defer f.Close()
	var (
		message string
		scanner *bufio.Scanner
		count   int
	)
	// write packed events
	scanner = bufio.NewScanner(f)
	w, err := os.OpenFile("/tmp/packedlog", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		t.Errorf("open file failed %v", err)
	}
	defer w.Close()

	//go writeJSON(w)
	go dispatchEventMozdef("http://localhost:4443")

	for scanner.Scan() {
		if count == 100 {
			time.Sleep(time.Second * 1)
			count = 0
		}
		message = scanner.Text()
		index := strings.Index(message, " ")
		if index == -1 {
			t.Errorf("message malformed")
		}
		// strip type=* from the message
		msgType := message[5:index]
		// strip ' msg='
		message = message[index+5:]
		headerType, ok := libaudit.MsgTypeTab[msgType]
		if !ok {
			t.Errorf("msg type not found")
		}
		event, err := libaudit.ParseAuditEvent(message, headerType, true)
		if err != nil {
			t.Errorf("parse failed %v", err)
		}
		messageHandler(event, err)
		count++
	}
	for len(jsonBuffChan) > 0 {
		time.Sleep(time.Millisecond * 10)
	}
}

func writeJSON(output io.Writer) {
	for {
		select {
		case msg := <-jsonBuffChan:
			// fmt.Println("Writing")
			msgBytes, err := json.Marshal(*msg)
			if err != nil {
				panic(err)
			}
			msgBytes = append(msgBytes, []byte("\n")...)
			left := len(msgBytes)
			for left > 0 {
				nb, err := output.Write(msgBytes)

				if err != nil {
					// let the agent know that the message sending is failing
					sendLogMessage(fmt.Sprintf("dispatching of event %v is failed", msg.Details["auditserial"]))
					panic(err)
				}
				left -= nb
				msgBytes = msgBytes[nb:]
			}

		}
	}
}

func BenchmarkModule(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t testing.T
		TestReadAuditMessages(&t)
	}
}
