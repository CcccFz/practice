package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc func(v interface{}) bool
)

type Pub struct {
	m          sync.RWMutex
	size       int
	timeout    time.Duration
	subscribers map[subscriber]topicFunc
}

func NewPub(size int, timeout time.Duration) *Pub {
	return &Pub{
		size:        size,
		timeout:     timeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *Pub) SubTopic(topic topicFunc) chan interface{} {
	sub := make(chan interface{}, p.size)

	p.m.Lock()
	p.subscribers[sub] = topic
	p.m.Unlock()

	return sub
}

func (p *Pub) SubAll() chan interface{} {
	return p.SubTopic(nil)
}

func (p *Pub) Publish(content interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topicFunc := range p.subscribers {
		wg.Add(1)
		go p.publishTopic(sub, topicFunc, content, &wg)
	}

	wg.Wait()
	time.Sleep(2 * time.Second)
}

func (p *Pub) publishTopic(sub subscriber, topicFunc topicFunc, content interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	if topicFunc != nil && !topicFunc(content) {
		return
	}

	select {
	case sub <- content:
	case <- time.After(p.timeout):
	}
}

func (p *Pub) Evict(sub subscriber) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

func (p *Pub) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func main() {
	pub := NewPub(10, 5)

	all := pub.SubAll()
	golang := pub.SubTopic(func(v interface{}) bool {
		if msg, ok := v.(string); ok {
			if strings.Contains(msg, "golang") {
				return true
			}
		}
		return false
	})

	go func() {
		for msg := range golang {
			fmt.Println("golang: ", msg)
		}
	}()

	go func() {
		for msg := range all {
			fmt.Println("all: ", msg)
		}
	}()

	pub.Publish("hello, golang")
	pub.Publish("hello, world")

	time.Sleep(5 * time.Second)
}