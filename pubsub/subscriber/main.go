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
	"github.com/loophole-labs/frisbee"
	"github.com/rs/zerolog/log"
	"hash/crc32"
	"os"
	"os/signal"
)

const PUB = uint32(1)
const SUB = uint32(2)

var topic = []byte("TOPIC 1")
var topicHash = crc32.ChecksumIEEE(topic)

// Handle the PUB message type
func handlePub(incomingMessage frisbee.Message, incomingContent []byte) (outgoingMessage *frisbee.Message, outgoingContent []byte, action frisbee.Action) {
	if incomingMessage.From == topicHash {
		log.Printf("Client Received Message on Topic %s: %s", string(topic), string(incomingContent))
	}
	return
}

func main() {
	router := make(frisbee.ClientRouter)
	router[PUB] = handlePub
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)

	c := frisbee.NewClient("127.0.0.1:8192", router)
	err := c.Connect()
	if err != nil {
		panic(err)
	}

	i := 0

	// First subscribe to the topic
	err = c.Write(&frisbee.Message{
		From:          0,
		To:            0,
		Id:            uint32(i),
		Operation:     SUB,
		ContentLength: uint64(len(topic)),
	}, &topic)
	if err != nil {
		panic(err)
	}

	// Now the handle pub function will be called
	// automatically whenever a message that matches the topic arrives

	<-exit
	err = c.Close()
	if err != nil {
		panic(err)
	}
}
