package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type User struct {
	RegAt time.Time // 注册时间
	Score int       // 积分
}

var (
	users = make([]*User, 0, 500)
	mu = sync.Mutex{}
	listenNumber int32      //监听全局变量users的次数
	prized       bool       //是否已经执行过积分奖励
	downstreamOver bool
)

const (
	H = 3
	M = 8
	N = 5
)

func prize() {
	sort.Slice(users, func(i, j int) bool { return users[i].RegAt.Before(users[j].RegAt) })
	for _, user := range users[:H] {
		user.Score += 1
	}
}

func main() {
	// BusinessModel()
	// SignalWithChannel()
	// SignalWithCond()
	// BroadcastWithChannel()
	BroadcastWithCond()
	fmt.Println("listenNumber: ", listenNumber)
}

func BusinessModel() {
	//上游
	for i := 0; i < M; i++ {
		go func() {
			for { //不停地注册新用户
				if downstreamOver {
					break
				}
				mu.Lock()
				users = append(users, &User{RegAt: time.Now()}) //注册用户
				mu.Unlock()
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) //随机休息一段时间，再注册下一个用户
			}
		}()
	}

	//下游
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			for {
				mu.Lock()
				if !prized {
					atomic.AddInt32(&listenNumber, 1)
					if len(users) >= 100 {
						prize()
						prized = true
					}
				}
				mu.Unlock()
				if prized {
					break
				}
			}
		}()
	}
	wg.Wait()
	downstreamOver = true
}

// 减少对全局变量users的监听次数。上游每次改变users时向一个channel里发送一条数据
func SignalWithChannel() {
	ch := make(chan struct{}, 10*N)

	//上游
	for i := 0; i < M; i++ {
		go func() {
			for { //不停地注册新用户
				if downstreamOver {
					break
				}
				mu.Lock()
				users = append(users, &User{RegAt: time.Now()}) //注册用户
				mu.Unlock()
				ch <- struct{}{}
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) //随机休息一段时间，再注册下一个用户
			}
		}()
	}

	//下游
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			for {
				<-ch //阻塞，直到users有改变
				mu.Lock()
				if !prized {
					atomic.AddInt32(&listenNumber, 1)
					if len(users) >= 100 {
						prize()
						prized = true
					}
				}
				mu.Unlock()
				if prized {
					break
				}
			}
		}()
	}
	wg.Wait()
	downstreamOver = true
}

func SignalWithCond() {
	cond := sync.NewCond(&mu) //cond.L等价于mu

	//上游
	for i := 0; i < M; i++ {
		go func() {
			for { //不停地注册新用户
				if downstreamOver {
					break
				}
				mu.Lock()
				users = append(users, &User{RegAt: time.Now()}) //注册用户
				mu.Unlock()
				cond.Signal()                                                //通知别人users有变化。Signal只能通知到一个协程
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) //随机休息一段时间，再注册下一个用户
			}
		}()
	}

	//下游
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			for {
				mu.Lock()   //等价于cond.L.Lock()
				cond.Wait() //阻塞，直到接收到通知。Wait内部会先执行mu.Unlock()，等接收到信号后再执行mu.Lock()，所以在调Wait()之前需要先上锁
				if !prized {
					atomic.AddInt32(&listenNumber, 1)
					if len(users) >= 100 {
						prize()
						prized = true
					}
				}
				mu.Unlock() //等价于cond.L.Unlock()
				if prized {
					break
				}
			}
		}()
	}
	wg.Wait()
	downstreamOver = true
}

func BroadcastWithChannel() {
	ch := make(chan struct{}, 10*N)

	//上游
	for i := 0; i < M; i++ {
		go func() {
			for { //不停地注册新用户
				if downstreamOver {
					break
				}
				mu.Lock()
				users = append(users, &User{RegAt: time.Now()}) //注册用户
				mu.Unlock()
				//把n个下游协程全部通知一遍。close channel也能实现通知的功能，但是一个channl只能close一次，本业务中我们需要多次通知。实际中上游一般不知道下游协程的数目，这种情况下只能用cond.Broadcast()
				for j := 0; j < N; j++ {
					ch <- struct{}{}
				}
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) //随机休息一段时间，再注册下一个用户
			}
		}()
	}

	//下游
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			for {
				<-ch //阻塞，直到users有改变
				atomic.AddInt32(&listenNumber, 1)
				mu.Lock()
				done := false
				if len(users) >= 100 {
					prize()
					done = true
				}
				mu.Unlock()
				if done {
					break
				}
			}
		}()
	}
	wg.Wait()
	downstreamOver = true
}

func BroadcastWithCond() {
	cond := sync.NewCond(&mu) //cond.L等价于mu

	//上游
	for i := 0; i < M; i++ {
		go func() {
			for { //不停地注册新用户
				if downstreamOver {
					break
				}
				mu.Lock()
				users = append(users, &User{RegAt: time.Now()}) //注册用户
				mu.Unlock()
				cond.Broadcast()                                             //通知所有下游协程
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) //随机休息一段时间，再注册下一个用户
			}
		}()
	}

	//下游
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			for {
				mu.Lock()
				cond.Wait()
				atomic.AddInt32(&listenNumber, 1)
				done := false
				if len(users) >= 100 {
					prize()
					done = true
				}
				mu.Unlock()
				if done {
					break
				}
			}
		}()
	}
	wg.Wait()
	downstreamOver = true
}