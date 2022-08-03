package main

import "fmt"

type Base struct {
    num int
}

func (b Base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

type Container struct {
    Base
    Str string
}

func main() {

    item := Container{
        Base: Base{
            num: 1,
        },
        Str: "some name",
    }

    fmt.Printf("co={num: %v, str: %v}\n", item.num, item.Str)

    fmt.Println("also num:", item.Base.num)

    fmt.Println("describe:", item.describe())

    type Describe interface {
        describe() string
    }

    var d Describe = item
    fmt.Println("describer:", d.describe())
}