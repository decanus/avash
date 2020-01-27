/*
Copyright © 2019 AVA Labs <collin@avalabs.org>
*/

package cmd

import (
	"fmt"
	"strconv"
	"time"

	pmgr "github.com/ava-labs/avash/processmgr"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// ProcmanagerCmd represents the procmanager command
var ProcmanagerCmd = &cobra.Command{
	Use:   "procmanager [operation]",
	Short: "Access the process manager for the avash client.",
	Long: `Access the process manager for the avash client. Using this 
	command you can list, stop, and start processes registered with the 
	process manager.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("procmanager requires an operation. Available: list, start, stop, stopall, startall, kill, killall, remove, metadata")
	},
}

// PMListCmd represents the list operation on the procmanager command
var PMListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the processes currently running.",
	Long:  `Lists the processes currently running in tabular format.`,
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(AvaShell.rl.Stdout())
		table = pmgr.ProcManager.ProcessTable(table)
		table.Render()
	},
}

// PMMetadataCmd represents the list operation on the procmanager command
var PMMetadataCmd = &cobra.Command{
	Use:   "metadata [node name]",
	Short: "Prints the metadata associated with the node name.",
	Long:  `Prints the metadata associated with the node name.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 && args[0] != "" {
			name := args[0]
			metadata, err := pmgr.ProcManager.Metadata(name)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(metadata)
		} else {
			cmd.Help()
		}
	},
}

// PMStartCmd represents the start operation on the procmanager command
var PMStartCmd = &cobra.Command{
	Use:   "start [node name] [optional: delay in secs]",
	Short: "Starts the process named if not currently running.",
	Long:  `Starts the process named if not currently running.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 && args[0] != "" {
			name := args[0]
			delay := time.Duration(0)
			if len(args) >= 2 {
				if v, e := strconv.ParseInt(args[1], 10, 64); e == nil && v > 0 {
					delay = time.Duration(v)
				}
			}
			timerv := time.NewTimer(delay * time.Second)
			go func() {
				fmt.Printf("process will start in %ds: %s\n", int(delay), name)
				<-timerv.C
				err := pmgr.ProcManager.StartProcess(name)
				if err != nil {
					fmt.Println(err)
				}
			}()
		} else {
			cmd.Help()
		}
	},
}

// PMStopCmd represents the stop operation on the procmanager command
var PMStopCmd = &cobra.Command{
	Use:   "stop [node name] [optional: delay in secs]",
	Short: "Stops the process named if currently running.",
	Long:  `Stops the process named if currently running.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 && args[0] != "" {
			name := args[0]
			delay := time.Duration(0)
			if len(args) >= 2 {
				if v, e := strconv.ParseInt(args[1], 10, 64); e == nil && v > 0 {
					delay = time.Duration(v)
				}
			}
			timerv := time.NewTimer(delay * time.Second)
			go func() {
				fmt.Printf("process will stop in %ds: %s\n", int(delay), name)
				<-timerv.C
				err := pmgr.ProcManager.StopProcess(name)
				if err != nil {
					fmt.Println(err)
				}
			}()
		} else {
			cmd.Help()
		}
	},
}

// PMKillCmd represents the stop operation on the procmanager command
var PMKillCmd = &cobra.Command{
	Use:   "kill [node name] [optional: delay in secs]",
	Short: "Kills the process named if currently running.",
	Long:  `Kills the process named if currently running.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 && args[0] != "" {
			name := args[0]
			delay := time.Duration(0)
			if len(args) >= 2 {
				if v, e := strconv.ParseInt(args[1], 10, 64); e == nil && v > 0 {
					delay = time.Duration(v)
				}
			}
			timerv := time.NewTimer(delay * time.Second)
			go func() {
				fmt.Printf("process will stop in %ds: %s\n", int(delay), name)
				<-timerv.C
				err := pmgr.ProcManager.KillProcess(name)
				if err != nil {
					fmt.Println(err)
				}
			}()
		} else {
			cmd.Help()
		}
	},
}

// PMKillAllCmd stops all processes in the procmanager
var PMKillAllCmd = &cobra.Command{
	Use:   "killall [optional: delay in secs]",
	Short: "Kills all processes if currently running.",
	Long:  `Kills all processes if currently running.`,
	Run: func(cmd *cobra.Command, args []string) {
		delay := time.Duration(0)
		if len(args) >= 1 {
			if v, e := strconv.ParseInt(args[0], 10, 64); e == nil && v > 0 {
				delay = time.Duration(v)
			}
		}
		timerv := time.NewTimer(delay * time.Second)
		go func() {
			fmt.Printf("all process will stop in %ds\n", int(delay))
			<-timerv.C
			errname, err := pmgr.ProcManager.KillAllProcesses()
			if err != nil {
				fmt.Println("Error on process '" + errname + "': " + err.Error())
			}
		}()
	},
}

// PMStopAllCmd stops all processes in the procmanager
var PMStopAllCmd = &cobra.Command{
	Use:   "stopall [optional: delay in secs]",
	Short: "Stops all processes if currently running.",
	Long:  `Stops all processes if currently running.`,
	Run: func(cmd *cobra.Command, args []string) {
		delay := time.Duration(0)
		if len(args) >= 1 {
			if v, e := strconv.ParseInt(args[0], 10, 64); e == nil && v > 0 {
				delay = time.Duration(v)
			}
		}
		timerv := time.NewTimer(delay * time.Second)
		go func() {
			fmt.Printf("all process will stop in %ds\n", int(delay))
			<-timerv.C
			errname, err := pmgr.ProcManager.StopAllProcesses()
			if err != nil {
				fmt.Println("Error on process '" + errname + "': " + err.Error())
			}
		}()
	},
}

// PMStartAllCmd starts all processes in the procmanager
var PMStartAllCmd = &cobra.Command{
	Use:   "startall [optional: delay in secs]",
	Short: "Starts all processes if currently stopped.",
	Long:  `Starts all processes if currently stopped.`,
	Run: func(cmd *cobra.Command, args []string) {
		delay := time.Duration(0)
		if len(args) >= 1 {
			if v, e := strconv.ParseInt(args[0], 10, 64); e == nil && v > 0 {
				delay = time.Duration(v)
			}
		}
		timerv := time.NewTimer(delay * time.Second)
		go func() {
			fmt.Printf("All processes will start in %ds\n", int(delay))
			<-timerv.C
			errname, err := pmgr.ProcManager.StartAllProcesses()
			if err != nil {
				fmt.Println("Error on process '" + errname + "': " + err.Error())
			}
		}()
	},
}

// PMRemoveCmd represents the list operation on the procmanager command
var PMRemoveCmd = &cobra.Command{
	Use:   "remove [node name] [optional: delay in secs]",
	Short: "Removes the process named.",
	Long:  `Removes the process named. It will stop the process if it is running.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(len(args) >= 1 && args[0] != "") {
			cmd.Help()
		}
		name := args[0]
		delay := time.Duration(0)
		if len(args) >= 2 {
			if v, e := strconv.ParseInt(args[1], 10, 64); e == nil && v > 0 {
				delay = time.Duration(v)
			}
		}
		timerv := time.NewTimer(delay * time.Second)
		go func() {
			fmt.Printf("process removed in %ds: %s\n", int(delay), name)
			<-timerv.C
			err := pmgr.ProcManager.RemoveProcess(name)
			if err != nil {
				fmt.Println(err)
			}
		}()
	},
}

func init() {
	ProcmanagerCmd.AddCommand(PMKillCmd)
	ProcmanagerCmd.AddCommand(PMKillAllCmd)
	ProcmanagerCmd.AddCommand(PMListCmd)
	ProcmanagerCmd.AddCommand(PMMetadataCmd)
	ProcmanagerCmd.AddCommand(PMRemoveCmd)
	ProcmanagerCmd.AddCommand(PMStopCmd)
	ProcmanagerCmd.AddCommand(PMStopAllCmd)
	ProcmanagerCmd.AddCommand(PMStartAllCmd)
	ProcmanagerCmd.AddCommand(PMStartCmd)
}