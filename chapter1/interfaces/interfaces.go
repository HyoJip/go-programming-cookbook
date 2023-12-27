package interfaces

import (
	"fmt"
	"io"
	"os"
)

func Copy(in io.ReadSeeker, out io.Writer) error {

	w := io.MultiWriter(out, os.Stdout)

	// 복사 - 1
	if _, err := io.Copy(w, in); err != nil { // Stdout 출력 - 1
		return err
	}

	in.Seek(0, io.SeekStart)

	// 복사 - 2
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil { // Stdout 출력 - 2
		return err
	}

	fmt.Println()
	return nil
}
