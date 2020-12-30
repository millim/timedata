package lib

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkPut(b *testing.B) {
	b.ResetTimer()
	userIDS := []int64{1, 2, 3}
	actions := []string{"z", "r", "p", "q", "v"}
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().Unix())
		uid := rand.Intn(3)
		a := rand.Intn(5)
		m := rand.Intn(1440)
		t := time.Now().Add(-time.Minute * time.Duration(m))
		Put(t, userIDS[uid], actions[a])
	}
}
