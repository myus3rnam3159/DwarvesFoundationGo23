package Cmd

import (
	sortUtils "SortUtils"
	fmt "fmt"
	os "os"

	cobra "github.com/spf13/cobra"
)

var flags = map[string]sortUtils.SortFunc {
	"int": sortUtils.SortIntegers,
	"string": sortUtils.SortStrings,
	"mix": sortUtils.SortMixs,
}

func runCommand(cmd *cobra.Command, args []string) {
	var flagUsed = false

	//Sai chỗ này vì chưa reset flag sau khi gọi
	for k, f := range flags{
		flagUsed, _ = cmd.Flags().GetBool(k)
		if flagUsed{
			f(&args)
			break
		}
	}
	fmt.Print("Output: ")

	for _, e := range args{
		fmt.Print(e + " ")
	}
	fmt.Println()

}

var rootCmd = &cobra.Command{
	Use:   "sort",
	Short: "sort a list",
	Long:  "sort a list of integers or float64s or mix types",
	Run:   runCommand,
}

func Excecute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init(){
	for flagKey:= range flags{
		rootCmd.Flags().Bool(flagKey, false, flagKey + " input list")
	}
}
