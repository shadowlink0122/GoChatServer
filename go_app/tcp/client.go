package main

import(
	"bufio"
	"fmt"
	"net"
	"os"
)

func Input()string{
	var text = bufio.NewScanner(os.Stdin)
	for{
		text.Scan()
		if(text.Text() != ""){
			break
		}
	}
	return text.Text()
}

func SendMessage(conn net.Conn)(net.Conn){
	text := Input()
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

func SetMyName(conn net.Conn)(net.Conn){
	fmt.Printf("YourName: ")
	SendMessage(conn)

	return conn
}

func chatting(conn net.Conn){
	for{
		go func(){
			conn = SendMessage(conn)
		}()

  	message := GetMessage(conn)

  	if(message == "Disconnection" || message == ""){
  		break
  	}else{
  		fmt.Println(message)
  	}
	}
}

func main() {
  conn, _ := net.Dial("tcp", "IPAdress:Port")
  conn = SetMyName(conn)
  chatting(conn)
}

