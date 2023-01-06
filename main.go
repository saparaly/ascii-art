package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	standard   = "b7e06e7f6a2d24d8da5d57d3cba6a2c7"
	shadow     = "0ca33a970e2a1c5b53ecbcad43d60b40"
	thinkertoy = "f7d527c38c0b2ea6df5c12dafb285fd1"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("ERROR: wrong input")
		return
	}
	// arg := os.Args[2]

	// 	if len(os.Args) == 4 {
	// 		if strings.HasPrefix(os.Args[1], "--output=") && strings.HasSuffix(os.Args[1], ".txt") {
	// 			os.Args[1] = os.Args[1][strings.Index(os.Args[1], "=")+1:]

	// fmt.Println(os.Args[1])
	// }else{
	// 	fmt.Println("error")
	// }
	if strings.HasPrefix(os.Args[1], "--output=") && strings.HasSuffix(os.Args[1], "txt") {
		os.Args[1] = os.Args[1][strings.Index(os.Args[1], "=")+1:]
		if os.Args[1][strings.Index(os.Args[1], "=")+1:] == ".txt" {
			fmt.Println("Incorrect input ->OPTION")
			return
		}
	} else {
		fmt.Println("Incorrect input ->OPTION")
		return
	}
	if !(os.Args[3] == "standard" || os.Args[3] == "shadow" || os.Args[3] == "thinkertoy") {
		fmt.Println("Incorrect input -> [BANNER]")
		return
	}

	txt, err := os.ReadFile(os.Args[3] + ".txt")
	if err != nil {
		return
	}
	s := os.Args[2]

	for i := 0; i < len(s); i++ {
		if !(s[i] >= 0 && s[i] <= 127) {
			fmt.Println("ERROR: wrong input")
			return
		}
	}

	checksum := MD5(string(txt))

	if os.Args[3] == "standard" {
		if checksum != standard {
			fmt.Println("standard.txt was changed")
			return
		}
	} else if os.Args[3] == "shadow" {
		if checksum != shadow {
			fmt.Println("shadow.txt was changed")
			return
		}
	} else if os.Args[3] == "thinkertoy" {
		if checksum != thinkertoy {
			fmt.Println("thinkertoy.txt was changed")
			return
		}
	} else {
		fmt.Println("Incorrect input -> [BANNER]")
		return
	}

	fixedTXT := strings.ReplaceAll(string(txt), "\r", "")
	banner := strings.Split(fixedTXT, "\n\n")

	s = strings.ReplaceAll(s, "\\n", "\n")

	replacer := strings.NewReplacer("\"", "\"", "\\", `\`, `\!`, "!")

	d := replacer.Replace(s)
	d = strings.ReplaceAll(d, "\\!", "!")

	words := Check(strings.Split(d, "\n"))
	// words = strings.Split(words, "\n")

	test := ""
	for _, word := range words {
		if word == "" {
			// fmt.Println()
			test += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, l := range word {
				// fmt.Print(strings.Split(banner[l-32], "\n")[i])
				test += strings.Split(banner[l-32], "\n")[i]
			}
			// fmt.Println()
			test += "\n"
		}
	}
	os.WriteFile(os.Args[1], []byte(test), 0666)
}

func Check(s []string) []string {
	c := 0
	for _, v := range s {
		if v == "" {
			c++
		}
	}
	if c == len(s) {
		return s[1:]
	}
	return s
}

// MD5 - Превращает содержимое из переменной data в md5-хеш
func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

// FileMD5 создает md5-хеш из содержимого нашего файла.
func FileMD5(path string) string {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
