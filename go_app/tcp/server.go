package main

import(
	"fmt"
	"net"
	"math/rand"
	"time"
)

type User struct{
	IPAdress string
	UserName string
	conn net.Conn
	seed int
}

func genSeed(conn net.Conn)int{
	rand.Seed(time.Now().UnixNano())

	passSeed := rand.Intn(25) + 1
	conn.Write([]byte(fmt.Sprintf("%d", passSeed)))

	return passSeed
}

func SendMessage(conn net.Conn, msg string, seed int)(net.Conn){
	msg = enc(msg, seed)
	conn.Write([]byte(msg))
	return conn
}

func GetMessage(userData *User)(string){
	response := make([]byte, 1024)
  n,_ := userData.conn.Read(response)
  message := string(response[:n])
  message = dec(message, userData.seed)

  return message
}

func PrintMessage(userData *User){
  message := GetMessage(userData)

  fmt.Println(message)
}

func GetPrintMessage(userData *User)(string){
	message := GetMessage(userData)

	fmt.Println(message)
	return message
}

func PrintUsersData(pUserData *[]User){
	for index, UserData := range *pUserData{
		fmt.Printf("[%d] %s : %s\n", index, UserData.IPAdress, UserData.UserName)
	}
}

func SendForAll(userData *User, msg string, allUsersData *[]User){
	msg = fmt.Sprintf("[%s]:%s", userData.UserName, msg)
	for _, UserData := range *allUsersData{
		if(UserData != *userData){
			UserData.conn = SendMessage(UserData.conn, msg, UserData.seed)
		}
	}
}

func DeleteUser(userData *User, allUsersData *[]User)([]User){
	var newUserData []User
	for _, UserData := range *allUsersData{
		if(UserData != *userData){
			newUserData = append(newUserData, UserData)
		}
	}
	return newUserData
}

func chatting(userData *User, allUsersData *[]User){
	for{
  	message := GetMessage(userData)
  	fmt.Printf("[%s]:%s\n\n", userData.UserName, message)

  	if(message == "UsersData"){
  		PrintUsersData(allUsersData)
  	}else if(message == "Exit" || message == ""){
  		quitMessage := userData.UserName + " exits this Room"
  		userData.conn = SendMessage(userData.conn, "Disconnection", userData.seed)
  		fmt.Printf("Exit: %s\n", userData.conn.RemoteAddr())
  		SendForAll(userData, quitMessage, allUsersData)

  		*allUsersData = DeleteUser(userData, allUsersData)
  		break
  	}else{
  		SendForAll(userData, message, allUsersData)
  		// conn = SendMessage(conn, "<GetMessage>")
  	}
	}
}

func main(){
	listen,_ := net.Listen("tcp", "192.168.33.10:8000")
	var UserData []User = make([]User, 0,100)
	fmt.Println("Running@192.168.33.10:8000")

	// UserData = append(UserData, User{"192.168.33.10:8000","HostServer"})

	for{
		conn,_ := listen.Accept()

		go func(){
			userSeed := genSeed(conn)
			IPAddr := conn.RemoteAddr().String()
			Name := GetMessage(&User{"NULL", "NULL", conn, userSeed})

			fmt.Printf("New Client: %s\n", IPAddr)
			fmt.Printf("Client Name: %s\n", Name)
			newUser := User{IPAddr, Name, conn, userSeed}
			UserData = append(UserData, newUser)
			// conn = SendMessage(conn, "<Enter The Room>")

			SendForAll(&newUser, " enters this room.", &UserData)

			chatting(&newUser, &UserData)

			conn.Close()
		}()

	}
}
