/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"fmt"
	"github.com/loophole-labs/frisbee"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"time"
)

const PING = uint32(1)
const PONG = uint32(2)

func handlePong(incomingMessage frisbee.Message, incomingContent []byte) (outgoingMessage *frisbee.Message, outgoingContent []byte, action frisbee.Action) {
	if incomingMessage.ContentLength > 0 {
		log.Printf("Client Received Message: %s", string(incomingContent))
	}
	return
}

func main() {
	router := make(frisbee.ClientRouter)
	router[PONG] = handlePong
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)

	c := frisbee.NewClient("127.0.0.1:8192", router)
	err := c.Connect()
	if err != nil {
		panic(err)
	}

	go func() {
		i := 0
		for {
			message := []byte(fmt.Sprintf("ECHO MESSAGE: %d", i))
			err := c.Write(&frisbee.Message{
				To:            0,
				From:          0,
				Id:            uint32(i),
				Operation:     PING,
				ContentLength: uint64(len(message)),
			}, &message)
			if err != nil {
				panic(err)
			}
			i++
			time.Sleep(time.Second)
		}
	}()

	<-exit
	err = c.Close()
	if err != nil {
		panic(err)
	}
}
