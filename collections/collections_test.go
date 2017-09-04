package collections

import (
    "testing"
    "strconv"
    "time"
    "reflect"
    "fmt"
)

func TestLinkedListCreate(t *testing.T) {
    l := NewLinkedList("FOO")
    l.PushBack("BAR")
    _, a := l.Get(0)
    _, b := l.Get(1)
    if a.(string) == "FOO" && b.(string) == "BAR" {
        return;
    }
    t.Error("Expected FOO and BAR, got: " + a.(string) + " and " + b.(string) + "!")
}

func TestLinkedListIterate(t *testing.T) {
    l := NewLinkedList("START")
    for i := 1; i <= 100; i++ {
        l.PushBack(i)
    }
    for l.Next() != nil { }
}

func TestLinkedListMultiType(t *testing.T) {
    l := NewLinkedList("START")
    for i := 1; i <= 50; i++ {
        l.PushBack(i)
        l.PushBack("STR_"+strconv.Itoa(i))
    }
    for l.Next() != nil {
        t, i := l.Current()
        switch reflect.TypeOf(t) {
        case reflect.TypeOf(""):
            fmt.Println(t.String() + ": " + i.(string))
        case reflect.TypeOf(1):
            fmt.Println(t.String() + ": " + strconv.Itoa(i.(int)))
        default:
        }
    }
}

func TestLinkedBufferSync(t *testing.T) {
    b := new(LinkedBuffer)
    err := b.Write("Test")
    if err != nil {
        t.Error(err)
    }
    _, v := b.Read()
    if v.(string) != "Test" {
        t.Error("Expected Test, got: " + v.(string) + "!")
    }
}

func TestLinkedBufferAsync(t *testing.T) {
    b := new(LinkedBuffer)
    b.BuffLimit = 2
    go func() {
        for i := 0; i < 10; i++ {
            b.WriteAsync("BUFF_O_"+strconv.Itoa(i))
        }
        fmt.Println("Writer is done!")
    }()
    for i := 0; i < 10; i++ {
        t, i := b.ReadAsync()
        fmt.Println("Reader | type: " + t.String() + " | value: " + i.(string))
        time.Sleep(100 * time.Millisecond)
    }
}


func TestLinkedBufferAsyncNoLimit(t *testing.T) {
    b := new(LinkedBuffer)
    go func() {
        for i := 0; i < 10; i++ {
            b.WriteAsync("BUFF_O_"+strconv.Itoa(i))
        }
        fmt.Println("Writer is done!")
    }()
    for i := 0; i < 10; i++ {
        t, i := b.ReadAsync()
        fmt.Println("Reader | type: " + t.String() + " | value: " + i.(string))
        time.Sleep(100 * time.Millisecond)
    }
}

// TestDataCannon tests if the data cannon fires any values. Some will be missed due to the nature of the datacannon.
// Its objective is to fire as much data into the goroutines as possible, as fast as possible.
func TestDataCannon(t *testing.T) {
    dc := new(DataCannon)
    a := []string{"FOO", "BAR"}
    b := []string{"BAZ", "X", "Y", "Z", "W"}
    c := []string{"OKKER", "GOKKER", "GUMMI", "KLOKKER"}
    helper := func(a []string) LinkedList {
        l := new(LinkedList)
        for _, i := range a {
            l.PushBack(i)
        }
        return *l
    }
    ll := helper(a)
    ll2 := helper(b)
    ll3 := helper(c)
    dc.Load(ll)
    dc.Load(ll2)
    dc.Load(ll3)
    done := false
    dc.Fire(func (t reflect.Type, v interface{}) {
        fmt.Println(t.String() + "... " + v.(string))
    }, func() {
        done = true
    }, 3)
    for !done {}
}
