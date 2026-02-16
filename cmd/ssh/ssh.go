package ssh

import (
	"errors"
	"os"

	"github.com/fosrl/cli/internal/logger"
	"github.com/spf13/cobra"
)

var errHostnameRequired = errors.New("--hostname is required")

func SSHCmd() *cobra.Command {
	opts := struct {
		User     string
		Hostname string
		Identity string
	}{}

	cmd := &cobra.Command{
		Use:   "ssh",
		Short: "Run an interactive SSH session",
		Long:  `Run an SSH client in the terminal. Uses the system ssh binary with a PTY for interactive sessions.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.Hostname == "" {
				return errHostnameRequired
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			exitCode, err := Run(RunOpts{
				User:        opts.User,
				Hostname:    opts.Hostname,
				Identity:    opts.Identity,
				PassThrough: args,
			})
			if err != nil {
				logger.Error("%v", err)
				os.Exit(1)
			}
			os.Exit(exitCode)
		},
	}

	cmd.Flags().StringVarP(&opts.User, "user", "u", "", "SSH login user (maps to ssh -l)")
	cmd.Flags().StringVar(&opts.Hostname, "hostname", "", "Target host (required)")
	cmd.Flags().StringVarP(&opts.Identity, "identity", "i", "", "Path to private key file")

	// Allow arbitrary args after flags (e.g. after --) to pass through to ssh
	cmd.Args = cobra.ArbitraryArgs

	return cmd
}
