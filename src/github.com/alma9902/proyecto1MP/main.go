
package main

import (
    "os"
    "flag"
    "strings"
    "github.com/alma9902/proyecto1MP/programs"
)

func main() {
    flagMode := flag.String("mode", "server", "start in clien or server")
    flag.Parse()
    if strings.ToLower(*flagMode) == "server"{
      programs.StartServerMode((os.Args[3]))
    }else{
      programs.StartClientMode((os.Args[3]),(os.Args[4]),(os.Args[5]))
    }
}
