package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/disintegration/imaging"
)

var (
	destination = "/tmp/lock-blur.png"
)

func main() {
	cmd := exec.Command("grim", "/tmp/lock.png")
	log.Printf("Running grim and waiting for it to finish...")
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	img, err := os.Open("/tmp/lock.png")
	if err != nil {
		panic(err)
	}
	defer img.Close()

	srcImage, _, err := image.Decode(img)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	blurimg := imaging.Blur(srcImage, 5)
	end := time.Since(start)
	fmt.Printf("Generated in: %.2fs\n", end.Seconds())
	generateImage(destination, blurimg)
	cmd = exec.Command("swaylock", "-i", "/tmp/lock-blur.png")
	log.Printf("Taking screenshot using grim")
	err = cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}
func generateImage(dst string, img image.Image) {
	fq, err := os.Create(destination)
	if err != nil {
		panic(err)
	}
	defer fq.Close()

	if err = png.Encode(fq, img); err != nil {
		log.Fatal(err)
	}

}
