package main

import (
	"github.com/reiver/go-oi"

	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func newsHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	var url string = "http://swapi.savla.su/news/krusty/"

	// try URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("X-Auth-Token", os.Getenv("SWAPI_TOKEN"))

	// get the raw Body
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// read Body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return err
	}

	// parse JSON stream into Go object
	var newsStream News
	if err := json.Unmarshal(data, &newsStream); err != nil {
		log.Print(err)
		return err
	}

	// loop over news items 
	for i, item := range newsStream.News {
		// hardcoded paging limit for VTxxx terminals
		if i > 4 {
			break
		}

		oi.LongWriteString(stdout, fmt.Sprintf("%s\n\r", item.Title))
		oi.LongWriteString(stdout, fmt.Sprintf("[ %s / %s ]\n\r\n\r", item.PubDate, item.Server))
		oi.LongWriteString(stdout, fmt.Sprintf("%s\n\r\n\r", item.Perex))
		oi.LongWriteString(stdout, fmt.Sprintf("--------------------------------------------\n\r\n\r"))
	}
	return nil
}

func helpHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	/*for _, cmd := range cmds {
		oi.LongWriteString(stdout, "\n\r   " + cmd.name + " - " + cmd.description)
	}*/

	oi.LongWriteString(stdout, "\r\n")

	return nil
}

func versionHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	oi.LongWriteString(stdout, "\n\rbbs-go telnet service")
	oi.LongWriteString(stdout, "\n\rversion: " + version + "\n\r\r\n")

	return nil
}

func fiveHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	oi.LongWriteString(stdout, "The number FIVE looks like this: 5\r\n")

	return nil
}

func danceHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	for i:=0; i<20; i++ {
		oi.LongWriteString(stdout, "\r⠋")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠙")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠹")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠸")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠼")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠴")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠦")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠧")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠇")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠏")
		time.Sleep(50*time.Millisecond)
	}
	oi.LongWriteString(stdout, "\r \r\n")

	return nil
}

