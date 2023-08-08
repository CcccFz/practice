package main

import (
	"log"
	"sync"
	"time"
)

type DaYe interface {
	Say(wg *sync.WaitGroup)
	Listen(wg *sync.WaitGroup)
}

type daYe struct {
	name  string
	total uint32
	trans Trans
}

type ZhangDaYe struct {
	*daYe
}

func NewZhangDaYe() DaYe {
	trans, err := Listen("localhost:9000")
	if err != nil {
		log.Panic(err)
	}

	return &ZhangDaYe{&daYe{name: "张大爷", total: 100000, trans: trans}}
}

func (z *ZhangDaYe) Say(wg *sync.WaitGroup) {
	defer wg.Done()

	i := uint32(0)

	go func() {
		for {
			time.Sleep(3 * time.Second)
			log.Println(z.name+" 说: ", i)
		}
	}()

	for i < z.total {
		if err := z.trans.Send(msgTypeEatZReq); err != nil {
			log.Panic(err)
		}
		i++
	}
}

func (z *ZhangDaYe) Listen(wg *sync.WaitGroup) {
	defer wg.Done()

	i := uint32(0)

	go func() {
		for {
			time.Sleep(3 * time.Second)
			log.Println(z.name+" 听: ", i)
		}
	}()

	for i < z.total*3 {
		msg, err := z.trans.Receive()
		if err != nil {
			log.Panic(err)
		}

		typ := msg.Type()
		switch typ {
		case msgTypePlaceLReq:
			if err = z.trans.Send(msgTypePlaceZRes); err != nil {
				log.Panic(err)
			}
		case msgTypePleaseLReq:
			if err = z.trans.Send(msgTypePleaseZRes); err != nil {
				log.Panic(err)
			}
		case msgTypeEatLRes:
		default:
			log.Panicf("听不懂的消息类型: %d", typ)
		}
		i++
	}
}

type LiDaYe struct {
	*daYe
}

func NewLiDaYe() DaYe {
	trans, err := Connect("localhost:9000")
	if err != nil {
		log.Panic(err)
	}

	return &LiDaYe{&daYe{name: "李大爷", total: 100000, trans: trans}}
}

func (l *LiDaYe) Say(wg *sync.WaitGroup) {
	defer wg.Done()

	i := uint32(0)

	go func() {
		for {
			time.Sleep(3 * time.Second)
			log.Println(l.name+" 说: ", i)
		}
	}()

	for i < l.total {
		err := l.trans.Send(msgTypePlaceLReq)
		if err != nil {
			log.Panic(err)
		}
		if err = l.trans.Send(msgTypePleaseLReq); err != nil {
			log.Panic(err)
		}
		i++
	}
}

func (l *LiDaYe) Listen(wg *sync.WaitGroup) {
	defer wg.Done()

	i := uint32(0)

	go func() {
		for {
			time.Sleep(3 * time.Second)
			log.Println(l.name+" 听: ", i)
		}
	}()

	for i < l.total*3 {
		msg, err := l.trans.Receive()
		if err != nil {
			log.Panic(err)
		}

		typ := msg.Type()
		switch typ {
		case msgTypeEatZReq:
			if err = l.trans.Send(msgTypeEatLRes); err != nil {
				log.Panic(err)
			}
		case msgTypePlaceZRes, msgTypePleaseZRes:
		default:
			log.Panicf("听不懂的消息类型: %d", typ)
		}
		i++
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(4)

	ch := make(chan bool)

	var li DaYe
	go func() {
		time.Sleep(2 * time.Second)
		li = NewLiDaYe()
		ch <- true
	}()

	zhang := NewZhangDaYe()
	<-ch

	go li.Listen(wg)
	go zhang.Listen(wg)
	zhang.Say(wg)
	li.Say(wg)

	wg.Wait()
}
