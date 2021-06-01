package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/fogleman/gg"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) < 7 {
		fmt.Println(fmt.Sprintf(`Invalid arguments!`))
		return
	}

	name := arguments[0]
	idolImage := arguments[1]
	frameImage := arguments[2]
	maskImage := arguments[3]
	hasGroup, err := strconv.ParseBool(arguments[4])
	groupLogo := arguments[5]
	finalUrl := arguments[6]

	if err != nil {
		log.Fatal(err)
		return
	}

	Draw(name, idolImage, frameImage, maskImage, hasGroup, groupLogo, finalUrl)
}

func Draw(name string, idolImage string, frameImage string, maskImage string, hasGroup bool, groupLogo string, final string) error {

	dc := gg.NewContext(350, 500)

	err := dc.LoadFontFace(`./src/assets/fonts/AlteHaasGroteskBold.ttf`, 30)

	if err != nil {
		log.Fatal(err)
		return err
	}

	idol, err := gg.LoadImage(idolImage)

	if err != nil {
		log.Fatal(err)
		return err
	}

	dc.DrawImage(idol, 47, 54)

	frame, err := gg.LoadImage(frameImage)

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = exec.Command("convert", maskImage, "-fill", "rgb(185,185,185)", "-colorize", "100", "colorized-mask.png").Run()

	if err != nil {
		log.Fatal(err)
		return err
	}

	mask, err := gg.LoadImage("./colorized-mask.png")

	if err != nil {
		log.Fatal(err)
		return err
	}

	dc.DrawImage(frame, 0, 0)

	dc.DrawImage(mask, 0, -1)

	// SetRGB takes a float64 for each argument (between 0 and 1),
	// so we need to divide by 255 to get that number.
	dc.SetRGB(185/255, 185/255, 185/255)

	textX := float64(50)
	nameY := float64(436)

	dc.DrawString(name, textX, nameY)

	if hasGroup {
		group, err := gg.LoadImage(groupLogo)

		if err != nil {
			log.Fatal(err)
			return err
		}

		dc.DrawImage(group, 0, 0)

	}

	dc.SavePNG(final)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf(final)

	return nil

}
