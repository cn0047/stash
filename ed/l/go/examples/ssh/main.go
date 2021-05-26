package main

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

const (
	User      = "ec2-user"
	Host      = ""
	Port      = 22
	PathToKey = ""
)

func main() {
	s, err := getSession()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	err = printUname(s)
	if err != nil {
		panic(err)
	}
}

func printUname(s *ssh.Session) error {
	b, err := s.CombinedOutput("uname -a")
	if err != nil {
		return fmt.Errorf("failed to run command, error: %w", err)
	}

	fmt.Printf("uname: %s", b)

	return nil
}

func getSession() (*ssh.Session, error) {
	key, err := ioutil.ReadFile(PathToKey)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file, error: %w", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("failed to parse key file, error: %w", err)
	}

	config := &ssh.ClientConfig{
		User:            User,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%d", Host, Port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial ssh host, error: %w", err)
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create new ssh session, error: %w", err)
	}

	return session, nil
}
