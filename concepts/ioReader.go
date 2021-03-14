package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)


type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	var n, error = r13.r.Read(b)
	for i , v := range b {
		if v >= 'a' && v <= 'z' {
			b[i] = (b[i] - 13 + 'a') % 26 + 'a'
		} else if v >= 'A' && v <= 'Z' {
			b[i] = (b[i] - 13 + 'A') % 26 + 'A'
		}
	}

	return n, error
}


func root13ReaderImple() {
	fmt.Println("==========rot13Reader implementation======")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}

func main(){
	var reader = strings.NewReader("Hello I am noob and trying to learn the goLang")

	var chunk = make([]byte, 8)

	for {
		n , error := reader.Read(chunk)
		fmt.Printf("chunk: %q\n", chunk[:n])

		if error == io.EOF {
			fmt.Println("file read successful man")
			break
		}
	}

	root13ReaderImple()

}
