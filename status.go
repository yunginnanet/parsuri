package parsuri

type status int32

const (
	statusRunning status = iota
	statusClosed
)

func (l *Loader) isRunning() bool {
	return l.running.CompareAndSwap(int32(statusRunning), int32(statusRunning))
}
