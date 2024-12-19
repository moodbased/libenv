package libenv

import "testing"

func TestLoad(t *testing.T) {
	LoadEnvBool("TEST_BOOL")
	LoadEnvInt64("TEST_INT64")
	LoadEnvInt("TEST_INT")
	LoadEnvString("TEST_STRING")
	LoadEnvStringMute("TEST_STRING_MUTE")
	LoadEnvStrings("TEST_STRINGS")
}
