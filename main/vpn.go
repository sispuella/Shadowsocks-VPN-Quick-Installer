// vpn installation project vpn installation.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

type Config struct {
	ip string
	pw string
	sp string
	lp string
}

func (nb *Config) GetIP() {
	client := &http.Client{Timeout: 10 * time.Second}
	curl := "http://ip.chinaz.com/"
	request, err := http.NewRequest("GET", curl, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:60.0) Gecko/20100101 Firefox/60.0")
	request.Header.Set("Referer", "http://ip.chinaz.com/siteip")
	if err != nil {
		fmt.Println("Request Fail, Please Restart the Program. If this keeps occuring, please contact \"kliest.yang@smus.ca\"")
	}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("Request Error, please retry. Code:", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		dom, _ := goquery.NewDocumentFromResponse(res)
		tmp := dom.Find("div dl dd").Text()
		tmp1 := strings.Fields(tmp)
		tmp2 := strings.TrimSpace(tmp1[0])
		ip := strings.TrimRightFunc(tmp2, func(r rune) bool {
			if unicode.IsDigit(r) == true {
				return false
			} else {
				return true
			}
		})
		nb.ip = ip
		fmt.Println("The detected IP address is :" + `"` + nb.ip + `"`)
	} else {
		fmt.Println("Status code error, please retry. Code: ", res.StatusCode)
		os.Exit(1)
	}
}

func (nb *Config) GetPW() {
	pw := bufio.NewReader(os.Stdin)
	var input string
	var passedType string
	for {
		fmt.Println("Please define your VPN password:")
		input, _ = pw.ReadString('\n')
		if strings.Contains(input, " ") == true {
			fmt.Println("Your VPN password should not contain any space")
			continue
		} else {
			passedType = strings.TrimSpace(input)
			fmt.Println("Your VPN password is:" + `"` + passedType + `"`)
			break
		}
	}
	nb.pw = passedType
}

func (nb *Config) GetSP() {
	pw := bufio.NewReader(os.Stdin)
	var input string
	var passedType string
	for {
		fmt.Println("Please define your Server Port:")
		input, _ = pw.ReadString('\n')
		if strings.Contains(input, " ") == true || (isDigit(strings.TrimSpace(input)) == false) {
			fmt.Println("Your Server Port # should not contain any space or non-digit")
			continue
		} else {
			passedType = strings.TrimSpace(input)
			break
		}
	}
	nb.sp = passedType
}

func (nb *Config) GetLP() {
	pw := bufio.NewReader(os.Stdin)
	var input string
	var passedType string
	for {
		fmt.Println("Please define your Local Port:")
		input, _ = pw.ReadString('\n')
		if strings.Contains(input, " ") == true || (strings.TrimSpace(input) == nb.sp) || (isDigit(strings.TrimSpace(input)) == false) {
			fmt.Println("Your Local Port # should not contain any space or non-digit, and must be different from Server Port #")
			continue
		} else {
			passedType = strings.TrimSpace(input)
			fmt.Println("-----------------------------------------------------------")
			break
		}
	}
	nb.lp = passedType
}
func (nb *Config) writeConfig() {
	file, err := os.OpenFile("/etc/shadowsocks.json", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0)
	if err != nil {
		log.Fatal(`Failed to create config file. Please check if path /etc/ exists, and is accessible and writable.`)
	}
	file.WriteString(`{`)
	file.WriteString(`"server":"` + nb.ip + `",`)
	file.WriteString(`"server_port":` + nb.sp + `,`)
	file.WriteString(`"local_address":"127.0.0.1",`)
	file.WriteString(`"local_port":` + nb.lp + `,`)
	file.WriteString(`"password":"` + nb.pw + `",`)
	file.WriteString(`"timeout":300,`)
	file.WriteString(`"method":"aes-256-cfb",`)
	file.WriteString(`"fast_open":false`)
	file.WriteString(`}`)
	defer file.Close()
	fmt.Println("The VPN is configure as following:")
	fmt.Println("Server IP is now configured as:" + nb.ip)
	fmt.Println("Server Port is now configured as Port #:" + nb.sp)
	fmt.Println("Local port is now configured as Port #:" + nb.lp)
	fmt.Println("VPN Password is now configured as:" + `"` + nb.pw + `".` + "Encoding: aes-256-cfb")

}

func isDigit(input string) bool {
	tmp := []rune(input)
	var a bool
	for i := range tmp {
		if unicode.IsDigit(tmp[i]) {
			a = true
		} else {
			a = false
			break
		}
	}
	return a
}

func main() {
	f := Config{}
	f.GetIP()
	f.GetPW()
	f.GetSP()
	f.GetLP()
	f.writeConfig()
}
