package main

import (
    "github.com/chuyval/qqs-helper/q2/pkg/models"
    "github.com/golang/protobuf/proto"
    "fmt"
    "log"
)

func main () {
    stuff := createProtoStuff()
    log.Printf("Stuff: %+v", stuff)

    // Convert to string
    stuffStr := proto.MarshalTextString(stuff)

    var stuff2 models.Stuff
    err := proto.UnmarshalText(stuffStr, &stuff2)
    if err != nil {
        fmt.Println("It didnt work", err)
    } else {
        fmt.Println("It worked", stuff2)
    }
}

func createProtoStuff() *models.Stuff {
    someValueList := []*models.SomeValue{&models.SomeValue{Id: "Some Value ID"}}

    valueList := &models.SomeValueList{SomeValue: someValueList}

    stuffValueList := &models.Stuff_ValueList{
        ValueList: valueList,
    }

    stuff := &models.Stuff{
        Id: "Stuff List Id",
        Stuff: stuffValueList,
    }

    return stuff
}