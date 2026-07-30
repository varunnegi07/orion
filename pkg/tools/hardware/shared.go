package hardwaretools

import toolshared "github.com/varunnegi07/orion/pkg/tools/shared"

type ToolResult = toolshared.ToolResult

func ErrorResult(message string) *ToolResult {
	return toolshared.ErrorResult(message)
}

func SilentResult(forLLM string) *ToolResult {
	return toolshared.SilentResult(forLLM)
}
