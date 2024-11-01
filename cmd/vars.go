package cmd

const (
	keyType    = "type"
	keyTopic   = "topic"
	keyMessage = "message"
	keyInput   = "input"

	defaultTemplate = `^(%s)%s:%s`
	defaultType     = `build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test`
	defaultTopic    = `(\([a-zA-Z0-9\-\.]+\))?(!)?`
	defaultMessage  = `.*`
)

var (
	cfgFile string
	Version string
	Commit  string
	Date    string
)
