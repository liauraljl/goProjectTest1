package json_demo

import (
	"encoding/json"
	"fmt"
)

func JsonStart() {
	serializeStructTest()
	serializeMapTest()
	DeSerialiseStructTest()
	DeSerialiseMapTest()
}

type Server struct {
	ServerName string
	ServerIp   string
	ServerPort int
}

/*type Server2 struct {
	ServerName2 string json:"name"
	ServerIp2   string 'json:"ip2"'
	ServerPort2 int 'json:"port2"'
}*/

func serializeStructTest() {
	server := new(Server)
	server.ServerName = "server_test_name"
	server.ServerIp = "192.0.0.1"
	server.ServerPort = 8080

	//序列化,系统自带序列化函数
	b, error := json.Marshal(server)
	if error != nil {
		fmt.Println("序列化strcut出现异常：", error.Error())
		return
	}
	fmt.Println("序列化struct结果：", string(b))
}

func serializeMapTest() {
	dog := make(map[string]string)
	dog["name"] = "dog1"
	dog["color"] = "red"
	dog["weight"] = "7kg"

	//序列化,系统自带序列化函数
	b, error := json.Marshal(dog)
	if error != nil {
		fmt.Println("序列化map出现异常：", error.Error())
		return
	}
	fmt.Println("序列化map结果：", string(b))
}

func DeSerialiseStructTest() {
	server := new(Server)
	serverJsonStr := "{\"ServerName\":\"server_test_name\",\"ServerIp\":\"192.0.0.1\",\"ServerPort\":8080}"
	err := json.Unmarshal([]byte(serverJsonStr), &server)
	if err != nil {
		fmt.Println("反序列化struct出现异常：", err.Error())
		return
	}
	fmt.Println("反序列化struct结果：", server)
}

func DeSerialiseMapTest() {
	dog := make(map[string]string)
	dogJsonStr := "{\"color\":\"red\",\"name\":\"dog1\",\"weight\":\"7kg\"}"
	err := json.Unmarshal([]byte(dogJsonStr), &dog)
	if err != nil {
		fmt.Println("反序列化map出现异常：", err.Error())
		return
	}
	fmt.Println("反序列化map结果：", dog)
}
