package slidingwindow

import (
	"log"
	"slidingwindow/bucket"
	"sync"
	"time"
)

type SlidingWindow struct {
	sync.RWMutex
	broken          bool             //是否已经触发熔断
	size            int              //滑动窗口大小
	buckets         []*bucket.Bucket //桶队列
	reqThreshold    int              //触发熔断的请求阈值
	failedThreshold float64          //触发熔断的失败阈值，failed次数/请求总次数
	lastBreakTIme   time.Time        //上一次触发熔断的时间
	// seeker bool 	//探测者是否被派出
	brokeTimeGap time.Duration //熔断恢复的时间间隔
}

func NewSlidingWindow(s int, reqT int, failedT float64, brokeTG time.Duration) *SlidingWindow {
	return &SlidingWindow{
		size:            s,
		buckets:         make([]*bucket.Bucket, 0, s),
		reqThreshold:    reqT,
		failedThreshold: failedT,
		brokeTimeGap:    brokeTG,
	}
}

//追加新桶
func (sw *SlidingWindow) AppendBucket() {
	sw.Lock()
	defer sw.Unlock()

	sw.buckets = append(sw.buckets, bucket.NewBucket())
	//若超出滑动窗口大小，则删除第一个桶
	if len(sw.buckets) >= sw.size+1 {
		sw.buckets = sw.buckets[1:]
	}
}

//获取最新的桶
func (sw *SlidingWindow) GetLastBucket() *bucket.Bucket {
	if len(sw.buckets) == 0 {
		sw.AppendBucket()
	}
	return sw.buckets[len(sw.buckets)-1]
}

//在桶中记录当次请求结果
func (sw *SlidingWindow) RecordReqResult(result bool) {
	sw.GetLastBucket().Record(result)
}

//启动滑动窗口
func (sw *SlidingWindow) Start() {
	go func() {
		//每隔100毫秒添加一个新桶
		for {
			sw.AppendBucket()
			time.Sleep(time.Millisecond * 100)
		}
	}()
}

//根据当前滑动窗口判断是否需要触发熔断
func (sw *SlidingWindow) BreakJudge() bool {
	sw.Lock()
	defer sw.Unlock()

	total := 0
	failed := 0
	for _, v := range sw.buckets {
		total += v.Total
		failed += v.Failed
	}
	//失败的阈值超过设置的失败阈值，并且请求总次数超过设置的请求阈值
	if float64(failed)/float64(total) > sw.failedThreshold && total > sw.reqThreshold {
		return true
	}
	return false
}

//判断是否超过熔断时间间隔
func (sw *SlidingWindow) OverBrokenTimeGap() bool {
	return time.Since(sw.lastBreakTIme) > sw.brokeTimeGap
}

// 监控滑动窗口的总失败次数与是否开启熔断
func (sw *SlidingWindow) Moniter() {
	go func() {
		for {
			if sw.broken {
				if sw.OverBrokenTimeGap() {
					sw.Lock()
					sw.broken = false
					sw.Unlock()
				}
				continue
			}
			if sw.BreakJudge() {
				sw.Lock()
				sw.broken = true
				sw.lastBreakTIme = time.Now()
				sw.Unlock()
			}
		}
	}()
}

//每隔1秒显示当前是否处于熔断状态
func (sw *SlidingWindow) ShowStatus() {
	go func() {
		for {
			broken := "false"
			if sw.broken {
				broken = "true"
			}
			log.Println(broken)
			time.Sleep(time.Second)
		}
	}()
}

// 获取当前熔断状态
func (sw *SlidingWindow) Broken() bool {
	return sw.broken
}
