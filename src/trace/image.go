package trace

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
)

type Image [][]Color

const maxColor = math.MaxUint16

// Convert a row of uint16s to a byte slice
func colorToBytes(r [][3]uint16) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, r)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

// Create an image w x h pixels
func makeImage(w, h uint) Image {
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
// Ensure the brightest component of any colour is 1.0
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

func (i Image) write(name string) {
	// TODO: fill this out as per goradiosity.
}

// Write the given image to a PPM format file of the given name
func WriteImageToPPM(image Image, name string) {
	f, err := os.Create(name + ".ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Printf("Normalizing image\n")
	normalizeImage(image)

	log.Printf("Converting image\n")
	outImage := make([][][3]uint16, len(image))
	for i := range outImage {
		outImage[i] = make([][3]uint16, len(image[i]))
		for j := range outImage[i] {
			outImage[i][j][0] = uint16(image[i][j].R * maxColor)
			outImage[i][j][1] = uint16(image[i][j].G * maxColor)
			outImage[i][j][2] = uint16(image[i][j].B * maxColor)
		}
	}

	log.Printf("Writing image to %s.ppm\n", name)
	_, err = f.WriteString(fmt.Sprintf("P6 %d %d %d\n", len(outImage[0]), len(outImage), maxColor))
	if err != nil {
		log.Fatal(err)
	}
	for y := range outImage {
		// Write a row of the image to the file
		_, err := f.Write(colorToBytes(outImage[y]))
		if err != nil {
			log.Fatal(err)
		}
	}
}
