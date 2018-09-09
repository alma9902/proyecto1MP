package main

import (
    "fmt"
    "net"
    "os"
    "log"
)

type ClientManager struct {
    //map[KeyType] valueType
    clients    map[*Client]bool
    //names      map[*Client]string
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
}

func (manager *ClientManager) start() string {
    for {
        select {
        case connection := <-manager.register:
            manager.clients[connection] = true
            fmt.Println("Added new connection!")
            return "La conexión salió bien"
            //fmt.Println("Added new connection!"+ manager.names[*client.name])
        case connection := <-manager.unregister:
            if _, ok := manager.clients[connection]; ok {
                close(connection.data)
                delete(manager.clients, connection)
                fmt.Println("A connection has terminated!")
            }
        case message := <-manager.broadcast:
            for connection := range manager.clients {
                select {
                case connection.data <- message:
                default:
                    close(connection.data)
                    delete(manager.clients, connection)
                }
            }
        }
    }
}

func (manager *ClientManager) receive(client *Client) {
    for {
        message := make([]byte, 4096)
        length, err := client.socket.Read(message)
        if err != nil {
            manager.unregister <- client
            client.socket.Close()
            break
        }
        if length > 0 {
            log.Println("RECEIVED :" + string(message))
            manager.broadcast <- message
        }
    }
}

func (manager *ClientManager) send(client *Client) {
    defer client.socket.Close()
    for {
        select {
        case message, ok := <-client.data:
            if !ok {
                return
            }
            client.socket.Write(message)
        }
    }
}

func startServerMode(port string) {
    fmt.Println("Starting server...")
    listener, error := net.Listen("tcp",":"+port)
    defer listener.Close()
    if error != nil {
        log.Fatal("socket listen port %s failed,%s", port, error)
        os.Exit(1)
    }
    log.Printf("Begin listen port : %s",port)
    manager := ClientManager{
        clients:    make(map[*Client]bool),
        broadcast:  make(chan []byte),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
    go manager.start()
    for {
        connection, _ := listener.Accept()
        if error != nil {
            log.Fatalln(error)
            continue
        }
        client := &Client{socket: connection, data: make(chan []byte)}
        manager.register <- client
        go manager.receive(client)
        go manager.send(client)
    }
}
