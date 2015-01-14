// Copyright (c) 2014  by f4lk0r_ 
// Licensed under the GPLv3, see the License.txt file for details

package main

import(
    "log"
    "io/ioutil"
    ts "github.com/janimo/textsecure"
)

func messageHandler(msg *ts.Message) {
    err := ts.SendMessage(msg.Source(), msg.Message())
    if err != nil {
        log.Println(err)
    }

    if msg.Message() != "" {
        //fmt.Printf("\r                                               %s%s : %s%s%s\n>", red, getName(msg.Source()), green, msg.Message(), blue)
    }

    for _, a := range msg.Attachments() {
        handleAttachment(msg.Source(), a)
    }
}

// Recieves incoming attachments
func handleAttachment(src string, b []byte) {
    f, err := ioutil.TempFile(".", "TextSecure_Attachment")
    if err != nil {
        debugLog.Println(err)
        return
    }
    log.Printf("Saving attachment of length %d from %s to %s", len(b), src, f.Name())
    f.Write(b)

}
//
//// prints to the message window screen.
//func msgWinPrint(a ... interface{}) {
//}