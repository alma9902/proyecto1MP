
package main

import (
    //"bufio"
    //"fmt"
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
      //reader := bufio.NewReader(os.Stdin)
      programs.StartClientMode((os.Args[3]),(os.Args[4]))

    }
}
