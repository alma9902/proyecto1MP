package programs

import(
  "bufio"
  "fmt"
  "net"
  "os"
  "strings"
)

type Client struct {
    Socket net.Conn
    Data   chan []byte
    Name   string
}

func (client *Client) Receive() {
    for {
        message := make([]byte, 4096)
        length, err := client.Socket.Read(message)
        if err != nil {
            client.Socket.Close()
            break
        }
        if length > 0 {
            fmt.Println("Yo :"+string(message))
        }
    }
}
func StartClientMode(ip string, port string, name string) {
    fmt.Println("Conectando usuario...")
    addr := strings.Join([]string{ip,port}, ":")
    connection, error := net.Dial("tcp", addr)
    if error != nil {
        fmt.Println(error)
    }
    client := &Client{Name: name, Socket: connection}
    fmt.Println(client.Name)
    go client.Receive()
    for {
        reader := bufio.NewReader(os.Stdin)
        message, _ := reader.ReadString('\n')
        connection.Write([]byte(strings.TrimRight(message, "\n")))
    }
}
