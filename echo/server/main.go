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
	"os"
	"os/signal"
)

const PING = uint32(1)
const PONG = uint32(2)

func handlePing(_ *frisbee.Conn, incomingMessage frisbee.Message, incomingContent []byte) (outgoingMessage *frisbee.Message, outgoingContent []byte, action frisbee.Action) {
	if incomingMessage.ContentLength > 0 {
		log.Printf("Server Received Message: %s", incomingContent)
		outgoingMessage = &frisbee.Message{
			From:          incomingMessage.From,
			To:            incomingMessage.To,
			Id:            incomingMessage.Id,
			Operation:     PONG,
			ContentLength: incomingMessage.ContentLength,
		}
		outgoingContent = incomingContent
	}

	return
}

func main() {
	router := make(frisbee.ServerRouter)
	router[PING] = handlePing
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)

	s := frisbee.NewServer(":8192", router)
	err := s.Start()
	if err != nil {
		panic(err)
	}

	<-exit
	err = s.Shutdown()
	if err != nil {
		panic(err)
	}
}
