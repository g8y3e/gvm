package helper

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	GoBaseAddress = "https://golang.org/dl/go"
)

func Download(version, systemName, arch, root string) bool {
	url := GoBaseAddress + version + "." + systemName + "-" + arch + ".zip"
	destFile := filepath.Join(root, "go" + version + ".zip")
	destDir := filepath.Join(root, "go" + version)

	fmt.Println("Downloading Go v" + version + "... Please wait...")
	// create a file to store downloaded data
	out, err := os.Create(destFile)
	if err != nil {
		fmt.Println("Can't create folder:", destDir, "; error:", err)
		return false
	}

	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Error while downloading", url, "; error:", err)
		return false
	}

	if response.StatusCode != 200 {
		fmt.Println("Download failed for url " + url + ". You can check the url manually. Rolling Back.")

		//remove the zip file after closing it.
		defer os.Remove(destDir)
		defer out.Close()
		return false
	}

	_, err = io.Copy(out, response.Body)
	if err != nil {
		fmt.Println("Error while copy file response ", destDir, "-", err)
	}

	fmt.Println("Unzipping files...")
	zipErr := unzip(destFile, destDir)
	if zipErr != nil {
		fmt.Println("Error while unzipping", destFile, "-", zipErr)
		return false
	}
	//remove the zip file after closing it.
	defer os.Remove(destFile)
	defer out.Close()
	return true
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()
	os.MkdirAll(dest, 0755)
	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			//make any needed folders for the file in question
			os.MkdirAll(filepath.Clean(path+"\\.."), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()
			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}
	return nil
}