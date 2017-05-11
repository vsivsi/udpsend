package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Host string `long:"host" default:"255.255.255.255" description:"IP destination address"`
	Port uint16 `long:"port" default:"1234" description:"UDP destination port"`
	Msg  string `long:"msg" default:"" description:"data to send"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			log.Printf("error: %v\n", err.Error())
			os.Exit(1)
		} else {
			// log.Printf("%v\n", err.Error())
			os.Exit(0)
		}
	}

	c, err := net.Dial("udp", fmt.Sprintf("%v:%d", opts.Host, opts.Port))
	if err != nil {
		log.Panic(err)
	}

	defer c.Close()

	var output []byte

	if len(opts.Msg) == 0 {
		output, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Panic(err)
		}
	} else {
		output = []byte(opts.Msg)
	}

	n, err := c.Write(output)
	if err != nil {
		log.Panic(err)
	} else {
		log.Printf("Wrote %d bytes\n", n)
	}

	os.Exit(0)
}
