package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/google/goexpect"
	"github.com/google/goterm/term"
)

const (
	timeout = 10 * time.Second
	oPass   = "OM"
	xPass   = "XM"
	jPass   = "JM"
)

var (
	username = flag.String("username", "", "username to use")
	password = flag.String("password", "", "alternate password to use")
	nodeAddr = flag.String("node-addr", "", "node address of ssh server")
	ucpeAddr = flag.String("ucpe-addr", "", "ucpe address of ssh server")
)

func exec() ([]expect.BatchRes, error) {
	conn, _, err := expect.Spawn(fmt.Sprintf("ssh %s@xx.jump.xx.net  -p 2222", *username), -1)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer conn.Close()

	ress, err := conn.ExpectBatch([]expect.Batcher {
		&expect.BExp{R: fmt.Sprintf("%s@xx.jump.xx.net's password:", *username)},
		&expect.BSnd{S: *password + "\r"},
		&expect.BExp{R: "Opt>"},

		&expect.BSnd{S: *nodeAddr + "\r"},
		&expect.BExp{R: "seven@"},

		&expect.BSnd{S: "su -\r"},
		&expect.BExp{R: "Password"},
		&expect.BSnd{S: xPass + "\r"},
		&expect.BExp{R: "root@"},

		&expect.BSnd{S: fmt.Sprintf("ssh -i /etc/openvpn/server/box seven@%s  -p 7722\r", *ucpeAddr)},
		&expect.BExp{R: "Enter passphrase for key '/etc/openvpn/server/box': "},
		&expect.BSnd{S: oPass + "\r"},
		&expect.BExp{R: "seven@7cloudos"},

		&expect.BSnd{S: "su -\r"},
		&expect.BExp{R: "Password"},
		&expect.BSnd{S: xPass + "\r"},
		&expect.BExp{R: "root@sas"},
		&expect.BSnd{S: "/usr/local/svx/xc_info|head -n 1" + "\r"},
		&expect.BExp{R: "root@sas"},
		
	}, timeout)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return ress, nil
}

func main() {
	flag.Parse()
	fmt.Println(term.Bluef("SSH jumpserver"))

	fmt.Println(*username, *password, *nodeAddr, *ucpeAddr)

	ress, err := exec()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(&ress)
	
	fmt.Println(term.Greenf("ok"))
}
