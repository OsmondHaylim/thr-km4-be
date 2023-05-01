package main

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	result := dataMap
	for keys, _ := range result{
		if keys == key{
			delete(result, keys)
		}
	}
	return result
}

func main() {}
