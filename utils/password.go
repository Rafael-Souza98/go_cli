package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)


func GeneratePassword(size int) string{
	rand.Seed(time.Now().UnixNano())
	if size < 8 {
		return "The minimum size must be 8"
	}
	b:= make([]byte, size)
	if _, err:= rand.Read(b); err!= nil{
		fmt.Println("Error to generate password")
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}




// b:= make([]byte, size)
	// rand.Shuffle(size, func(i, j int) {
	// 	if size < 8 {
	// 		fmt.Println("The minimium size must be 8")
	// 	}
	// 	b[i], b[j] = b[j], b[i]
	// })
	// for i := range b {
	// 	fmt.Printf("%c", b[i])
	// }