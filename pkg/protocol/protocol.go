package protocol

type EmptyChecker interface {
	Empty() bool
}

type ProtocolDetailer interface {
	EmptyChecker
	Protocol() int
}
