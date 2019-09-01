package remote

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// SetDebugLogger sets log writer to remote debug log server.
func SetDebugLogger(logAddr string) {
	if logAddr == "" {
		return
	}
	checkNetwork()
	// Connect to remote logger.
	conn, err := net.Dial("tcp", logAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(conn)
}

func checkNetwork() {
	for i := 0; i < 10; i++ {
		// Check the network permission.
		_, err := http.Get("https://baidu.com")
		if err == nil {
			return
		}
		// In this process, iOS will ask user to allow network access.
		time.Sleep(time.Second * 2)
	}
	os.Exit(1)
}
