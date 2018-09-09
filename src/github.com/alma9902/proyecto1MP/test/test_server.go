package test

import(
  "testing"
  "github.com/alma9902/proyecto1MP/programs"
)

func startServerTest(t *testing.T){
  c1 := programs.ClientManager{
    clients:   make(map[*Client]bool),
    broadcast: make(chan []byte),
    register:  make(chan *Client),
    unregister:make(chan *Client),
  }
  listener, error:= net.Listen("tcp" ,":12345")
  defer listener.Close()
  connection,_:=listener.Accept()
  c1.register <- programs.Client{
    socket:connection,
    data: "Hola mundo",
    name: "Alma",
  }
  output := c1.programs.Start()
  expected := "La conexión salió bien"
  if output != expected{
    t.Errorf("expected %s, encontrado %s",expected,output)
  }
}
