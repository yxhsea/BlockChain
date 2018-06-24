package main

import (
	"encoding/json"
	"BlockChain/core"
	"io"
	"net/http"
)

var blockChain *core.BlockChain

func run() {
	http.HandleFunc("/block_chain/get", blockChainGetHandle)
	http.HandleFunc("/block_chain/write", blockChainWriteHandle)
	http.ListenAndServe(":8888", nil)
}

func blockChainGetHandle(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockChain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockChainWriteHandle(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockChain.SendData(blockData)
	blockChainGetHandle(w, r)
}

func main() {
	blockChain = core.NewBlockChain()
	run()
}
