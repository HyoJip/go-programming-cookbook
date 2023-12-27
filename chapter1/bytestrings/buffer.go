package bytestrings

import (
	"bytes"
	"io"
)

func Buffer(rawString string) *bytes.Buffer {

	// string -> Bytes
	//rawBytes := []byte(rawString)

	// 초기화 선언과 값 셋팅 분리 (잘 안씀)
	//var b = new(bytes.Buffer)
	//b.Write(rawBytes)

	// byte -> buffer
	//b = bytes.NewBuffer(rawBytes)

	// string -> buffer
	b := bytes.NewBufferString(rawString)

	return b
}

func toString(r io.Reader) (string, error) {
	//b, err := ioutil.ReadAll(r)
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
