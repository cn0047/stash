package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const (
	HOST = ""
	PORT = ""
	USER = ""
	PASS = ""
)

func main() {
	one()
}

func one() {
	conn, c := getClient()
	_ = conn

	p := "testing/test"
	//mkdir(c, p)
	write(c, p, "It works!\n")
	lstat(c, p)
}

func write(c *sftp.Client, path string, msg string) {
	data := strings.NewReader("it works!\n")

	f, err := c.Create(path + "/" + "test.txt")
	if err != nil {
		fmt.Printf("[write] c.Create error: %+v \n", err)
		return
	}

	_, err = io.Copy(f, data)
	if err != nil {
		fmt.Printf("[write] io.Copy error: %+v \n", err)
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Printf("[write] f.Close error: %+v \n", err)
		return
	}
}

func mkdir(c *sftp.Client, path string) {
	err := c.Mkdir(path)
	if err != nil {
		fmt.Printf("[mkdir] error: %+v \n", err)
	}
}

func lstat(c *sftp.Client, path string) {
	info, err := c.Lstat(path)
	if err != nil {
		fmt.Printf("[lstat] error: %+v \n", err)
		return
	}

	fmt.Printf("🎾 %+v \n", info)
}

func getClient() (*ssh.Client, *sftp.Client) {
	addr := fmt.Sprintf("%s:%s", HOST, PORT)
	auths := []ssh.AuthMethod{
		ssh.Password(PASS),
	}
	config := ssh.ClientConfig{
		User:            USER,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         180 * time.Second,
	}
	conn, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		panic(err)
	}

	c, err := sftp.NewClient(conn, sftp.MaxConcurrentRequestsPerFile(1), sftp.UseFstat(false))
	if err != nil {
		panic(err)
	}

	return conn, c
}
