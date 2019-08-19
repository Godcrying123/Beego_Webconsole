package controllers

// import (
// 	"fmt"
// 	"log"
// )

// func init() {
// 	go handleMessages()
// }

// //广播发送至页面
// func handleMessages() {
// 	for {
// 		msg := <-Hostchan
// 		for client := range Clients {
// 			err := client.WriteJSON(msg.Memory.UsedMemory)
// 			if err != nil {
// 				log.Printf("client.WriteJSON error: %v", err)
// 				client.Close()
// 				delete(Clients, client)
// 			}
// 		}
// 	}
// }
