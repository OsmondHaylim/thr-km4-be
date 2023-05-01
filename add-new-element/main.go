package main

func AddElement(data []int, newData int, position string) []int {
	switch position{
	case "up":
		result := []int{newData}
		result = append(result, data...)
		return result
	case "down":
		data = append(data, newData)
		return data
	default:
		return nil
	}
}

func main() {}
