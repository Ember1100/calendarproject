package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func runPeriodicTasks(r *gin.Engine) {
	// 设置多个定时任务时间点
	targetTimes := []time.Time{
		time.Date(2024, 6, 18, 22, 36, 57, 333000000, time.Local),
		time.Date(2024, 6, 18, 22, 56, 20, 926000000, time.Local),
	}

	tickers := make([]*time.Ticker, len(targetTimes))

	for i, targetTime := range targetTimes {
		tickers[i] = time.NewTicker(time.Until(targetTime))
		defer tickers[i].Stop()

		go func(idx int) {
			for {
				select {
				case <-tickers[idx].C:
					runTask(r, idx)
					// 重置定时器,等待下一次执行
					tickers[idx] = time.NewTicker(time.Until(targetTimes[idx].Add(24 * time.Hour)))
				}
			}
		}(i)
	}
}

func runTask(r *gin.Engine, idx int) {
	fmt.Printf("执行定时任务 %d...\n", idx+1)

	// 在这里添加你需要执行的任务代码
	// 例如: 更新数据库、发送邮件等
}
