package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func SearchString() {
	s := "this is a test"

	fmt.Println(strings.Contains(s, "this"))   // true
	fmt.Println(strings.ContainsAny(s, "abc")) // true
	fmt.Println(strings.HasPrefix(s, "this"))  // true
	fmt.Println(strings.HasSuffix(s, "test"))  // true
}

func ModifyString() {
	s := "simple string"

	fmt.Println(strings.Split(s, " "))
	fmt.Println(strings.Title(s))
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

func StringReader() {
	s := "simple string\n"
	r := strings.NewReader(s)

	io.Copy(os.Stdout, r)
}
