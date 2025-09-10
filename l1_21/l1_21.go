package main

import (
	"encoding/json"
	"fmt"
)

type (
	JSONPrinter interface {
		PrintStyled(text string) string
	}
	Printer interface {
		Print(s string) string
	}

	Client struct {
		printer JSONPrinter
	}
	PrinterAdapter struct {
		printer Printer
	}
	PrintResult struct {
		Print string `json:"print"`
	}
	TextPrinter struct{}
)

func (t *TextPrinter) Print(s string) string {
	return fmt.Sprintf("Text: %s", s)
}

func (p *PrinterAdapter) PrintStyled(text string) string {
	printedText := p.printer.Print(text)
	result := PrintResult{Print: printedText}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return fmt.Sprintf(`{"error": "JSON marshaling failed: %s"}`, err.Error())
	}

	return string(jsonBytes)
}

func NewPrinterAdapter(printer Printer) *PrinterAdapter {
	return &PrinterAdapter{
		printer: printer,
	}
}

func (c *Client) PrintMessage(msg string) {
	result := c.printer.PrintStyled(msg)
	fmt.Println(result)
}

func main() {
	textPrinter := &TextPrinter{}
	adapter := NewPrinterAdapter(textPrinter)
	client := &Client{printer: adapter}

	client.PrintMessage("Hello World")
}
