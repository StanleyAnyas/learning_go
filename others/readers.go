package others

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	if err != nil {
		return 0, err
	}

	for i := 0; i < n; i++ {
		if (p[i] >= 'A' && p[i] <= 'M') || (p[i] >= 'a' && p[i] <= 'm') {
			p[i] += 13
		} else if (p[i] >= 'N' && p[i] <= 'Z') || (p[i] >= 'n' && p[i] <= 'z') {
			p[i] -= 13
		}
	}
	return n, nil
}

type MyImage struct {
	Width, Height int
}

func (img MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img MyImage) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}
func ShowImage(m image.Image) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	io.WriteString(w, "IMAGE:")
	b64 := base64.NewEncoder(base64.StdEncoding, w)
	err := (&png.Encoder{CompressionLevel: png.BestCompression}).Encode(b64, m)
	if err != nil {
		panic(err)
	}
	b64.Close()
	io.WriteString(w, "\n")
}

func reader() {
	r := strings.NewReader("Hello, Readerrrrrrrr!")
	b := make([]byte, 89)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rRot := rot13Reader{s}
	io.Copy(os.Stdout, &rRot)
}

func MainFuncForReaders() {
	reader()
	// m := MyImage{256, 256}
	// ShowImage(m)
}
