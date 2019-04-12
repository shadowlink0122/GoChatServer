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

//入力と、その文字列を暗号化して渡す
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

//暗号化のためのシードをもらう
func getSeed(conn net.Conn)(int){
	seed, _ := strconv.Atoi(GetMessage(conn))
	return seed
}

// メッセージの送信
func SendMessage(conn net.Conn, seed *int)(net.Conn){
	text := Input(seed)
	conn.Write([]byte(text))
	return conn
}

//メッセージの受信
func GetMessage(conn net.Conn)(string){
	response := make([]byte, 1024)
	n,_ := conn.Read(response)
	message := string(response[:n])

	return message
}

//受信したメッセージを表示する
func PrintMessage(conn net.Conn)(string){
	response := make([]byte, 1024)
	n,_ := conn.Read(response)
	message := string(response[:n])

	fmt.Println(message)
	return message
}

//識別のため、自分の名前を入力する
func SetMyName(conn net.Conn, seed *int){
	fmt.Printf("YourName: ")
	SendMessage(conn, seed)
}

//受信したメッセージが、ユーザのものでなく、サーバのものである時
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

//チャットのメインコード
func chatting(conn net.Conn, seed *int){
	for{
		go func(){
			for{
				SendMessage(conn, seed)	
			}
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

