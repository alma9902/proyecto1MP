package programs

import (
    "fmt"
    "net"
    "os"
    "log"
)

type ClientManager struct {
    //map[KeyType] valueType
    Clients    map[*Client]bool
    Broadcast  chan []byte
    Register   chan *Client
    Unregister chan *Client
}

func (manager *ClientManager) Start(){
    for {
        select {
        case connection := <-manager.Register:
            manager.Clients[connection] = true
            fmt.Println("Added new connection!")
        case connection := <-manager.Unregister:
            if _, ok := manager.Clients[connection]; ok {
                close(connection.Data)
                delete(manager.Clients, connection)
                fmt.Println("A connection has terminated!")
            }
        case message := <-manager.Broadcast:
            for connection := range manager.Clients {
                select {
                case connection.Data <- message:
                default:
                    close(connection.Data)
                    delete(manager.Clients, connection)
                }
            }
        }
    }
}

func (manager *ClientManager) Receive(client *Client) {
    for {
        message := make([]byte, 4096)
        length, err := client.Socket.Read(message)
        if err != nil {
            manager.Unregister <- client
            client.Socket.Close()
            break
        }
        if length > 0 {
            log.Println("RECEIVED from :" +string(message))
            manager.Broadcast <- message
        }
    }
}

func (manager *ClientManager) Send(client *Client) {
    defer client.Socket.Close()
    for {
        select {
        case message, ok := <-client.Data:
            if !ok {
                return
            }
            client.Socket.Write(message)
        }
    }
}

func StartServerMode(port string) {
    fmt.Println("Starting server...")
    listener, error := net.Listen("tcp",":"+port)
    defer listener.Close()
    if error != nil {
        log.Fatal("socket listen port %s failed,%s", port, error)
        os.Exit(1)
    }
    log.Printf("Begin listen port : %s",port)
    manager := ClientManager{
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan []byte),
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
    }
    go manager.Start()
    for {
        connection, _ := listener.Accept()
        if error != nil {
            log.Fatalln(error)
            continue
        }
        client := &Client{Socket: connection, Data: make(chan []byte)}
        manager.Register <- client
        go manager.Receive(client)
        go manager.Send(client)
    }
}
