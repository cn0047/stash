package main

import (
	"encoding/base64"
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
	Host             = ""
	Port             = "22"
	User             = "ec2-user"
	Pass             = ""
	PathToPrivateKey = ""
	PrivateKeyBase64 = ``
	PrivateKey       = ``
	PublicKey        = ``
)

func main() {
	one()
}

func one() {
	//cfg := getPasswordBasedConfig()
	//cfg := getPrivateKeyFileBasedConfig()
	//cfg := getPrivateKeyBasedConfig()
	//cfg := getRawPrivateKeyBasedConfig(fromBase64(toBase64(PrivateKey)))
	cfg := getRawPrivateKeyBasedConfig(fromBase64(PrivateKeyBase64))
	//cfg := getPublicKeyBasedConfig()
	conn, c := getClient(cfg)
	_ = conn

	p := "/tmp"
	//mkdir(c, p)
	write(c, p, "It works!!!\n")
	lstat(c, p)
}

func toBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func fromBase64(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(fmt.Errorf("failed to decode base64 string, err: %w", err))
	}

	return string(decoded)
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

	fmt.Printf("[lstat] info: %+v \n", info)
}

func getClient(config *ssh.ClientConfig) (*ssh.Client, *sftp.Client) {
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

func getPasswordBasedConfig() *ssh.ClientConfig {
	auth := []ssh.AuthMethod{ssh.Password(Pass)}
	config := &ssh.ClientConfig{
		User:            User,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         180 * time.Second,
	}

	return config
}

func getPrivateKeyFileBasedConfig() *ssh.ClientConfig {
	key, err := ioutil.ReadFile(PathToPrivateKey)
	if err != nil {
		panic(fmt.Errorf("failed to read key file, error: %w", err))
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(fmt.Errorf("failed to parse key file, error: %w", err))
	}

	auth := []ssh.AuthMethod{ssh.PublicKeys(signer)}

	return getConfig(auth)
}

func getPrivateKeyBasedConfig() *ssh.ClientConfig {
	signer, err := ssh.ParsePrivateKey([]byte(PrivateKey))
	if err != nil {
		panic(fmt.Errorf("failed to parse private key, error: %w", err))
	}

	auth := []ssh.AuthMethod{ssh.PublicKeys(signer)}

	return getConfig(auth)
}

func getRawPrivateKeyBasedConfig(key string) *ssh.ClientConfig {
	k, err := ssh.ParseRawPrivateKey([]byte(key))
	if err != nil {
		panic(fmt.Errorf("failed to parse raw private key, error: %w", err))
	}

	signer, err := ssh.NewSignerFromKey(k)
	if err != nil {
		panic(fmt.Errorf("failed to create new signer, error: %w", err))
	}

	auth := []ssh.AuthMethod{ssh.PublicKeys(signer)}

	return getConfig(auth)
}

func getPublicKeyBasedConfig_invalid_1() *ssh.ClientConfig {
	k, err := ssh.ParsePublicKey([]byte(PublicKey))
	if err != nil {
		panic(fmt.Errorf("failed to parse public key, error: %w", err))
	}

	signer, err := ssh.NewSignerFromKey(k)
	if err != nil {
		panic(fmt.Errorf("failed to create new signer, error: %w", err))
	}

	auth := []ssh.AuthMethod{ssh.PublicKeys(signer)}

	return getConfig(auth)
}

func getPublicKeyBasedConfig_invalid_2() *ssh.ClientConfig {
	k, comment, options, rest, err := ssh.ParseAuthorizedKey([]byte(PublicKey))
	if err != nil {
		panic(fmt.Errorf("failed to parse public key, error: %w", err))
	}
	_ = comment
	_ = options
	_ = rest

	signer, err := ssh.NewSignerFromKey(k)
	if err != nil {
		panic(fmt.Errorf("failed to create new signer, error: %w", err))
	}

	auth := []ssh.AuthMethod{ssh.PublicKeys(signer)}

	return getConfig(auth)
}

func getPublicKeyBasedConfig() *ssh.ClientConfig {
	return getPublicKeyBasedConfig_invalid_1()
}

func getConfig(auth []ssh.AuthMethod) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: User,
		Auth: auth,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 180 * time.Second,
	}

	return config
}
