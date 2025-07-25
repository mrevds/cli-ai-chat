package utils

import (
    "fmt"
    "strings"
)

// ANSI color codes
const (
    ColorReset  = "\033[0m"
    ColorRed    = "\033[31m"
    ColorGreen  = "\033[32m"
    ColorYellow = "\033[33m"
    ColorBlue   = "\033[34m"
    ColorPurple = "\033[35m"
    ColorCyan   = "\033[36m"
    ColorWhite  = "\033[37m"
    ColorBold   = "\033[1m"
    ColorDim    = "\033[2m"
)

func PrintBanner() {
    fmt.Printf("%s%s", ColorCyan, ColorBold)
    fmt.Println("╔══════════════════════════════════════╗")
    fmt.Println("║          🤖 AIMLAPI Chat CLI          ║")
    fmt.Println("║                                      ║")
    fmt.Println("║     Powered by GPT-4o & Go Lang     ║")
    fmt.Println("╚══════════════════════════════════════╝")
    fmt.Printf("%s", ColorReset)
    fmt.Println()
}

func PrintPrompt() {
    fmt.Printf("%s%s❯ %s", ColorGreen, ColorBold, ColorReset)
}

func PrintResponse(content string) {
    fmt.Printf("\n%s%s🤖 Ассистент:%s\n", ColorBlue, ColorBold, ColorReset)
    fmt.Printf("%s┌─%s\n", ColorDim, ColorReset)
    
    lines := strings.Split(content, "\n")
    for _, line := range lines {
        fmt.Printf("%s│ %s%s\n", ColorDim, line, ColorReset)
    }
    fmt.Printf("%s└─%s\n\n", ColorDim, ColorReset)
}

func PrintError(err error) {
    fmt.Printf("%s%s❌ Ошибка: %v%s\n\n", ColorRed, ColorBold, err, ColorReset)
}

func PrintInfo(message string) {
    fmt.Printf("%s%sℹ️  %s%s\n", ColorYellow, ColorBold, message, ColorReset)
}

func PrintSuccess(message string) {
    fmt.Printf("%s%s✅ %s%s\n", ColorGreen, ColorBold, message, ColorReset)
}

func PrintSeparator() {
    fmt.Printf("%s%s\n", ColorDim, strings.Repeat("─", 50))
    fmt.Printf("%s", ColorReset)
}