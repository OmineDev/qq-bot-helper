package cqhttp_helper

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type CQHttpConnection struct {
	cqHttpAddr        string
	cqHttpBearerToken string
	readerNoBlockCBs  []func([]byte)
	sendNoBlock       func([]byte)
	log               func(string)
}

func NewCQHttpConnection(cqHttpAddr, cqHttpBearerToken string, log func(string)) *CQHttpConnection {
	m := &CQHttpConnection{
		cqHttpAddr:        cqHttpAddr,
		cqHttpBearerToken: cqHttpBearerToken,
		readerNoBlockCBs:  []func([]byte){},
		sendNoBlock:       func([]byte) {},
		log:               log,
	}
	return m
}

func (c *CQHttpConnection) RegisterReaderNoBlockCB(cb func([]byte)) {
	c.readerNoBlockCBs = append(c.readerNoBlockCBs, cb)
}

func (c *CQHttpConnection) SendNoBlock(data []byte) {
	c.sendNoBlock(data)
}

func (c *CQHttpConnection) startRoutine(ctx context.Context) {
	sendChan, readChan := connectToWS(ctx, c.cqHttpAddr, c.cqHttpBearerToken, c.log)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-readChan:
				for _, cb := range c.readerNoBlockCBs {
					cb(data)
				}
			}
		}
	}()
	c.sendNoBlock = func(data []byte) {
		select {
		case <-ctx.Done():
			return
		case sendChan <- data:
		default:
			c.log("send chan full, drop data")
		}
	}
}

func connectToWS(ctx context.Context, addr, bearer string, log func(string)) (sendChan chan []byte, readChan chan []byte) {
	sendChan = make(chan []byte, 128)
	readChan = make(chan []byte, 128)
	if addr == "" {
		log("cqhttp addr is empty, void connection")
		go func() {
			for {
				<-sendChan
			}
		}()
		return
	}
	if !strings.HasPrefix(addr, "ws://") && !strings.HasPrefix(addr, "wss://") {
		addr = "ws://" + addr
	}

	baseReconnectTime := time.Second * 3
	maxReConnectTime := time.Second * 60
	thisReconnectTime := baseReconnectTime
	waitNextReconnectTime := func() time.Duration {
		t := thisReconnectTime
		thisReconnectTime *= 2
		if thisReconnectTime > maxReConnectTime {
			thisReconnectTime = maxReConnectTime
		}
		return t
	}
	clearReconnectTime := func() {
		thisReconnectTime = baseReconnectTime
	}

	go func() {
		for {
			if ctx.Err() != nil {
				return
			}
			log(fmt.Sprintf("connecting to %v...", addr))
			header := http.Header{}
			if bearer != "" {
				header.Set("Authorization", "Bearer "+bearer)
			}
			conn, _, err := websocket.DefaultDialer.Dial(addr, header)
			if err != nil {
				sleepTime := waitNextReconnectTime()
				log(fmt.Sprintf("connect to %v failed, retry after %v", addr, sleepTime))
				time.Sleep(sleepTime)
				continue
			}
			log(fmt.Sprintf("connect to %v success", addr))
			clearReconnectTime()
			connectionDead := make(chan struct{})
			go func() {
				defer conn.Close()
				defer close(connectionDead)
				for {
					_, data, err := conn.ReadMessage()
					if err != nil {
						log(fmt.Sprintf("read message failed: %v", err))
						break
					}
					readChan <- data
				}
			}()
			go func() {
				for {
					select {
					case <-ctx.Done():
						return
					case <-connectionDead:
						return
					case data := <-sendChan:
						err := conn.WriteMessage(websocket.TextMessage, data)
						if err != nil {
							log(fmt.Sprintf("write message failed: %v", err))
							return
						}
					}
				}
			}()
			<-connectionDead
			sleepTime := waitNextReconnectTime()
			log(fmt.Sprintf("connection to %v dead, retry after %v", addr, sleepTime))
			time.Sleep(sleepTime)
		}
	}()
	return sendChan, readChan
}
