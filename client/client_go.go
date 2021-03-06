// +build !js

package client

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/the729/go-libra/config"
	"github.com/the729/go-libra/generated/pbac"
	"github.com/the729/go-libra/types/validator"
)

func (c *Client) connect(server string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("grpc dial error: %v", err)
	}
	c.closeFunc = func() { conn.Close() }
	c.ac = pbac.NewAdmissionControlClient(conn)
	return nil
}

func (c *Client) loadTrustedPeers(file string) error {
	peerconf, err := config.LoadTrustedPeersFromFile(file)
	if err != nil {
		return fmt.Errorf("load conf err: %v", err)
	}
	verifier, err := validator.NewConsensusVerifier(peerconf)
	if err != nil {
		return fmt.Errorf("new verifier err: %v", err)
	}
	c.verifier = verifier
	return nil
}
