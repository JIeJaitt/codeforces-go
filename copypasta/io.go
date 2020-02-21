package copypasta

import (
	"bufio"
	. "fmt"
	"io"
)

func simpleIO(_r io.Reader, _w io.Writer) {
	// NOTE: just a bufio.NewReader is enough, there is no difference between this and ioutil.ReadAll
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)

	Fprintln(out, n)
	// NOTE: to print int as char, use Fprintf(out, "%c", 'a'+1)
	// NOTE: to print []byte as string, use Fprintf(out, "%s", data)
}

// 一般来说读 1e5 个 int 需要 50-100ms
func fastIO(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}
	// 注意：若有负数请使用下面这个！
	read = func() (x int) {
		in.Scan()
		data := in.Bytes()
		if data[0] == '-' {
			for _, b := range data[1:] {
				x = x*10 + int(b-'0')
			}
			return -x
		}
		for _, b := range data {
			x = x*10 + int(b-'0')
		}
		return
	}

	// FIXME: why AC with Fscan but WA5 with readS() in CF827A ?
	in.Buffer(nil, 1e9)
	readS := func() string { in.Scan(); return string(in.Bytes()) }

	_ = []interface{}{read, readS}
}

func lineIO(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Buffer(nil, 1e9) // default maxTokenSize is 65536
	out := bufio.NewWriter(_w)
	defer out.Flush()

	for in.Scan() {
		line := in.Text()

		Fprintln(out, line)
	}
}

// 由于瓶颈在 Read() 上，使用 buffer 减少调用次数才是关键，优化其余逻辑已无明显意义（对于 2e6 只能减 60ms）
//scanToken := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
//	n := len(data)
//	// Skip leading spaces.
//	start := 0
//	for ; start < n; start++ {
//		if r := data[start]; r != ' ' && r != '\n' && r != '\r' {
//			break
//		}
//	}
//	// Scan until space, marking end of word.
//	for i := start; i < n; i++ {
//		if r := data[i]; r == ' ' || r == '\n' || r == '\r' {
//			return i + 1, data[start:i], nil
//		}
//	}
//	// If we're at EOF and have a non-empty, non-terminated word. Return it.
//	if atEOF && start < n {
//		return len(data), data[start:], nil
//	}
//	// Request more data.
//	return start, nil, nil
//}
//in.Split(scanToken)
