package gom

import "fmt"

type NewCaptcha struct {
}

const (
	dx = 150
	dy = 50
)

func Draw() {
	captchaImage := gocaptcha.NewCaptchaImage(dx, dy, gocaptcha.RandLightColor())

	err := captchaImage.DrawNoise(gocaptcha.CaptchaComplexLower).
		DrawTextNoise(gocaptcha.CaptchaComplexLower).
		DrawText(gocaptcha.RandText(4)).
		DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A)).
		DrawSineLine().Error

	if err != nil {
		fmt.Println(err)
	}
}
