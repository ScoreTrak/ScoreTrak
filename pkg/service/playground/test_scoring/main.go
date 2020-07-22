package main

import (
	_ "crypto/tls"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	//"gorm.io/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/miekg/dns"
	"log"
)

func main() {

	//test_database()
}

// required:
//  1) b
//  2) Password
// Optional Set:
// {
//		independent 1: Port
//		independent 2: Expected Mailboxes
//}
func imap_test() {

	//log.Println("Connecting to server...")
	//insecure_config := &tls.Config{InsecureSkipVerify: true}
	//
	//// Connect to server
	//c, err := client.DialTLS("172.17.120.203:993", insecure_config)

	c, err := client.Dial("172.17.120.203:143")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login("testuser", "changeme"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}
	//
	//if err := <-done; err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Select INBOX
	//mbox, err := c.Select("INBOX", false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Flags for INBOX:", mbox.Flags)
	//
	//// Get the last 4 messages
	//from := uint32(1)
	//to := mbox.Messages
	//if mbox.Messages > 3 {
	//	// We're using unsigned integers here, only substract if the result is > 0
	//	from = mbox.Messages - 3
	//}
	//seqset := new(imap.SeqSet)
	//seqset.AddRange(from, to)
	//
	//messages := make(chan *imap.Message, 10)
	//done = make(chan error, 1)
	//go func() {
	//	done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	//}()
	//
	//log.Println("Last 4 messages:")
	//for msg := range messages {
	//	log.Println("* " + msg.Envelope.Subject)
	//}
	//
	//if err := <-done; err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("Done!")
}

// required:
//  1) Username
//  2) Password
// 	3) Database Type [MYSQL, POSTGRE, SQLITE, MSSQL]
// Optional Set:
// {
//		independent 1: Port
//		independent 2: Database to look for
//		independent 2: Schema to look for
//
//}

//notes IMPORT CODE FOR SIDE EFFECT.
// MORE CUSTOMIZATION REQUIRED
//func test_database() {
//	db, err := gorm.Open("mysql", "root:changeme@tcp(172.17.120.203:3306)/testdb")
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//}

//func rdp_test(){ //NOTE TESTED YET. //Make sure not to kick out existing user sessions. // Make sure to increase overall cap of users logged into the server
//	client := grdp.NewClient("172.17.126.181:3389", glog.NONE) //USE XRDP???
//	err := client.Login("testuser", "Change.me!")
//	fmt.Println("LOGGING IN ERRORS/SUCCESS RETURN:")
//	if err != nil {
//		fmt.Println("login failed,", err)
//	} else {
//		fmt.Println("login success")
//	}
//}
