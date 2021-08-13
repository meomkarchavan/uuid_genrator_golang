package uuid_genrator

import (
	"crypto/rand"
	"fmt"
)

func Uuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return err.Error()
	}
	fmt.Println(b)

	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:3], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid

}
