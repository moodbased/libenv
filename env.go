package libenv

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/moodbased/libenv/sig"
)

const (
	WrongInt   = -1024
	DefaultInt = 0
)

func LoadEnvBool(v string) bool {
	env := os.Getenv(v)
	if env != "" {
		b, err := strconv.ParseBool(env)
		if err != nil {
			sig.Int(fmt.Sprintf("[ENV] failed to parse '%s' of value '%s'\n", v, env))
			return false
		}
		slog.Info("[ENV] loaded", v, b)
		return b
	}
	slog.Info("[ENV] use default value", v, false)
	return false
}

func LoadEnvInt64(v string) int64 {
	env := os.Getenv(v)
	if env != "" {
		i, err := strconv.ParseInt(env, 10, 64)
		if err != nil {
			sig.Int(fmt.Sprintf("[ENV] failed to parse '%s' of value '%s'\n", v, env))
			return WrongInt
		}
		slog.Info("[ENV] loaded", v, i)
		return i
	}
	slog.Info("[ENV] use default value", v, 0)
	return DefaultInt
}

func LoadEnvInt(v string) int {
	return int(LoadEnvInt64(v))
}

func LoadEnvString(v string) string {
	env := os.Getenv(v)
	if env != "" {
		slog.Info("[ENV] loaded", v, env)
	} else {
		slog.Info("[ENV] use default value", v, env)
	}
	return env
}

// LoadEnvStringMute loads a string from the environment variable without logging.
func LoadEnvStringMute(v string) string {
	env := os.Getenv(v)
	if env != "" {
		slog.Info("[ENV] loaded", v, "******")
	} else {
		slog.Info("[ENV] use default value", v, env)
	}
	return env
}

// LoadEnvStrings loads a comma-separated string from the environment variable.
func LoadEnvStrings(v string) (strs []string) {
	env := os.Getenv(v)
	if len(env) == 0 {
		slog.Info("[ENV] use default value", v, strs)
		return
	}
	strs = strings.Split(env, ",")
	slog.Info("[ENV] loaded", v, strs)
	return strs
}
