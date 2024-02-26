package cmd

import (
	"github.com/pedromotita/tcp-server/pkg/client"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use: "client",
	Run: nil,
}

var connectCmd = &cobra.Command{
	Use: "connect",
	RunE: func(cmd *cobra.Command, args []string) error {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			return err
		}

		port, err := cmd.Flags().GetString("port")
		if err != nil {
			return err
		}

		msg, err := cmd.Flags().GetString("message")
		if err != nil {
			return err
		}

		c := client.New()

		conn, err := c.Connect(host, port)
		if err != nil {
			return err
		}

		client.Write(conn, msg)

		return nil
	},
}

func init() {
	connectCmd.Flags().StringP("host", "H", "", "Server IP address")
	connectCmd.Flags().StringP("port", "p", "", "Server port number")
	connectCmd.Flags().StringP("message", "m", "", "Message to send to server")

	connectCmd.MarkFlagRequired("host")
	connectCmd.MarkFlagRequired("flag")
	connectCmd.MarkFlagRequired("message")

	clientCmd.AddCommand(connectCmd)
}
