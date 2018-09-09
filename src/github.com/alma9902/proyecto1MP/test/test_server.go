package testProyect

import(
  "testing"
  //"github.com/alma9902/proyecto1MP/programs/main"
  package main
)

func startServerTest(t *testing.T){
  c1 := ClientManager{
    clients:   make(map[*Client]bool),
    broadcast: make(chan []byte),
    register:  make(chan *Client),
    unregister:make(chan *Client),
  }
  listener, error:= net.Listen("tcp" ,":12345")
  defer listener.Close()
  connection,_:=listener.Accept()
  manager.register <- Client{
    socket:connection,
    data: "Hola mundo",
    name: "Alma",
  }
  output := c1.start()
  expected := "La conexión salió bien"
  if output != expected{
    t.Errorf("expected %s, found %s",expected,output)
  }
}
