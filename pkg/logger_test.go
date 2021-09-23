package pkg

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestTracePrintMessage(t *testing.T){
	var buf bytes.Buffer
	log := New(&buf, Trace)

	msg := &message {}

	log.Trace("Example msg", nil)

	json.Unmarshal(buf.Bytes(), msg)

	if msg.Level != Trace.ToString() {
		t.Logf("Expected level: TRACE. Got %s", msg.Level)
		t.FailNow()
	}

	if msg.Message != "Example msg" {
		t.Logf("Expected message: \"Example msg\". Got %s", msg.Message)
		t.FailNow()
	}

	if msg.Time == "" {
		t.Log("Expected time to not be empty")
		t.FailNow()
	}
}

func TestInfoPrintMessage(t *testing.T){
	var buf bytes.Buffer
	log := New(&buf, Trace)

	msg := &message {}

	log.Info("Example msg", nil)

	json.Unmarshal(buf.Bytes(), msg)

	if msg.Level != Info.ToString() {
		t.Logf("Expected level: INFO. Got %s", msg.Level)
		t.FailNow()
	}

	if msg.Message != "Example msg" {
		t.Logf("Expected message: \"Example msg\". Got %s", msg.Message)
		t.FailNow()
	}

	if msg.Time == "" {
		t.Log("Expected time to not be empty")
		t.FailNow()
	}
}

func TestErrorPrintMessage(t *testing.T){
	var buf bytes.Buffer
	log := New(&buf, Trace)

	msg := &message {}

	log.Error("Example msg", nil)

	json.Unmarshal(buf.Bytes(), msg)

	if msg.Level != Error.ToString() {
		t.Logf("Expected level: ERROR. Got %s", msg.Level)
		t.FailNow()
	}

	if msg.Message != "Example msg" {
		t.Logf("Expected message: \"Example msg\". Got %s", msg.Message)
		t.FailNow()
	}

	if msg.Time == "" {
		t.Log("Expected time to not be empty")
		t.FailNow()
	}
}

func TestFatalPrintMessage(t *testing.T){
	var buf bytes.Buffer
	log := New(&buf, Trace)

	msg := &message {}

	log.Fatal("Example msg", nil)

	json.Unmarshal(buf.Bytes(), msg)

	if msg.Level != Fatal.ToString() {
		t.Logf("Expected level: FATAL. Got %s", msg.Level)
		t.FailNow()
	}

	if msg.Message != "Example msg" {
		t.Logf("Expected message: \"Example msg\". Got %s", msg.Message)
		t.FailNow()
	}

	if msg.Time == "" {
		t.Log("Expected time to not be empty")
		t.FailNow()
	}
}

func TestWarnPrintMessage(t *testing.T){
	var buf bytes.Buffer
	log := New(&buf, Trace)

	msg := &message {}

	log.Warn("Example msg", nil)

	json.Unmarshal(buf.Bytes(), msg)

	if msg.Level != Warn.ToString() {
		t.Logf("Expected level: WARN. Got %s", msg.Level)
		t.FailNow()
	}

	if msg.Message != "Example msg" {
		t.Logf("Expected message: \"Example msg\". Got %s", msg.Message)
		t.FailNow()
	}

	if msg.Time == "" {
		t.Log("Expected time to not be empty")
		t.FailNow()
	}
}

func TestDebugPrintMessage(t *testing.T){
	var buf bytes.Buffer
	log := New(&buf, Trace)

	msg := &message {}

	log.Debug("Example msg", nil)

	json.Unmarshal(buf.Bytes(), msg)

	if msg.Level != Debug.ToString() {
		t.Logf("Expected level: DEBUG. Got %s", msg.Level)
		t.FailNow()
	}

	if msg.Message != "Example msg" {
		t.Logf("Expected message: \"Example msg\". Got %s", msg.Message)
		t.FailNow()
	}

	if msg.Time == "" {
		t.Log("Expected time to not be empty")
		t.FailNow()
	}
}

func TestLogBelowMinLevel(t *testing.T) {
	var buf bytes.Buffer
	log := New(&buf, Error)

	msg := &message {}

	log.Warn("Example msg", nil)

	json.Unmarshal(buf.Bytes(), msg)


	if msg.Level == Warn.ToString() {
		t.Logf("Expect empty level. Got %s", msg.Level)
		t.FailNow()
	}

	if msg.Message == "Example msg" {
		t.Logf("Expected empty message. Got %s", msg.Message)
		t.FailNow()
	}

	if msg.Time != "" {
		t.Log("Expected time to be empty")
		t.FailNow()
	}
}