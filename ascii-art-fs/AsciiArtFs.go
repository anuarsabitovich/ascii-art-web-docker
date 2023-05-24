package asciiartfs

// package main
import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	// "errors"
	// "fmt"
	"os"
	"strings"
)

var (
	hash           = "b7e06e7f6a2d24d8da5d57d3cba6a2c7"
	shadowHash     = "0ca33a970e2a1c5b53ecbcad43d60b40"
	thinkertoyHash = "f7d527c38c0b2ea6df5c12dafb285fd1"
	InvalidStyle   = errors.New("Error: Invalid style file! ")
	InvalidInput   = errors.New("Error: Invalid input")
	BadRequestErr  = errors.New("Error: Bad request")
)

func AsciiArtFs(text string, source string) (string, error) {
	standard := ""
	if source == "standard.txt" {
		temp, _ := os.ReadFile("./ascii-art-fs/standard.txt")
		standard = string(temp)
	} else if source == "shadow.txt" {
		temp, _ := os.ReadFile("./ascii-art-fs/shadow.txt")
		standard = string(temp)
	} else if source == "thinkertoy.txt" {
		temp, _ := os.ReadFile("./ascii-art-fs/thinkertoy.txt")
		standard = string(temp)
	} else {
		return "", nil
	}
	standard2 := strings.ReplaceAll(string(standard), "\r", "")
	// if source == "thinkertoy.txt" {}
	arg := strings.ReplaceAll(text, "\r\n", "\n")
	switch source {
	case "standard.txt":
		if GetMD5Hash(string(standard)) != hash {
			return "Invalid style file!", InvalidStyle
		}
	case "shadow.txt":
		if GetMD5Hash(string(standard)) != shadowHash {
			return "Invalid style file!", InvalidStyle
		}
	case "thinkertoy.txt":
		if GetMD5Hash(string(standard)) != thinkertoyHash {
			return "Invalid style file!", InvalidStyle
		}
	}

	splittedStandard := strings.Split(string(standard2), "\n\n")
	newline := 0
	for i, k := range text {
		if i+1 < len(text) {
			if k == '\\' && text[i+1] == 'n' {
				newline++
			}
		}
	}
	// fmt.Println(arg)
	// fmt.Println(newline)
	// argFix = strings.ReplaceAll(argFix, "\n", " .")
	// fmt.Println("Arg: ", argFix)
	argFix2 := strings.Split(arg, "\n")
	// for i := 0; i < len(argFix2); i++ {
	// 	fmt.Print(i, "--:", argFix2[i], " ")
	// }
	var err error
	err = nil
	var final string
	for _, rune := range arg {
		if (rune >= ' ' && rune <= '~') || rune == '\n' || rune == '\r' {
			continue
		} else {
			final = "Incorrect input!"

			return final, InvalidInput
		}
	}
	for _, d := range argFix2 {
		if d == "" {
			if newline != 0 {
				final = final + "\n"
				newline--
				continue
			}
			continue
		}
		for i := 0; i < 8; i++ {
			for _, k := range d {
				temp := strings.Split(splittedStandard[k-32], "\n")
				final = final + string(temp[i])
			}
			if i != 7 {
				final = final + "\n"
			}
		}
		final = final + "\n"
	}
	return final, err
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
