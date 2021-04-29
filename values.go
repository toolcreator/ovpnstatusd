package main

import (
	"bytes"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

type ovpnStatus struct {
	commonName  string
	clientCount uint
}

func catOvpnStatus(args *argBundle) (string, error) {
	sshClientCfg := &ssh.ClientConfig{
		User: *args.user,
		Auth: []ssh.AuthMethod{
			ssh.Password(*args.password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(*args.timeout) * time.Second,
	}

	sshClient, err := ssh.Dial("tcp", *args.destination, sshClientCfg)
	if err != nil {
		return "", err
	}
	defer sshClient.Close()

	sshSession, err := sshClient.NewSession()
	if err != nil {
		return "", err
	}
	defer sshSession.Close()

	var buffer bytes.Buffer
	sshSession.Stdout = &buffer
	if err := sshSession.Run("cat " + *args.remotePath); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func parseOvpnStatusStr(str string) []ovpnStatus {
	lines := strings.Split(str, "\n")

	const (
		SEARCH_START = iota
		READ         = iota
		STOP         = iota
	)

	connections := make(map[string]uint)

	parseState := SEARCH_START
	for _, line := range lines {
		switch parseState {

		case SEARCH_START:
			columns := strings.Split(line, ",")
			if len(columns) > 0 {
				if columns[0] == "Common Name" {
					parseState = READ
				}
			}

		case READ:
			columns := strings.Split(line, ",")
			if len(columns) > 0 {
				if columns[0] == "ROUTING TABLE" {
					parseState = STOP
				} else {
					connections[columns[0]] += 1
				}
			}

		}
		if parseState == STOP {
			break
		}
	}

	var ret []ovpnStatus
	for commonName, clientCount := range connections {
		ret = append(ret, ovpnStatus{
			commonName:  commonName,
			clientCount: clientCount,
		})
	}

	return ret
}

func getValues(args *argBundle) ([]ovpnStatus, error) {
	ovpnStatusStr, err := catOvpnStatus(args)
	if err != nil {
		return nil, err
	}

	return parseOvpnStatusStr(ovpnStatusStr), nil
}
