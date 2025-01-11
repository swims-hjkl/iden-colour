package main


import "log"
import "os"
import "fmt"

import "github.com/swims-hjkl/iden-colour/png"


func processIHDR(IHDRChunk []byte) {
	fmt.Printf(string(IHDRChunk))
}

func main()  {
	fileName := "test.png"
	pngContent, err := os.ReadFile(fileName);
	if err != nil {
		log.Fatal("file reading failed")
		return
	}

	png, err := png.NewPNG(nil, fileName)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(png)
	processIHDR(pngContent)

}
