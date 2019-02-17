// Basic implementation of Publisher/Subscriber pattern
package main

import (
    "log"
    "sync"
    "time"

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
    if p.psc/Conn != nil {
        p.psc.Conn != nil {
            p.psc.Conn.Send("QUIT")
            p.psc.Conn.Flush()
    }
    p.mu.Unlock()
}
