package cli

import (
	"strconv"

	"exabits/x/computing/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateDevices() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-devices [address] [machine-id] [machine-token] [coords] [ip]",
		Short: "Create a new devices",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argMachineId := args[1]
			argMachineToken := args[2]
			argCoords := args[3]
			argIp := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDevices(clientCtx.GetFromAddress().String(), argAddress, argMachineId, argMachineToken, argCoords, argIp)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateDevices() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-devices [id] [address] [machine-id] [machine-token] [coords] [ip]",
		Short: "Update a devices",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argAddress := args[1]

			argMachineId := args[2]

			argMachineToken := args[3]

			argCoords := args[4]

			argIp := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDevices(clientCtx.GetFromAddress().String(), id, argAddress, argMachineId, argMachineToken, argCoords, argIp)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteDevices() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-devices [id]",
		Short: "Delete a devices by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteDevices(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
