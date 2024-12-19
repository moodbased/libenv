package sig

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"syscall"
)

var once sync.Once

func Int(reason string) {
	slog.Info("trying to shut down process by SIGINT", "reason", reason)
	once.Do(func() {
		err := syscall.Kill(os.Getpid(), syscall.SIGINT)
		if err != nil {
			fmt.Println("Error sending signal:", err)
		}
	})
}
