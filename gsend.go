// Copyright 2013 - by Jim Lawless
// License: MIT / X11
// See: http://www.mailsend-online.com/license2013.php
//
// Bear with me ... I'm a Go noob.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func main() {
	to := flag.String("t", "", "destination Internet mail address")
	from := flag.String("f", os.Getenv("GSEND_USER"), "the sender's GMail address")
	pwd := flag.String("p", os.Getenv("GSEND_PWD"), "the sender's password")
	subject := flag.String("s", "", "subject line of email")
	msg := flag.String("m", "", "a one-line email message")
	flag.Usage = func() {
		fmt.Println("Syntax:\n\tgsend [flags]\nwhere flags are:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("-f and -p may also be set using the environmental variables " +
			"GSEND_USER and GSEND_PWD respectively.")
	}

	fmt.Printf("GSend v 1.01 by Jim Lawless\n")

	flag.Parse()

	for _, s := range [...]*string{to, from, pwd, subject, msg} {
		if *s == "" {
			flag.Usage()
			os.Exit(1)
		}
	}

	body := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", *to, *from, *msg)
	auth := smtp.PlainAuth("", *from, *pwd, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, *from,
		[]string{*to}, []byte(body))
	if err != nil {
		log.Fatal(err)
	}
}
