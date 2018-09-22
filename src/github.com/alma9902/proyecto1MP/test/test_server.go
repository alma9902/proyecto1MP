package test

import(
  "testing"
  "github.com/alma9902/proyecto1MP/programs"
)

func startServerTest(t *testing.T){
  c1 := programs.ClientManager{
    Clients:   make(map[*Client]bool),
    Broadcast: make(chan []byte),
    Register:  make(chan *Client),
    Unregister:make(chan *Client),
  }
  listener, error:= net.Listen("tcp" ,":12345")
  defer listener.Close()
  connection,_:=listener.Accept()
  c1.register <- programs.Client{
    Socket:connection,
    Data: "Hola mundo",
    Name: "Alma",
  }
  output := c1.programs.Start()
  expected := "La conexión salió bien"
  if output != expected{
    t.Errorf("expected %s, encontrado %s",expected,output)
  }
}
