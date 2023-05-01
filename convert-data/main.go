package main

import "strconv"

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	temp := []string{}
	temp2 := ""
	for i := 0; i < len(data); i++ {
		if data[i] == ','{
			temp = append(temp, temp2)
			temp2 = ""
		}else{
			temp2 += string(data[i])
		}
	}
	temp = append(temp, temp2)
	temp2 = ""
	temp3,_ := strconv.Atoi(temp[1])
	var result = User{
		Name:		temp[0],
		Age:		temp3,
		Address:	temp[2],
	}
	return result
}

func main() {}
