package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime/pprof"
	"syscall"
	"time"
)

// Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep,
// которая приостанавливает выполнение текущей горутины.
//
// Важно: в отличие от настоящей time.Sleep, ваша функция должна именно блокировать выполнение
// (например, через таймер или цикл), а не просто вызывать time.Sleep :) — это упражнение.
//
// Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).

// BusySleep burns the CPU by continiously checking if
// time.Since(start) has exceeded d.
func BusySleep(d time.Duration) {
	start := time.Now()
	for time.Since(start) < d {
	}
}

func ContextSleep(d time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	<-ctx.Done()
}

func ChanSleep(d time.Duration) {
	select {
	case <-time.After(d):
	}
}

func OSSleep(d time.Duration) {
	tv := syscall.NsecToTimeval(d.Nanoseconds())
	syscall.Select(0, nil, nil, nil, &tv)
}

func NanoSleep(d time.Duration) {
	ts := syscall.NsecToTimespec(d.Nanoseconds())
	syscall.Nanosleep(&ts, nil)
}

func TickerSleep(d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	<-ticker.C
}

func TimedSleep(d time.Duration) {
	timer := time.NewTimer(d)
	defer timer.Stop()

	<-timer.C
}

func main() {
	var (
		d     = time.Second
		funcs = map[string]func(time.Duration){
			"  Looped  ": BusySleep, // this one burns the cpu a lot
			"time.After": ChanSleep,
			"  Context ": ContextSleep,

			// these block the os thread, but its kinda cool
			"OS Select ": OSSleep,
			"Unix Nano ": NanoSleep, // this one is unix only

			// these ones are just for examples
			"  Ticker  ": TickerSleep,
			"   Timer  ": TimedSleep,
			"time.Sleep": time.Sleep,
		}
	)

	for fname, f := range funcs {
		displayResults(d, fname, f)
	}
}

func displayResults(d time.Duration, fname string, f func(time.Duration)) {
	tmpFile, err := os.CreateTemp("", "cpu_profile_*.prof")
	if err != nil {
		panic("idk what happened but im scared")
	}
	defer os.Remove(tmpFile.Name())

	fmt.Printf("|------------------- %v -------------------|\n", fname)
	start := time.Now()
	fmt.Printf("Started at: %v\n", start.UTC())

	pprof.StartCPUProfile(tmpFile)
	f(d)
	pprof.StopCPUProfile()

	fmt.Printf(" Exited at: %v\n\n", time.Now().UTC())

	if fname == "  Looped  " {
		cmd := exec.Command("go", "tool", "pprof", "-top", os.Args[0], tmpFile.Name())
		out, _ := cmd.CombinedOutput()
		fmt.Print(string(out))
	}

	fmt.Printf("\n\n\n")
}
