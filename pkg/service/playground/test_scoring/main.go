package main

import (
	_ "crypto/tls"
	_ "github.com/miekg/dns"
	"net"
	"net/http"
	"os"

	"fmt"
	"github.com/bogdanovich/dns_resolver"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/go-ldap/ldap/v3"
	"github.com/hirochachacha/go-smb2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"time"
)

func main() {

	test_database()
}

// required: None

// required:
//	1) USERNAME
//  2) Password
//  3) DomainName
// Optional Set:
// {
//		independent 1: base DN
// 		independent 2: port
//		independent 3: SSL/nonSSL
// }

func ldap_test() {
	//
	//tlsConfig := &tls.Config{InsecureSkipVerify: true}
	//l, err := ldap.DialTLS("tcp", "172.17.124.119:636", tlsConfig) //DOESN"T WORK YET

	l, err := ldap.Dial("tcp", "172.17.124.119:389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind("Administrator@testdomain.aibek", "Change.me!")
	if err != nil {
		log.Println(err)
	}

	searchRequest := ldap.NewSearchRequest(
		"dc=testdomain,dc=aibek", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))", // The filter to apply
		[]string{"dn", "cn"},                    // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}

// required:
//	1) Entry To resolve
// Optional Set:
// {
//		independent 1: Entry Type (A, NX, AAAA, etc)
// 		independent 2: Expected Output
// 		independent 3: port
// 		Independent 4: TCP/UDP
// }

func dns_test() {
	resolver := dns_resolver.New([]string{"172.17.124.119"})
	resolver.RetryTimes = 1

	ip, err := resolver.LookupHost("WIN-UT74J0G9VRD.testdomain.aibek")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(ip)
}

// required:
//	1) USERNAME
//  2) Password
// Optional Set:
// {
// 		independent 1: text to upload as a file. (Text, and Filename)
// 		independent 2: file to read
//			optional: Expected Output
//		independent 3: Port
//		independent 4: Domain (Use Domain to lookup, or IP)?
//}

func test_smb() {
	conn, err := net.Dial("tcp", "172.17.124.119:445")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     "Administrator",
			Password: "Change.me!",
			Domain:   "",
		},
	}

	c, err := d.Dial(conn)
	if err != nil {
		panic(err)
	}
	defer c.Logoff()

	fs, err := c.Mount(`\\172.17.124.119\SMBShare`)
	if err != nil {
		panic(err)
	}
	defer fs.Umount()

	f, err := fs.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	//defer fs.Remove("hello.txt")
	defer f.Close()

	_, err = f.Write([]byte("Hello world!"))
	if err != nil {
		panic(err)
	}

	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}

// required:
//  1) URL
// Optional Set:
// {
//		independent 1: Port
//		independent 2: Contains Expected Output
//}

func http_test() {

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
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
func test_database() {
	db, err := gorm.Open("mysql", "root:changeme@tcp(172.17.120.203:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

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
