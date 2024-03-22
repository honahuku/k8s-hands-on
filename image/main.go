package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// `stress-ng` コマンドを非同期で実行
		go func() {
			cmd := exec.Command("bash", "-c", "stress-ng --cpu 8 --timeout 60s")
			if err := cmd.Start(); err != nil {
				fmt.Println("Error starting stress-ng command:", err)
				return
			}
			if err := cmd.Wait(); err != nil {
				fmt.Println("Error waiting for stress-ng command:", err)
			}
		}()

		// 少し待ってから `ps` コマンドを実行し、`stress-ng` プロセスをカウント
		cmd := exec.Command("bash", "-c", "ps aux | grep [s]tress-ng | wc -l")
		output, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// `stress-ng` プロセスの数をレスポンスとして返す
		processCount := strings.TrimSpace(string(output))
		fmt.Fprintf(w, "stress-ng processes: %s\n", processCount)
	})

	fmt.Println("Server is listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
