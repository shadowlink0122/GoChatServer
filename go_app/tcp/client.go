package main

import(
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

const(
	DISCONNECT = iota
	SEED

	ELSE
)

func Input(seed *int)string{
	var text = bufio.NewScanner(os.Stdin)
	for{
		text.Scan()
		if(text.Text() != ""){
			break
		}
	}

	enc_text := enc(text.Text(), *seed)
	return enc_text
}

func getSeed(conn net.Conn)(int){
	seed, _ := strconv.Atoi(GetMessage(conn))
	return seed
}

func SendMessage(conn net.Conn, seed *int)(net.Conn){
	text := Input(seed)
	conn.Write([]byte(text))
	return conn
}

func GetMessage(conn net.Conn)(string){
	response := make([]byte, 1024)
  n,_ := conn.Read(response)
  message := string(response[:n])

  return message
}

func PrintMessage(conn net.Conn)(string){
	response := make([]byte, 1024)
  n,_ := conn.Read(response)
  message := string(response[:n])

  fmt.Println(message)
  return message
}

func SetMyName(conn net.Conn, seed *int){
	fmt.Printf("YourName: ")
	SendMessage(conn, seed)
}

func system_message(conn net.Conn, sys_msg string)(int){
	switch sys_msg{
	case "Disconnection":
		return DISCONNECT
	case "":
		return DISCONNECT

	default:
		return ELSE
	}

	return (ELSE) // Else
}

func chatting(conn net.Conn, seed *int){
	for{
		go func(){
			SendMessage(conn, seed)
		}()

  	message := GetMessage(conn)
  	message = dec(message, *seed)

  	switch system_message(conn, message){
  	case DISCONNECT:
  		fmt.Println("<Disconnected!>")
  		return
  	default:
  		fmt.Println(message)
  	}
	}
}

func main() {
  conn, _ := net.Dial("tcp", "192.168.33.10:8000")
  msg_seed := getSeed(conn)
  SetMyName(conn, &msg_seed)
  chatting(conn, &msg_seed)
}

