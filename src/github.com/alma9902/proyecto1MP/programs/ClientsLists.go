package programs

import(
    "container/list"
    "fmt"
    "bytes"
)
type ClientsList struct{
    AllClients *list.List
}
func (e *ClientsList) AddToList(cliente *Client){
  e.AllClients.PushBack(cliente)
}
/*func (e *ClientsList) DeleteClient(cliente *Client){
  e.AllClients.Remove(cliente)
}*/
func (e *ClientsList) ShowClients() string{
  var cad bytes.Buffer
  for temp := e.AllClients.Front(); temp != nil; temp = temp.Next(){
    fmt.Println(temp.Value.(*Client).Name+"id :"+temp.Value.(*Client).Id)
    cad.WriteString( temp.Value.(*Client).Name + "\n")
  }
  return (cad.String())
}
func (e *ClientsList) SearchClient(id string) bool{
  for temp := e.AllClients.Front(); temp != nil; temp = temp.Next(){
    if(temp.Value.(*Client).Id == id){
      return true
    }
  }
  return false
}
