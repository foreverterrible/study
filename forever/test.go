package forever

import (
	"fmt"
	"io"
	"net/http"
)

func Test() {
	resp, err := http.Get("https://magiceden.io/")
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(body))
		}
	}
}
