package main

import (
    "os"
    "flag"
    "strings"
)

func main() {
    flagMode := flag.String("mode", "server", "start in clien or server")
    flag.Parse()
    if strings.ToLower(*flagMode) == "server"{
      startServerMode((os.Args[3]))
    }else{
      startClientMode((os.Args[3]),(os.Args[4]),(os.Args[5]))
    }
}
