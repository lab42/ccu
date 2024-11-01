package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ccu",
	Short: "Validates commit messages against conventional commit format",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the input message to validate
		message := viper.GetString(keyInput)

		// Validate required message
		if message == "" {
			return fmt.Errorf("commit message is required (use --input flag)")
		}

		// Construct the regex pattern using the configured parts
		typePattern := viper.GetString(keyType)
		topicPattern := viper.GetString(keyTopic)
		messagePattern := viper.GetString(keyMessage)

		regex := fmt.Sprintf(defaultTemplate, typePattern, topicPattern, messagePattern)

		// Compile and validate regex
		re, err := regexp.Compile(regex)
		if err != nil {
			return fmt.Errorf("invalid regex pattern: %w", err)
		}

		// Check if message matches pattern
		if !re.MatchString(message) {
			return fmt.Errorf("commit message does not match conventional commit format: %s", message)
		}

		fmt.Println("✅ Commit message is valid")
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccu.yaml)")

	// Local flags for the root command
	rootCmd.Flags().String(keyType, defaultType, "Regular expression pattern for type")
	rootCmd.Flags().String(keyTopic, defaultTopic, "Regular expression pattern for topic")
	rootCmd.Flags().String(keyMessage, defaultMessage, "Regular expression pattern for message")
	rootCmd.Flags().String(keyInput, "", "Commit message to validate")

	// Bind flags to viper
	viper.BindPFlag(keyType, rootCmd.Flags().Lookup(keyType))
	viper.BindPFlag(keyTopic, rootCmd.Flags().Lookup(keyTopic))
	viper.BindPFlag(keyMessage, rootCmd.Flags().Lookup(keyMessage))
	viper.BindPFlag(keyInput, rootCmd.Flags().Lookup(keyInput))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// 1. Set defaults first
	viper.SetDefault(keyType, defaultType)
	viper.SetDefault(keyTopic, defaultTopic)
	viper.SetDefault(keyMessage, defaultMessage)

	// 2. Read config file
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ccu")
	}

	// Read config file if it exists
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// 3. Set up environment variables
	viper.SetEnvPrefix("CCU")
	viper.AutomaticEnv()

	// Configure environment variable mapping
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}
