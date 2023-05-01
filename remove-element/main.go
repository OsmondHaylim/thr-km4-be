package main

func removeLeftRight(arr []any, position string) []any {
	switch position{
	case "left":
		result := arr[1:]
	case "right":
		result := arr[:len(arr)-1]
	default:
		result := arr
	}
	return result
}

func main() {}
