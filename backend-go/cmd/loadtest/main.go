// Command loadtest 是压测工具：固定并发打指定时长，统计 QPS / 延迟分位 / 状态码分布。
//
// 用 Go 原生并发，避免客户端成为瓶颈（Python 线程池受 GIL 限制，测不出服务端真实能力）。
//
// 示例：
//   go run ./cmd/loadtest -base http://127.0.0.1:8080 -ids PERF-01,PERF-02 -missing 100 -c 200 -d 10s
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	base := flag.String("base", "http://127.0.0.1:8080", "公开 API 基址")
	idsRaw := flag.String("ids", "", "存在的标签 ID，逗号分隔（为空则全部打不存在 ID）")
	missing := flag.Int("missing", 0, "不存在 ID 的数量（前缀 MISS）")
	concurrency := flag.Int("c", 100, "并发数")
	duration := flag.Duration("d", 10*time.Second, "压测时长")
	flag.Parse()

	ids := make([]string, 0, 32)
	if *idsRaw != "" {
		for _, s := range strings.Split(*idsRaw, ",") {
			if v := strings.TrimSpace(s); v != "" {
				ids = append(ids, v)
			}
		}
	}
	for i := 0; i < *missing; i++ {
		ids = append(ids, fmt.Sprintf("MISS-%04d", i))
	}
	if len(ids) == 0 {
		fmt.Println("必须提供 -ids 或 -missing")
		os.Exit(1)
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        *concurrency * 2,
			MaxIdleConnsPerHost: *concurrency * 2,
			MaxConnsPerHost:     *concurrency * 2,
			DisableCompression:  true,
		},
	}

	var (
		requested atomic.Int64
		wg        sync.WaitGroup
	)
	deadline := time.Now().Add(*duration)
	perWorker := make([][]time.Duration, *concurrency)
	perWorkerCodes := make([][]int, *concurrency)

	for w := 0; w < *concurrency; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			lat := make([]time.Duration, 0, 4096)
			codes := make([]int, 0, 4096)
			i := w
			for time.Now().Before(deadline) {
				target := ids[i%len(ids)]
				i++
				start := time.Now()
				resp, err := client.Get(*base + "/api/public/nfc/" + target + "/page")
				elapsed := time.Since(start)
				code := -1
				if err != nil {
					code = -1
				} else {
					code = resp.StatusCode
					// 必须读干响应体，否则连接无法复用，会造成端口耗尽（TIME_WAIT 爆炸）
					_, _ = io.Copy(io.Discard, resp.Body)
					_ = resp.Body.Close()
				}
				lat = append(lat, elapsed)
				codes = append(codes, code)
				requested.Add(1)
			}
			perWorker[w] = lat
			perWorkerCodes[w] = codes
		}(w)
	}

	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)

	// 汇总
	var all []time.Duration
	codeCount := map[int]int{}
	for w := range perWorker {
		all = append(all, perWorker[w]...)
		for _, c := range perWorkerCodes[w] {
			codeCount[c]++
		}
	}
	if len(all) == 0 {
		fmt.Println("没有完成任何请求")
		os.Exit(1)
	}
	sort.Slice(all, func(i, j int) bool { return all[i] < all[j] })

	p := func(pct float64) time.Duration {
		idx := int(float64(len(all)) * pct / 100)
		if idx >= len(all) {
			idx = len(all) - 1
		}
		return all[idx]
	}
	total := len(all)
	qps := float64(total) / elapsed.Seconds()

	fmt.Printf("并发=%d 时长=%s 实际=%s\n", *concurrency, duration.String(), elapsed.Round(time.Millisecond))
	fmt.Printf("总请求=%d  QPS=%.1f\n", total, qps)
	fmt.Printf("延迟  p50=%v  p90=%v  p95=%v  p99=%v  max=%v\n",
		p(50).Round(time.Microsecond), p(90).Round(time.Microsecond),
		p(95).Round(time.Microsecond), p(99).Round(time.Microsecond), all[len(all)-1])
	fmt.Printf("状态码 ")
	keys := make([]int, 0, len(codeCount))
	for k := range codeCount {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%d:%d  ", k, codeCount[k])
	}
	fmt.Println()

	if codeCount[500] > 0 || codeCount[-1] > 0 {
		fmt.Println("存在 5xx 或连接错误")
		os.Exit(1)
	}
}
