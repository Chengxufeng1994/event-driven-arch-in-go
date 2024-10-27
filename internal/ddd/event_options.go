package ddd

type EventOption interface {
	configureEvent(*eventBase)
}
