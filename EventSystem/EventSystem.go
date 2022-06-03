package EventSystem

type EventHandler func(name string, arguments ...interface{})

type Event struct {
	handlers map[string]EventHandler
}

func (event *Event) AddHandler(name string, handler EventHandler) {
	event.handlers[name] = handler
}

func (event *Event) DeleteHandler(name string) {
	delete(event.handlers, name)
}

func (event *Event) Call(arguments ...interface{}) {
	for name, handler := range event.handlers {
		handler(name, arguments...)
	}
}

func CreateEvent() Event {
	return Event{make(map[string]EventHandler)}
}
