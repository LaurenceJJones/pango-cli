//go:build !linux && !darwin
// +build !linux,!darwin

package ssh

import (
	"errors"
	"runtime"
)

// RunOpts is shared by both the exec and native SSH runners.
// PrivateKeyPEM and Certificate are set just-in-time (JIT) before connect; no file paths.
// Port is optional: 0 means use default (22 or whatever is in Hostname); >0 overrides.
type RunOpts struct {
	User          string
	Hostname      string
	Port          int    // optional; 0 = default
	PrivateKeyPEM string // in-memory private key (PEM, OpenSSH format)
	Certificate   string // in-memory certificate from sign-key API
	PassThrough   []string
}

// RunExec runs an interactive SSH session by executing the system ssh binary
// (with a PTY when stdin is a terminal on Unix). Requires ssh to be installed.
// opts.PrivateKeyPEM and opts.Certificate must be set (JIT key + signed cert).
func RunExec(opts RunOpts) (int, error) {
	return 1, errors.New("exec SSH runner is not supported on this platform: " + runtime.GOOS)
}
