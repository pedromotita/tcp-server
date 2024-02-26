package cmd

import (
	"github.com/pedromotita/tcp-server/pkg/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: nil,
}

var runCmd = &cobra.Command{
	Use: "run",
	RunE: func(cmd *cobra.Command, args []string) error {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			return err
		}

		port, err := cmd.Flags().GetString("port")
		if err != nil {
			return err
		}

		s := server.New(host, port)
		s.Run()

		return nil
	},
}

func init() {
	runCmd.Flags().StringP("host", "H", "", "Server IP address")
	runCmd.Flags().StringP("port", "p", "", "Server port number")

	runCmd.MarkFlagRequired("host")
	runCmd.MarkFlagRequired("flag")

	serverCmd.AddCommand(runCmd)
}
