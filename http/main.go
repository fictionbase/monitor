package main

import (
	"fmt"

	"github.com/fictionbase/fictionbase"
	"github.com/fictionbase/fictionbase/type/fbhttp"
)

func main() {
	fictionbase.SetViperConfig()
	fb := fbhttp.FictionBase{}
	fb.InitKey()
	// check http
	resp, elapsed, err := fbhttp.GetResponseAndTime()
	defer resp.Body.Close()
	if err != nil {
		// @TODO Really???
		fb.Message.Status = 500
		err = fictionbase.SendFictionbaseMessage(fb)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	// @TODO OkStatusCode NgStatusCode get From Config
	if resp.StatusCode > 399 {
		fb.Message.Status = float64(resp.StatusCode)
		err = fictionbase.SendFictionbaseMessage(fb)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	// @TODO ResponseTime get From Config
	if elapsed > 2 {
		fb.Message.ResponseTime = elapsed
		err = fictionbase.SendFictionbaseMessage(fb)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
}
