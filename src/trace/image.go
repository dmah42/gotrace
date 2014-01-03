package trace

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

type Image [][]Color

const maxColor = 65535

func colorComponentToBytes(c uint16) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, c)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func makeImage(w, h uint32) Image {
	image := make([][]Color, h)
	for i := range image {
		image[i] = make([]Color, w)
		for j := range image[i] {
			image[i][j] = Color{0.0, 0.0, 0.0}
		}
	}
	return image
}

// TODO: gamma correct
func normalizeImage(image Image) {
	// find max
	max := 0.0
	for i := range image {
		for j := range image[i] {
			if image[i][j].R > max {
				max = image[i][j].R
			}
			if image[i][j].G > max {
				max = image[i][j].G
			}
			if image[i][j].B > max {
				max = image[i][j].B
			}
		}
	}

	// normalize
	for i := range image {
		for j := range image[i] {
			image[i][j].Scale(1.0 / max)
		}
	}
}

func WriteImageToPPM(image Image, name string) {
	f, err := os.Create(name + ".ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Printf("Writing image to %s.ppm\n", name)
	normalizeImage(image)

	_, err = f.WriteString(fmt.Sprintf("P6 %d %d %d\n", len(image[0]), len(image), maxColor))
	if err != nil {
		log.Fatal(err)
	}
	for y := range image {
		for x := range image[y] {
			n, err := f.Write(colorComponentToBytes(uint16(image[y][x].R * maxColor)))
			if err != nil {
				log.Fatal(err)
			}
			if n != 2 {
				log.Fatal(fmt.Sprintf("r != 2 bytes: %.3f", image[y][x].R))
			}
			n, err = f.Write(colorComponentToBytes(uint16(image[y][x].G * maxColor)))
			if err != nil {
				log.Fatal(err)
			}
			if n != 2 {
				log.Fatal(fmt.Sprintf("g != 2 bytes: %.3f", image[y][x].G))
			}
			n, err = f.Write(colorComponentToBytes(uint16(image[y][x].B * maxColor)))
			if err != nil {
				log.Fatal(err)
			}
			if n != 2 {
				log.Fatal(fmt.Sprintf("b != 2 bytes: %.3f", image[y][x].B))
			}
		}
	}

}


