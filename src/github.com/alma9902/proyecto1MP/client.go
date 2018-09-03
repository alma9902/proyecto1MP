package main

import(
  "bufio"
  "fmt"
  "net"
  "os"
  "strings"
)

type Client struct {
    socket net.Conn
    data   chan []byte
    name string
}
const(
  StopCharacter= "\r\n\r\n"
)
func (client *Client) receive() {
    for {
        message := make([]byte, 4096)
        length, err := client.socket.Read(message)
        if err != nil {
            client.socket.Close()
            break
        }
        if length > 0 {
            fmt.Println("says " + string(message))
        }
    }
}
func startClientMode(ip string, port string, name string) {
    fmt.Println("Conectando usuario...")
    addr := strings.Join([]string{ip,port}, ":")
    connection, error := net.Dial("tcp", addr)
    if error != nil {
        fmt.Println(error)
    }
    client := &Client{name: name, socket: connection}
    //client := &Client{socket: connection}
    fmt.Println(client.name)
    go client.receive()
    for {
        reader := bufio.NewReader(os.Stdin)
        message, _ := reader.ReadString('\n')
        connection.Write([]byte(strings.TrimRight(message, "\n")))
    }
}
