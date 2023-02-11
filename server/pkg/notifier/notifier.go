package notifier

type Notifier[T any] struct {
	Listeners []func(T)
}

func Create[T any]() Notifier[T] {
	return Notifier[T]{
		Listeners: []func(T){},
	}
}

func (notifier Notifier[T]) Notify(value T) {
	for _, listener := range notifier.Listeners {
		listener(value)
	}
}

func (notifier *Notifier[T]) AddListener(listener func(T)) {
	notifier.Listeners = append(notifier.Listeners, listener)
}

func (notifier *Notifier[T]) RemoveAllListeners() {
	notifier.Listeners = []func(T){}
}
