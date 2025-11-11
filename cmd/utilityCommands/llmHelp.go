package utilityCommands

import (
	"fmt"

	"github.com/samhodg1993/toto/internal/embedded"
	"github.com/spf13/cobra"
)

var LlmHelpCmd = &cobra.Command{
	Use:     "llm-help",
	Aliases: []string{"llm"},
	Short:   "Display comprehensive LLM-focused usage documentation",
	Long: `Outputs the complete LLMs.txt usage guide designed for Large Language Models.

This documentation provides detailed command usage, examples, workflows, and tips
for working with toto. Perfect for AI assistants like Claude Code.

The same content is also available at: ~/.config/toto/LLMs.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(embedded.LLMUsageDoc)
	},
}
