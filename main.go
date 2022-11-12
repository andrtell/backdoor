package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/podman/auto-update", podmanAutoUpdate)
	log.Println("Server is starting on port 9999")
	log.Fatal(http.ListenAndServe(":9999", mux))
}

func podmanAutoUpdate(w http.ResponseWriter, r *http.Request) {
	runCommand("podman", "auto-update")
}

func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)

	stdout, _ := cmd.StdoutPipe()

	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	outputScanner := bufio.NewScanner(stdout)
	outputScanner.Split(bufio.ScanLines)
	for outputScanner.Scan() {
		m := outputScanner.Text()
		fmt.Println(m)
	}

	cmd.Wait()
}
