package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var _url string = `https://www.google.com.br/search?q=`
var _input string = `#tsf > div:nth-child(2) > div.A8SBwf > div.RNNXgb > div > div.a4bIc > input`
var _button string = `#tsf > div:nth-child(2) > div.A8SBwf > div.FPdoLc.tfB0Bf > center > input.gNO89b`

func main() {
	fmt.Println("--------------------------------------------")
	fmt.Println("--       Golang simple search links       --")
	fmt.Println("--------------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a search parameter: ")
	text, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Println(_url + text)

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Navigate(_url+text),
		chromedp.Nodes("#rso a", &nodes),
		chromedp.Stop(),
	)
	if err != nil {
		log.Println(err)
		return
	}

	for _, n := range nodes {
		if strings.HasPrefix(n.AttributeValue("href"), "/") == false {
			fmt.Print(n.AttributeValue("href"), "\t \n")
		}
	}
}
