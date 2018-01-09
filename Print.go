package print

import (
	"bytes"
	"image/png"
	"log"
	"os"
	"path"
	"time"

	"GO/src/github.com/tebeka/selenium"
)

type Test struct {
	wd selenium.WebDriver
}

func (t *Test) Printscreen(namePrint string) error {

	time.Sleep(time.Second * 3)
	ss, err := t.wd.Screenshot()
	if err != nil {
		return err
	}
	r := bytes.NewReader(ss)

	im, err := png.Decode(r)
	if err != nil {
		return err
	}

	if _, err := os.Stat("Screenshots"); err != nil {
		os.Mkdir("Screenshots", 0774)
	}

	fname := path.Join("Screenshots", namePrint)
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_CREATE, 0744)
	if err != nil {
		return err
	}

	png.Encode(f, im)

	log.Printf(" - Screenshot successful")

	return nil

}
