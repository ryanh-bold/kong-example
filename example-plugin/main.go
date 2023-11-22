package main

import (
	"fmt"
	"sync/atomic"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

const Version = "1.0.0"
const Priority = 1000

func main() {
	server.StartServer(New, Version, Priority)
}

var newCalls atomic.Uint64

func New() interface{} {
	newCalls.Add(1)
	return &Config{}
}

type Config struct{}

func (c *Config) Access(kong *pdk.PDK) {
	kong.Log.Debug(fmt.Sprintf("New() has been called %d times", newCalls.Load()))
	kong.Response.SetHeader("X-Example-Message", "Hello World")
}
