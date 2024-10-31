package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	Version string
	Commit  string
	Date    string
	JSON    bool
)

type Info struct {
	Version string `json:"version" yaml:"version"`
	Commit  string `json:"commit" yaml:"commit"`
	Date    string `json:"date" yaml:"date"`
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version information (Defaults to Yaml output)",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		info := Info{
			Version: Version,
			Commit:  Commit,
			Date:    Date,
		}

		switch true {
		case JSON:
			jsonData, err := json.Marshal(&info)
			if err != nil {
				return err
			}
			fmt.Println(string(jsonData))
		default:
			yamlData, err := yaml.Marshal(info)
			if err != nil {
				return err
			}
			fmt.Println(string(yamlData))
		}

		return
	},
}

func init() {
	versionCmd.Flags().BoolVar(&JSON, "json", false, "Output in JSON format")
	rootCmd.AddCommand(versionCmd)
}
