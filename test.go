package main

import (
	"errors"
	"net"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	gossh "golang.org/x/crypto/ssh"
)

const (
	host = "localhost"
	port = "2222"
)

func main() {
	srv, err := wish.NewServer(
		// The address the server will listen to.
		wish.WithAddress(net.JoinHostPort(host, port)),

		// The SSH server need its own keys, this will create a keypair in the
		// given path if it doesn't exist yet.
		// By default, it will create an ED25519 key.
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		// Middlewares do something on a ssh.Session, and then call the next
		// middleware in the stack.
		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(sess ssh.Session) {
					wish.Println(sess, "Hello, world!")
					next(sess)
				}
			},
		),
	)
	srv.ServerConfigCallback = func(ctx ssh.Context) *gossh.ServerConfig {
		return &gossh.ServerConfig{
			Config: gossh.Config{
				KeyExchanges: []string{
					"mlkem768x25519-sha256",
					"sntrup761x25519-sha512",
				},
			},
		}
	}

	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	log.Info("Starting SSH server", "host", host, "port", port)
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		// We ignore ErrServerClosed because it is expected.
		log.Error("Could not start server", "error", err)
	}
}
