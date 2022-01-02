package bucket

import (
	"sync"
	"time"
)

type Bucket struct {
	sync.RWMutex
	Total     int //请求总数
	Failed    int //失败总数
	Timestamp time.Time
}

func NewBucket() *Bucket {
	return &Bucket{
		Timestamp: time.Now(),
	}
}

func (b *Bucket) Record(result bool) {
	b.Lock()
	defer b.Unlock()

	if !result {
		b.Failed++
	}
	b.Total++
}
