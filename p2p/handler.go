package p2p

import (
	"bufio"
	"log"
	"fmt"
	"os"
	"sync"
	"encoding/json"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	blockchain "github.com/matheusroleal/Votechain/blockchain"
)

var mutex = &sync.Mutex{}

func ReadData(rw *bufio.ReadWriter) {
		for {
			str, err := rw.ReadString('\n')
			if err != nil {
				log.Println("Error reading from buffer")
				return
			}

			if str == "" {
				return
			}

			if str != "\n" {
				chain := make([]blockchain.Block, 0)

				if err := json.Unmarshal([]byte(str), &chain); err != nil {
					log.Println(err)
				}
				mutex.Lock()
				blockchain.ReplaceChain(chain)
				mutex.Unlock()

			}
		}
}

func WriteData(rw *bufio.ReadWriter) {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			mutex.Lock()
			bytes, err := json.Marshal(blockchain.Chain)
			if err != nil {
				log.Println(err)
			}
			mutex.Unlock()

			mutex.Lock()
			rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
			rw.Flush()
			mutex.Unlock()
		}
	}()

	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Println("Error reading from stdin")
			return
		}

		sentData := strings.Replace(sendData, "\n", "", -1)

		tb := blockchain.Transaction{}
		transaction := &tb
		decoder := json.NewDecoder(strings.NewReader(sentData))
		err = decoder.Decode(transaction)

		blockchain.CreateChain(transaction)

		bytes, err := json.Marshal(blockchain.Chain)
		if err != nil {
			log.Println(err)
		}

		spew.Dump(blockchain.Chain)

		mutex.Lock()
		_, err = rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		if err != nil {
			log.Println("Error writing to buffer")
			return
		}
		err = rw.Flush()
		if err != nil {
			log.Println("Error flushing buffer")
			return
		}
		mutex.Unlock()

	}
}
