package collections

import "reflect"

// CannonFire is a function that the cannon disposes to during a fire session.
type CannonFire func(reflect.Type, interface{})

// DoneFunc is run after the fire session has been sent.
// There's NO guarantee that all the data has been processed.
// You'll have to handle that yourself.
type DoneFunc func()

// DataCannon is an wrapper ontop of a linked list, it accept linked lists.
// When fired it tries to dispose the data as quickly as possible into the specified CannonFire goroutine.
// Remember to mutex lock these if you do alot of bursts!
// When all the shots have been fire DoneFunc will be called.
type DataCannon struct {
    buffers LinkedList
}

// Load loads a burst in to the data cannon.
func (c *DataCannon) Load(burst LinkedList) {
    c.buffers.PushBack(burst)
}

// Fire fires a single burst of data to X amount of goroutines with the specified delegate.
// omitting the times variable will make the cannon only fire once.
func (c *DataCannon) Fire(fireDelegate CannonFire, doneFunc DoneFunc, times ...int) {
    t := 1
    if len(times) > 0 {
        t = times[0]
    }
    for i := 0; i < t; i++ {
        _, buffp := c.buffers.Get(0)
        buff := buffp.(LinkedList)
        for buff.Next() != nil {
            te, v := buff.Current()
            go fireDelegate(te, v)
        }
        c.buffers.Remove(0)
    }
    doneFunc()
}
