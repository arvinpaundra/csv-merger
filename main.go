package main

import (
	"os"
	"time"

	"github.com/arvinpaundra/csv-merger/internal/core"
	"github.com/spf13/cobra"
)

var (
	source    string
	target    string
	sourceKey string
	targetKey string
	out       string
)

var rootCmd = &cobra.Command{
	Use: "csv-merger",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
		HiddenDefaultCmd:  true,
	},
	Short: "a command-line tool to merge two CSV files based on specified key columns.",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		csv1, err := core.ReadCsv(source)
		if err != nil {
			core.LogError("failed read source csv: ", err.Error())
			os.Exit(1)
		}

		csv2, err := core.ReadCsv(target)
		if err != nil {
			core.LogError("failed read target csv: ", err.Error())
			os.Exit(1)
		}

		keyFromCsv1, err := core.FindKeyIndex(sourceKey, csv1)
		if err != nil {
			core.LogError("failed finding key from source csv: ", err.Error())
			os.Exit(1)
		}

		keyFromCsv2, err := core.FindKeyIndex(targetKey, csv2)
		if err != nil {
			core.LogError("failed finding key from target csv: ", err.Error())
			os.Exit(1)
		}

		results, err := core.MergeCsv(keyFromCsv1, keyFromCsv2, csv1, csv2)
		if err != nil {
			core.LogError("failed merge file: ", err.Error())
			os.Exit(1)
		}

		err = core.WriteCsv(out, results)
		if err != nil {
			core.LogError("failed save csv: ", err.Error())
			os.Exit(1)
		}

		core.LogSuccess("finished in ", time.Since(start).Milliseconds(), "ms")
	},
}

func main() {
	rootCmd.Flags().StringVar(&source, "source", "", "path to the main csv file")
	rootCmd.Flags().StringVar(&target, "target", "", "path to the csv file to merge with the source")
	rootCmd.Flags().StringVar(&sourceKey, "source-key", "", "column name in the source file to match on")
	rootCmd.Flags().StringVar(&targetKey, "target-key", "", "column name in the target file to match on")
	rootCmd.Flags().StringVar(&out, "out", "result.csv", "specify the output file")

	rootCmd.Execute()
}
