package programs

// Esta clase se encargará de validar las líneas que pase
// el usuario, también de mandar los parámetros necesarios
// para que lo reciba ActionsKeyWords de enumMessenger
import(
  "strings"
  "github.com/rs/xid"
  "fmt"
)
//Esta función se encargará de analizar las cadenas, mandarlas
//al enum y no devuelve nada, las demás funciones
//se encargarán de eso
//Recibe el mensaje tipo cadena
func genXid() string{
    id := xid.New()
    return id.String()
}
func Actions(messa string, client *Client, listaClientes ClientsList)string{
  cad := strings.Fields(messa)
  id := cad[0]
  key := cad[1]
  //mensaje sin la palabra identificadora
  //message := strings.Replace(messa, " ",1)
  if(key == "IDENTIFY" && len(cad) == 3){
      client := &Client{Name: cad[2], Id: id}
      listaClientes.AddToList(client)
      fmt.Println(listaClientes.ShowClients())
      return "Se ha identificado exitosamente : "+client.Name
  }else{
    //Aquí busco al cliente en la lista de registrados por su id
    //si el id del cliente no pertenece a la lista entonces no
    //se ha identificado
    if(! listaClientes.SearchClient(id)){
      return ShowOptions()
    }else{
      //ActionsKeyWords(key, message)
      return"ok"
    }
  }
  return ShowOptions()
}
