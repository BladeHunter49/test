package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"

	// "os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
)

type Block struct {
	Index     int
	TimeStamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var BlockChain []Block

//计算块的哈希值
func calcHash(block Block) string {
	record := string(block.Index) + block.TimeStamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//生成新块
func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.TimeStamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calcHash(newBlock)

	return newBlock, nil
}

//校验块
func isBlockValid(newBlock, oldBlock Block) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}
	if newBlock.PrevHash != oldBlock.Hash {
		return false
	}
	if calcHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

//两个节点都生成块并添加到各自的链上,始终选择最长的链!
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(BlockChain) {
		BlockChain = newBlocks
	}
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(BlockChain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //code 500
		return
	}
	io.WriteString(w, string(bytes))
}

type Message struct {
	BPM int
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}

	// m.BPM = 23
	newBlock, err := generateBlock(BlockChain[len(BlockChain)-1], m.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if isBlockValid(newBlock, BlockChain[len(BlockChain)-1]) {
		newBlockChain := append(BlockChain, newBlock)
		replaceChain(newBlockChain)
		spew.Dump(BlockChain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func run() error {
	mux := makeMuxRouter()
	httpAddr := "8080"
	log.Println("Listening on 8080")
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, //1左移20位
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	go func() {
		t := time.Now()
		firstBlock := Block{0, t.String(), 42, "", ""}
		spew.Dump(firstBlock) //方便查看struct
		BlockChain = append(BlockChain, firstBlock)
	}()
	log.Fatal(run())
}
