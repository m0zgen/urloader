package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

// downloadFile - download file from URL and save it to savePath
func downloadFile(url, savePath string) error {
	// send GET request
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}
	defer response.Body.Close()

	// check if response status code is 200 OK
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("download file error: %v", response.Status)
	}

	// create file for writing
	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("create file error: %v", err)
	}
	defer file.Close()

	// copy response body to file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("write file error: %v", err)
	}

	fmt.Printf("Success. File downloaded and saved to: %s\n", savePath)
	return nil
}

// downloadAndCombine - download files from URLs and combine them into one file
func downloadAndCombine(urls []string, savePath string) error {
	// Open file for writing
	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("create file error: %v", err)
	}
	defer file.Close()

	// Download files from passed URLs
	for _, url := range urls {
		fmt.Printf("Downloading file from: %s\n", url)
		response, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("download file error %s: %v", url, err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("can not download file from %s: status %v", url, response.Status)
		}

		// Копируем содержимое в итоговый файл
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return fmt.Errorf("error writing cntent from %s: %v", url, err)
		}
		fmt.Printf("Content from %s added in %s\n", url, savePath)
	}

	fmt.Printf("Success. All data merged and saved to %s\n", savePath)
	return nil
}

func main() {
	url := flag.String("url", "", "URL-address for downloading file")
	savePath := flag.String("save_path", "", "Path to save downloaded file")
	combineFlag := flag.Bool("combine", false, "Combine multiple files into one")
	flag.Parse()

	if *combineFlag {

		urls := flag.Args() // Get all URLs from command line

		if err := downloadAndCombine(urls, *savePath); err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			os.Exit(1)
		}

	} else {

		if *url == "" || *savePath == "" {
			fmt.Println("Usage: go run download_script.go -url <url> -save_path <save_path>")
			os.Exit(1)
		}

		if err := downloadFile(*url, *savePath); err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			os.Exit(1)
		}
	}
}
