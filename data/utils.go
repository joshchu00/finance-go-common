package data

import "fmt"

func GetPath(dataDirectory string, dateString string) string {
	return fmt.Sprintf("%s/%s.json", dataDirectory, dateString)
}
