package programs
/**
  * Esta clase checará qué se recibe como identificador
  * en los mensajes y hará lo correspondiente
  * tengo que checar también que no haga nada más mientras no se identifica
  *antes tengo que descomponer los mensajes y verificar que sean válidos
  **/
import (
  "fmt"
)
//Esto retornará el mensaje con las opciones que el usuario puede ingresar
//para activar las diferentes funciones
func ShowOptions()string{
  var s string =
  "Lo que tiene que saber : \n"+
  "DEBE IDENTIFICARSE PRIMERO, DE LO CONTRARIO NO PODRÁ SEGUIR CONECTADO"+
  "ESCRIBA LAS SIGUIENTES PALABRAS CLAVE, AL PRINCIPIO PARA : \n"+
  "IDENTIFY      : para que te identifiques como usuario\n"+
  "STATUS        : para poner tu estado al usuario\n"+
  "USERS         : mostrar los usuarios identificados\n"+
  "MESSAGE       : enviar mensaje, el formato es MESSAGE username messageContent\n"+
  "PUBLICMESAGGE : enviar mensaje a todos los usuarios identificados, formato : PUBLICMESAGGE message\n"+
  "CREATEROOM    : crear sala donde los usuarios invitados puedan comunicarse\n"+
  "INVITE        : invitar a usuario a la sala, formato: INVITE username1 username2\n"+
  "JOINROOM      : unirse a la sala\n"+
  "ROOMESSAGE    : enviar mensaje a sala de usuarios: ROOMESSAGE roomname messageContent\n"+
  "DISCONNECT    : desconectarse del servidor"
  return s
}
type KeyWords int

const(
  //Season_Winter Season = 0
  IDENTIFY      KeyWords= 0
  STATUS        KeyWords= 1
  USERS         KeyWords= 2
  MESSAGE       KeyWords= 3
  PUBLICMESAGGE KeyWords= 4
  CREATEROOM    KeyWords= 5
  INVITE        KeyWords= 6
  JOINROOM      KeyWords= 7
  ROOMESSAGE    KeyWords= 8
  DISCONNECT    KeyWords= 9
)

func ActionsKeyWords(key KeyWords) string{
  switch (key) {
  case IDENTIFY:
    return "ok, identifica el usuario"
  case STATUS:
    return "ok, tengo que dar el status"
  case USERS:
    return "ok, users"
  case MESSAGE:
    //aquí conecto al cliente para que mande mensaje
    return "vale, message"
  case PUBLICMESAGGE:
    return "tiene que ser mensaje público"
  case CREATEROOM:
    return "tenemos que crear una sala"
  case INVITE:
    return "invite wuuuu"
  case JOINROOM:
    return "JOINROOM"
  case ROOMESSAGE:
    return "ROOMESSAGE"
  case DISCONNECT:
    //aquí de alguna manera tendré que usar ctrl c
    return "DISCONNECT"
  default:
    fmt.Println("Por favor, ingrese una opción válida")
    return ShowOptions()
  }
}
