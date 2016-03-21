package main

import (
	"bufio"
	"fmt"
	"hash/crc64"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	// "github.com/rubyist/go-dnsimple"
	// "net/http"
	// "io"
	// "encoding/json"
)

func getHash(buffer []byte) uint64 {
	table := crc64.MakeTable(crc64.ISO)
	return crc64.Checksum(buffer, table)
}

func makeDir(name string) {
	os.Mkdir(name, 0777)
}

func removeDir(name string) {
	if err := os.RemoveAll(name); err != nil {
		panic(err)
	}
}

func upload_finished(filename string, uploadId string, temp *os.File, etags map[int]string, total_parts int) {

	w := bufio.NewWriter(temp)

	for index := 1; index < total_parts; index++ {
		fmt.Printf("%d: %s\n", index, etags[index])

		full_path := uploadId + "/" + etags[index] + ".tmp"

		buffer, err := ioutil.ReadFile(full_path)
		if err != nil {
			panic(err)
		}

		if _, err := w.Write(buffer); err != nil {
			panic(err)
		}

		if err = w.Flush(); err != nil {
			panic(err)
		}
	}
}

func main() {

	const UPLOAD_ID = "ABCDE1234"
	// const FILENAME = "jaden_playing_basketball.mov"
	const FILENAME = "BMD002416.mov"
	const TEMP = "BMD002416.mov"
	const BUFFER_SIZE = 5 * 1048576

	removeDir(UPLOAD_ID)
	makeDir(UPLOAD_ID)

	source_file, err := os.Open(FILENAME)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := source_file.Close(); err != nil {
			panic(err)
		}
	}()
	r := bufio.NewReader(source_file)

	buf := make([]byte, BUFFER_SIZE)
	etags := make(map[int]string)
	var count int

	for count = 1; ; count++ {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		hash := getHash(buf)
		hash_file_name := strconv.FormatUint(hash, 10)
		etags[count] = hash_file_name
		full_path := UPLOAD_ID + "/" + hash_file_name + ".tmp"

		fo, err := os.Create(full_path)
		if err != nil {
			panic(err)
		}
		w := bufio.NewWriter(fo)

		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}

		if err = w.Flush(); err != nil {
			panic(err)
		}

		fo.Close()
	}

	temp_file, err := os.Create(UPLOAD_ID + "/" + TEMP)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := temp_file.Close(); err != nil {
			panic(err)
		}
	}()

	upload_finished(FILENAME, UPLOAD_ID, temp_file, etags, count)

	fmt.Println("Done")

}
