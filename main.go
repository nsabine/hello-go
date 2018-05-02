// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func listDir() {
	files, err := ioutil.ReadDir("/etc/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func fileWrite(logURL string) {

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	// d1 := []byte("hello\ngo\n")
	// err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	// check(err)

	// Create /etc/hello-data
	err := os.Mkdir("/etc/hello-data", 0666)
	f, err := os.Create("/etc/hello-data/" + logURL)
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	defer f.Close()

	// You can `Write` byte slices as you'd expect.
	// d2 := []byte{115, 111, 109, 101, 10}
	// n2, err := f.Write(d2)
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n2)

	// A `WriteString` is also available.
	//n3, err := f.WriteString("writes\n")
	// fmt.Printf("wrote %d bytes\n", n3)

	// Issue a `Sync` to flush writes to stable storage.
	//f.Sync()

	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.
	//w := bufio.NewWriter(f)
	//n4, err := w.WriteString(logURL)
	fmt.Printf("created file called %s\n", logURL)

	// Use `Flush` to ensure all buffered operations have
	// been applied to the underlying writer.
	//w.Flush()
	//listDir()

}

func handler(w http.ResponseWriter, r *http.Request) {
	var pathToLog = r.URL.Path[1:]
	if len(pathToLog) == 0 {
		pathToLog = "rootPath"
	}
	fmt.Fprintf(w, "Hi there, I love persistent apps on GKE.\n Creating file /etc/hello-data/%s!\n",
		pathToLog)
	fileWrite(pathToLog)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
