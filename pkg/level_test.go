package pkg

import "testing"

func TestTraceLevelToString(t *testing.T) {
	lvl := Trace.ToString()
	if lvl != "TRACE"{
		t.Logf("Expected value TRACE. Got %s", lvl)
		t.Fail()
	}
}

func TestDebugLevelToString(t *testing.T) {
	lvl := Debug.ToString()
	if lvl != "DEBUG"{
		t.Logf("Expected value DEBUG. Got %s", lvl)
		t.Fail()
	}
}

func TestInfoLevelToString(t *testing.T) {
	lvl := Info.ToString()
	if lvl != "INFO"{
		t.Logf("Expected value INFO. Got %s", lvl)
		t.Fail()
	}
}

func TestWarningLevelToString(t *testing.T) {
	lvl := Warn.ToString()
	if lvl != "WARNING"{
		t.Logf("Expected value WARNING. Got %s", lvl)
		t.Fail()
	}
}

func TestErrorLevelToString(t *testing.T) {
	lvl := Error.ToString()
	if lvl != "ERROR"{
		t.Logf("Expected value ERROR. Got %s", lvl)
		t.Fail()
	}
}

func TestFatalLevelToString(t *testing.T) {
	lvl := Fatal.ToString()
	if lvl != "FATAL"{
		t.Logf("Expected value FATAL. Got %s", lvl)
		t.Fail()
	}
}