package main

import (
	"fmt"
	"github.com/lionelhiepga/go/singleton"
	"time"
	"github.com/lionelhiepga/go/builder"
)

 
// ví dụ thư mục của bạn là '.../A/singleton/singleton.go' và '.../A/main.go'. 
// Thì bạn phải cd vào thư mục A, chạy lệnh 'go mod init singleton'. 
// Trong file main.go để import package single: import "singleton/singleton" 

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p\n", singleton.GetInstance())
		} ()
	}
	time.Sleep(time.Second*10)
	// s1 := singleton.GetInstance()
	// s2 :=singleton.GetInstance()
}