package collections

import "reflect"
import "errors"
import "sync"

type LinkedBuffer struct {
    // BuffName is the name of the buffer.
    BuffName string

    BuffLimit int
    buffCont int

    capLock bool
    asyncLock sync.Mutex

    // ll is the backend linked list.
    ll LinkedList
}

func (lb *LinkedBuffer) Read() (itype reflect.Type, item interface{}) {
    t, i := lb.ll.Get(0)
    lb.ll.Remove(0)
    lb.buffCont--
    return t, i
}

func (lb *LinkedBuffer) Write(item interface{}) error {
    if lb.BuffLimit > 0 && lb.buffCont > lb.BuffLimit {
        return errors.New("Item does not fit in the buffer due to buffer capping.")
    }
    lb.ll.PushBack(item)
    lb.buffCont++
    return nil
}

func (lb *LinkedBuffer) readasync() (itype reflect.Type, item interface{}) {
    lb.asyncLock.Lock()
    defer lb.asyncLock.Unlock()
    t, i := lb.ll.Get(0)
    if t != nil {
        lb.ll.Remove(0)
        lb.buffCont--
    }
    return t, i
}

func (lb *LinkedBuffer) ReadAsync() (itype reflect.Type, item interface{}) {
    t, i := lb.readasync()
    for t == nil {
        t, i = lb.readasync()
    }
    lb.capLock = false
    return t, i
}

func (lb *LinkedBuffer) WriteAsync(item interface{}) {
    if lb.BuffLimit > 0 && lb.buffCont > lb.BuffLimit-1 {
        lb.capLock = true
    }

    for lb.capLock == true {}

    lb.asyncLock.Lock()
    defer lb.asyncLock.Unlock()
    lb.ll.PushBack(item)
    lb.buffCont++
}
