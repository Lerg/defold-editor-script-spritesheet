package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type SpriteJson struct {
	Path   string `json:"path"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type SpriteSheetJson struct {
	SpritesheetFilename string       `json:"spritesheetFilename"`
	Sprites             []SpriteJson `json:"sprites"`
}

func main() {
	if len(os.Args) == 2 {
		if err := split_spritesheet(".", os.Args[1]); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln("Missing path to the sprite sheet JSON file as an argument.\nMust be called from the root directory of the Defold project.")
	}
}

func split_spritesheet(project_root_path string, spritesheet_json_path string) error {
	content, err := ioutil.ReadFile(spritesheet_json_path)
	if err != nil {
		log.Fatal("Error when opening JSON file: ", err)
	}

	var spritesheet_json SpriteSheetJson
	err = json.Unmarshal(content, &spritesheet_json)
	if err != nil {
		log.Fatal("Error during reading JSON: ", err)
	}

	spritesheet_image_path := path.Join(path.Dir(spritesheet_json_path), spritesheet_json.SpritesheetFilename)

	spritesheet_image, err := readImage(spritesheet_image_path)
	if err != nil {
		return err
	}

	for _, sprite := range spritesheet_json.Sprites {
		path_parts := []string{project_root_path}
		json_path_parts := strings.Split(sprite.Path[1:], "/")
		path_parts = append(path_parts, json_path_parts...)
		sprite_image_path := path.Join(path_parts...)
		sprite_img, err := cropImage(spritesheet_image, image.Rect(sprite.X, sprite.Y, sprite.X+sprite.Width, sprite.Y+sprite.Height))
		if err != nil {
			return err
		}
		os.MkdirAll(path.Dir(sprite_image_path), os.ModePerm)
		err = writeImage(sprite_img, sprite_image_path)
		if err != nil {
			return err
		}
		fmt.Println("Extracted sprite:", sprite_image_path)
	}

	err = os.Remove(spritesheet_image_path)
	if err != nil {
		return err
	}
	fmt.Println("Deleted spritesheet:", spritesheet_image_path)

	err = os.Remove(spritesheet_json_path)
	if err != nil {
		return err
	}
	fmt.Println("Deleted spritesheet JSON:", spritesheet_json_path)
	return nil
}

// readImage reads a image file from disk. We're assuming the file will be png format.
func readImage(name string) (image.Image, error) {
	fd, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	// image.Decode requires that you import the right image package. We've
	// imported "image/png", so Decode will work for png files. If we needed to
	// decode jpeg files then we would need to import "image/jpeg".
	//
	// Ignored return value is image format name.
	img, _, err := image.Decode(fd)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// cropImage takes an image and crops it to the specified rectangle.
func cropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(crop), nil
}

// writeImage writes an Image back to the disk.
func writeImage(img image.Image, name string) error {
	fd, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fd.Close()

	return png.Encode(fd, img)
}
