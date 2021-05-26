package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const (
	Host      = ""
	Port      = "22"
	User      = "ec2-user"
	Pass      = ""
	PathToKey = ""
	Key       = `-----BEGIN RSA PRIVATE KEY-----\n MIIEogIBAAKCAQ...\n -----END RSA PRIVATE KEY-----\n`
)

func main() {
	one()
}

func one() {
	//conn, c := getPasswordBasedClient()
	conn, c := getKeyFileBasedClient()
	_ = conn

	p := "/tmp"
	//mkdir(c, p)
	write(c, p, "It works!!!\n")
	lstat(c, p)
}

func write(c *sftp.Client, path string, msg string) {
	data := strings.NewReader(msg)

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

func getPasswordBasedClient() (*ssh.Client, *sftp.Client) {
	addr := fmt.Sprintf("%s:%s", Host, Port)
	auths := []ssh.AuthMethod{
		ssh.Password(Pass),
	}
	config := ssh.ClientConfig{
		User:            User,
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

func getKeyFileBasedClient() (*ssh.Client, *sftp.Client) {
	key, err := ioutil.ReadFile(PathToKey)
	if err != nil {
		panic(fmt.Errorf("failed to read key file, error: %w", err))
	}
	// or
	//key = []byte(Key)

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(fmt.Errorf("failed to parse key file, error: %w", err))
	}

	config := &ssh.ClientConfig{
		User: User,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 180 * time.Second,
	}
	addr := fmt.Sprintf("%s:%s", Host, Port)
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic(fmt.Errorf("failed to dial ssh host, error: %w", err))
	}

	c, err := sftp.NewClient(conn, sftp.MaxConcurrentRequestsPerFile(1), sftp.UseFstat(false))
	if err != nil {
		panic(fmt.Errorf("failed to create new client, error: %w", err))
	}

	return conn, c
}
