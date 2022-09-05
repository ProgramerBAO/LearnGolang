package main

import (
	"encoding/json"
	"fmt"
)

type User struct {

	// json标签
	UserId    string `json:“UserId”`
	SignH     string `json::SignH`
	PublicKey string `json:"Name"`
}

// Contract 描述构成的基本细节
type Contract struct {
	HashValueOfContract string            `json:"HashValueOfContract"`
	UsersInfo           []User            `json:"UsersInfo"`
	Users               map[string]string `json:"users"` //参与者
}

// 添加合同

// 根据合同哈希查询合同

// 根据的UserId查询

func main() {
	contract1 := Contract{
		HashValueOfContract: "sdsdsdsdsdsdsd",
		UsersInfo: []User{
			{
				UserId:    "1111111",
				SignH:     "1111111",
				PublicKey: "1111111",
			},
			{
				UserId:    "2222222",
				SignH:     "2222222",
				PublicKey: "2222222",
			},
		},
		Users: map[string]string{
			"1": "bob",
			"2": "kenny",
		},
	}

	encodeInfo, err := json.Marshal(&contract1)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	fmt.Println(contract1.Users)

	fmt.Println("encodeInfo:", string(encodeInfo))

	// 反序列化 解码
	var contract2 Contract

	if err = json.Unmarshal([]byte(encodeInfo), &contract2); err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	for i := 0; i < len(contract2.UsersInfo); i++ {
		fmt.Println("Users", i, "UserId", contract2.UsersInfo[i].UserId)
		fmt.Println("Users", i, "SignH", contract2.UsersInfo[i].SignH)
		fmt.Println("Users", i, "PublicKey", contract2.UsersInfo[i].PublicKey)
	}

}
