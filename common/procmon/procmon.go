package procmon

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/process"
	"go.temporal.io/server/common/clock"
	"go.temporal.io/server/common/dynamicconfig"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	"go.temporal.io/server/common/metrics"
)

type ProcessMonitor struct {
	sampleInterval dynamicconfig.DurationPropertyFn
	metricsHandler metrics.Handler
	timeSource     clock.TimeSource
	logger         log.Logger
	stop           chan struct{}
	done           chan struct{}
	ticks          chan struct{}
}

func NewProcessMonitor(
	sampleInterval dynamicconfig.DurationPropertyFn,
	metricsHandler metrics.Handler,
	timeSource clock.TimeSource,
	logger log.Logger,
) *ProcessMonitor {
	return &ProcessMonitor{
		sampleInterval: sampleInterval,
		metricsHandler: metricsHandler,
		timeSource:     timeSource,
		logger:         logger,
		stop:           make(chan struct{}),
		done:           make(chan struct{}),
		ticks:          make(chan struct{}, 1),
	}
}

func (m *ProcessMonitor) Start() {
	m.ticks <- struct{}{}
	go m.loop()
}

func (m *ProcessMonitor) Stop() {
	close(m.stop)
	<-m.done
}

func (m *ProcessMonitor) loop() {
	defer close(m.done)
	var timer clock.Timer
	for {
		select {
		case <-m.stop:
			if timer != nil {
				timer.Stop()
			}
			return
		case <-m.ticks:
			if err := m.record(); err != nil {
				m.logger.Error("Error recording process metrics", tag.Error(err))
			}
			delay := m.sampleInterval()
			if delay < 0 {
				delay = 0
			}
			timer = m.timeSource.AfterFunc(delay, m.tick)
		}
	}
}

func (m *ProcessMonitor) tick() {
	select {
	case m.ticks <- struct{}{}:
	default:
	}
}

func (m *ProcessMonitor) record() error {
	proc, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return fmt.Errorf("get current process info: %w", err)
	}

	cpuPercent, err := proc.CPUPercent()
	if err != nil {
		return fmt.Errorf("get process CPU percent: %w", err)
	}
	metrics.ProcessStats.CPUUsage.With(m.metricsHandler).Record(cpuPercent)

	memInfo, err := proc.MemoryInfo()
	if err != nil {
		return fmt.Errorf("get process memory info: %w", err)
	}
	memBytes := float64(memInfo.RSS)
	metrics.ProcessStats.MemoryUsage.With(m.metricsHandler).Record(memBytes)
	return nil
}
