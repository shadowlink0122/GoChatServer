package main

import(
	"fmt"
	"net"
)

type User struct{
	IPAdress string
	UserName string
	conn net.Conn
}

func SendMessage(conn net.Conn, msg string)(net.Conn){
	conn.Write([]byte(msg))
	return conn
}

func GetMessage(conn net.Conn)(string){
	response := make([]byte, 1024)
  n,_ := conn.Read(response)
  message := string(response[:n])

  return message
}

func PrintMessage(conn net.Conn){
  message := GetMessage(conn)

  fmt.Println(message)
}

func GetPrintMessage(conn net.Conn)(string){
	message := GetMessage(conn)

	fmt.Println(message)
	return message
}

func PrintUsersData(pUserData *[]User){
	for index, UserData := range *pUserData{
		fmt.Printf("[%d] %s : %s\n", index, UserData.IPAdress, UserData.UserName)
	}
}

func SendForAll(conn net.Conn, msg string, pUserData *[]User, name string, IPAddr string){
	msg = fmt.Sprintf("[%s]:%s", name, msg)
	for _, UserData := range *pUserData{
		if(UserData != User{IPAddr, name, conn}){
			UserData.conn = SendMessage(UserData.conn, msg)
		}
	}
}

func DeleteUser(conn net.Conn, pUserData *[]User, name string, IPAddr string)([]User){
	var newUserData []User
	for _, UserData := range *pUserData{
		if(UserData != User{IPAddr, name, conn}){
			newUserData = append(newUserData, UserData)
		}
	}
	return newUserData
}

func chatting(conn net.Conn, pUserData *[]User, name string, IPAddr string){
	for{
  	message := GetMessage(conn)
  	fmt.Printf("[%s]:%s\n\n", name, message)

  	if(message == "UsersData"){
  		PrintUsersData(pUserData)
  	}else if(message == "Exit" || message == ""){
  		quitMessage := name + " exits this Room"
  		conn = SendMessage(conn, "Disconnection")
  		fmt.Printf("Exit: %s\n", conn.RemoteAddr())

  		*pUserData = DeleteUser(conn, pUserData, name, IPAddr)
  		SendForAll(conn, quitMessage, pUserData, "ADMIN", "NULL")
  		break
  	}else{
  		SendForAll(conn, message, pUserData, name, IPAddr)
  		// conn = SendMessage(conn, "<GetMessage>")
  	}
	}
}

func main(){
	listen,_ := net.Listen("tcp", "IPAdress:Port")
	var UserData []User = make([]User, 0,100)
	fmt.Println("Running@IPAdress:Port")

	// UserData = append(UserData, User{"192.168.33.10:8000","HostServer"})

	for{
		conn,_ := listen.Accept()
		IPAddr := conn.RemoteAddr().String()
		Name := GetMessage(conn)

		fmt.Printf("New Client: %s\n", IPAddr)
		fmt.Printf("Client Name: %s\n", Name)
		UserData = append(UserData, User{IPAddr, Name, conn})
		conn = SendMessage(conn, "<Enter The Room>")

		SendForAll(conn, " enters this room.", &UserData, Name, "NULL")
		go func(){
			chatting(conn, &UserData, Name, IPAddr)

			conn.Close()
		}()

	}
}
