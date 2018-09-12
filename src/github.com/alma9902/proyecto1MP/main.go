
package main

import (
    "bufio"
    "fmt"
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
      reader := bufio.NewReader(os.Stdin)
      // despliega las opciones que tiene el cliente
      fmt.Println(programs.ShowOptions())
      input, _ := reader.ReadString('\n')
      words := strings.Split(input, " ")
      if words[0] == "IDENTIFY" && words[1] != "" && words[1] != " "{
          programs.StartClientMode((os.Args[3]),(os.Args[4]),words[1])
      }else{
        os.Exit(3)
      }
    }
}
