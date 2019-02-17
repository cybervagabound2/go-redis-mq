// Basic implementation of Publisher/Subscriber pattern
package main

import (
    "log"
    "sync"
    "time"
    "strconv"

    "github.com/garyburd/redigo/redis"
)

// Our main struct that implements all stuff.
type processor struct {
    mu    sync.Mutex
    psc   redis.PubSubConn
    pool  *redis.Pool
    topic string
}

// Handling errors
func (p *processor) forceError() {
    p.mu.Lock()
    if p.psc.Conn != nil {
         p.psc.Conn.Send("QUIT")
         p.psc.Conn.Flush()
    }
    p.mu.Unlock()
}

// Function allows to listen for Publisher
func (p *processor) listen() {
main:
    for {
        p.mu.Lock()
        if p.psc.Conn != nil {
            p.psc.Conn.Close()
            p.psc.Conn = nil
        }
        log.Println("new connection")
        p.psc.Conn = p.pool.Get()
        if err := p.psc.Subscribe(p.topic); err != nil {
            log.Printf("Subscribe(%s) returned %v", p.topic, err)
            continue
        }
        p.mu.Unlock()
        for {
            switch v := p.psc.Receive().(type) {
            case redis.Message:
                log.Printf("incoming message, %s %s\n", v.Channel, v.Data)
            case error:
                log.Printf("Receive() returned %v", v)
                continue main
            }
        }
    }
}

func main() {
    pool := &redis.Pool{Dial: func() (redis.Conn, error) {
        return redis.Dial("tcp", ":6379")
    },
    }

    p := &processor{pool: pool, topic: "topic:"}
    go p.listen()

    c := pool.Get()
    defer c.Close()

    for i := 0; i < 10; i++ {
        time.Sleep(time.Second)
        c.Do("PUBLISH", p.topic, strconv.Itoa(i) + ":( ͡° ͜ʖ ͡°)")
        time.Sleep(time.Second)
    }
    // check that we can force error
    p.forceError()
    time.Sleep(time.Second)
    // continue messaging
    c.Do("PUBLISH", p.topic, "Придется применить DDos-атаку")
    time.Sleep(time.Second)
}

