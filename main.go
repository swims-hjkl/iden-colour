package main

import "log"
import "os"
import "fmt"

import "github.com/swims-hjkl/iden-colour/png"

func main()  {
	fileName := "test.png"
	_, err := os.ReadFile(fileName);
	if err != nil {
		log.Fatal("file reading failed")
		return
	}

	png, err := png.NewPNG(nil, fileName)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(png)
	var _ = []byte{0X89, 0X50, 0X4E, 0X47, 0X0D, 0X0A, 0X1A, 0x0A};

}
