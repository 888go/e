package e

import (
	"fmt"
	"math/rand"
	"time"
)

func X取随机ip() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", X取随机数(50, 254), X取随机数(50, 254), X取随机数(50, 254), X取随机数(50, 254))
	return ip
}
