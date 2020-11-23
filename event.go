package gom

import (
	"context"
	"fmt"
)

// 事件 y
// 监听者=处理者 y

type EventFactory struct {
	maps map[string]*Event
}

var imageFactory *EventFactory

func GetEventFactory() *EventFactory {
	if imageFactory == nil {
		imageFactory = &EventFactory{
			maps: make(map[string]*Event),
		}
	}
	return imageFactory
}

func (f *EventFactory) Get(ctx *context.Context, eventName string) *Event {
	event := f.maps[eventName]
	if event == nil {
		event = NewEvent(ctx, eventName) // 监听者-处理类
		f.maps[eventName] = event
	}

	return event
}

type Event struct {
	data      string
	context   *context.Context
	listeners []Listener
}

func NewEvent(ctx *context.Context, eventName string) *Event {
	// Load image file
	data := fmt.Sprintf("image data %s", eventName)
	return &Event{
		data:    data,
		context: ctx,
	}
}

func (i *Event) Data() string {
	return i.data
}

func (i *Event) Attach(l Listener) string {
	i.listeners = append(i.listeners, l)
	return i.data
}

// 监听者
type Listener struct {
	*Event
}

func NewListener(ctx *context.Context, eventName string) *Listener {
	event := GetEventFactory().Get(ctx, eventName)
	return &Listener{
		Event: event,
	}
}

func (i *Listener) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}
