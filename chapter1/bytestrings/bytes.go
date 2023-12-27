package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array ❤️"

	b := Buffer(rawString)

	// buffer 객체를 Bytes(), String()을 통해 convert 가능
	fmt.Println(b.String())

	// Reader 인터페이스를 만족하므로(Read(byte[]) 구현-덕타이핑) toString(Reader) 사용 가능
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	// byte로 새로운 byteReader 생성
	reader := bytes.NewReader([]byte(rawString))

	// 버퍼링처리된 읽기 및 토큰화를 허용하는 스캐너 연결
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords) // 공백을 split 기준으로

	for scanner.Scan() { // while (true | false)
		fmt.Print(scanner.Text())
	}

	return nil
}
