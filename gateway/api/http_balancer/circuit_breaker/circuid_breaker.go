package circuit_breaker

import (
	"time"

	"github.com/sirupsen/logrus"
)

type CircuitStatus int

//go:generate stringer -type=CircuitStatus -output=./circuid_types_string.go
const (
	// Closed - закрытое состояние
	Closed CircuitStatus = iota
	// Opened - открытое состояние
	Opened
	// HalfOpened - полуоткрытое состояние
	HalfOpened
)

type circuitBreaker struct {
	currStatus     CircuitStatus
	requestCounter int
	// лимит, когда разрешены ошибки в закрытом состоянии
	requestLimit int
	defaultDelay time.Duration

	// до какого времени закрыт circuit breaker
	closedTo time.Time
}

func NewCircuiBreaker(requestLimit int, defaultDelay time.Duration) *circuitBreaker {
	return &circuitBreaker{
		currStatus:     Closed,
		requestCounter: 0,
		requestLimit:   requestLimit,
		defaultDelay:   defaultDelay,
	}
}

// CountError считает ошибку в запросе
func (cb *circuitBreaker) CountError() {
	cb.requestCounter++
	if cb.currStatus == HalfOpened {
		cb.currStatus = Opened
		cb.reset()
		logrus.Infof("circuid status now %s until %s", cb.currStatus.String(), cb.closedTo.Format(time.RFC3339))
	}

	if cb.requestCounter >= cb.requestLimit {
		switch cb.currStatus {
		case Closed:
			cb.currStatus = Opened
			cb.reset()
		}
		logrus.Infof("circuid status now %s until %s", cb.currStatus.String(), cb.closedTo.Format(time.RFC3339))
	}
}

func (cb *circuitBreaker) reset() {
	cb.closedTo = time.Now().Add(cb.defaultDelay)
	cb.requestCounter = 0
}

func (cb *circuitBreaker) GetStatus() CircuitStatus {
	if cb.currStatus == Closed {
		return Closed
	}
	if cb.closedTo.Before(time.Now()) {
		switch cb.currStatus {
		case Opened:
			cb.currStatus = HalfOpened
		}
		cb.reset()
		logrus.Infof("circuid status now %s until %s", cb.currStatus.String(), cb.closedTo.Format(time.RFC3339))
	}
	return cb.currStatus
}

func (cb *circuitBreaker) CountHalfOpenedRequests() int {
	cb.requestCounter++
	if cb.requestCounter >= cb.requestLimit {
		cb.currStatus = Closed
		cb.reset()
		logrus.Infof("circuid status now %s until %s", cb.currStatus.String(), cb.closedTo.Format(time.RFC3339))
	}
	return cb.requestCounter
}
