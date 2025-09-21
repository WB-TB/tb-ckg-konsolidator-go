package request

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"sync/atomic"
)

var reqid uint64

func ID() string {
	var sb strings.Builder

	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		rand.Read(buf[:])
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}

	sb.WriteString(hostname)
	sb.WriteString("-")
	sb.WriteString(b64[0:10])

	return fmt.Sprintf("%s-%08d", sb.String(), atomic.AddUint64(&reqid, 1))
}
