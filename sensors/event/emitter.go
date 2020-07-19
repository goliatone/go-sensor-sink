package event

import (
	"reflect"
	"sync"
)

//Emitter event emitter struct
type Emitter struct {
	listeners map[interface{}][]Listener
	mutex     *sync.Mutex
}

//Listener struct
type Listener struct {
	callback func(...interface{})
	once     bool
}

//NewEmitter returns a new Emitter
func NewEmitter() *Emitter {
	return &Emitter{
		make(map[interface{}][]Listener),
		&sync.Mutex{},
	}
}

//AddListener add event listener
func (e *Emitter) AddListener(event string, callback func(...interface{})) *Emitter {
	return e.On(event, callback)
}

//On adds a new event listener
func (e *Emitter) On(event string, callback func(...interface{})) *Emitter {
	e.mutex.Lock()
	if _, ok := e.listeners[event]; !ok {
		e.listeners[event] = []Listener{}
	}

	e.listeners[event] = append(e.listeners[event], Listener{callback, false})
	e.mutex.Unlock()

	e.EmitSync("listener.add", []interface{}{event, callback})

	return e
}

//Once register a new one-time listener for a given event type
func (e *Emitter) Once(event string, callback func(...interface{})) *Emitter {
	e.mutex.Lock()
	if _, ok := e.listeners[event]; !ok {
		e.listeners[event] = []Listener{}
	}

	e.listeners[event] = append(e.listeners[event], Listener{callback, true})
	e.mutex.Unlock()

	e.EmitSync("listener.add", []interface{}{event, callback})
	return e
}

//RemoveListener removes given event callback
func (e *Emitter) RemoveListener(event string, callback func(...interface{})) *Emitter {
	return e.removeListenerInternal(event, callback, false)
}

//RemoveAllListeners remove all listeners or all listeners from an event
func (e *Emitter) RemoveAllListeners(event interface{}) *Emitter {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if event == nil {
		e.listeners = make(map[interface{}][]Listener)
		return e
	}
	if _, ok := e.listeners[event]; !ok {
		return e
	}
	delete(e.listeners, event)
	return e
}

//Listeners return slice with all listener callbacks
func (e *Emitter) Listeners(event string) []Listener {
	if _, ok := e.listeners[event]; !ok {
		return nil
	}
	return e.listeners[event]
}

//ListenersCount number of listeners
func (e *Emitter) ListenersCount(event string) int {
	return len(e.Listeners(event))
}

//EmitSync run all listeners of the given event in synchronous mode
func (e *Emitter) EmitSync(event string, args ...interface{}) *Emitter {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for _, v := range e.Listeners(event) {
		if v.once {
			e.removeListenerInternal(event, v.callback, true)
		}
		v.callback(args...)
	}

	return e
}

//EmitAsync run all listeners of the given event using goroutines
func (e *Emitter) EmitAsync(event string, args ...interface{}) *Emitter {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	for _, v := range e.Listeners(event) {
		if v.once {
			e.removeListenerInternal(event, v.callback, true)
		}
		go v.callback(args...)
	}
	return e
}

func (e *Emitter) removeListenerInternal(event string, callback func(...interface{}), surpress bool) *Emitter {
	if _, ok := e.listeners[event]; !ok {
		return e
	}

	for k, v := range e.listeners[event] {
		if reflect.ValueOf(v.callback).Pointer() == reflect.ValueOf(callback).Pointer() {
			e.listeners[event] = append(e.listeners[event][:k], e.listeners[event][k+1:]...)
			if !surpress {
				e.EmitSync("listener.remove", []interface{}{event, callback})
			}
			return e
		}
	}
	return e
}
