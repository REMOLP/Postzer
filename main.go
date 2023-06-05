package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Function to parse the aliases file and create a map of aliases
func parseAliases() map[string]string {
	aliases := make(map[string]string)
	if data, err := ioutil.ReadFile("aliases.pz"); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" {
				parts := strings.SplitN(line, " ", 2)
				if len(parts) == 2 {
					alias := parts[0]
					value := parts[1]
					aliases[alias] = value
				}
			}
		}
	} else {
		fmt.Println("Aliases file not found.")
	}
	return aliases
}

// Function to parse the user-specified .pzp file and replace aliases with values
func generateHTML(filePath string, aliases map[string]string) string {
	htmlOutput := ""

	if data, err := ioutil.ReadFile(filePath); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" {
				parts := strings.SplitN(line, " ", 2)
				alias := parts[0]
				if value, ok := aliases[alias]; ok {
					if len(parts) > 1 {
						replacements := strings.Split(parts[1], ";;")
						for i, replacement := range replacements {
							value = strings.Replace(value, fmt.Sprintf("{{%d}}", i), replacement, -1)
						}
					}
					htmlOutput += value + "\n" // Add a new line after each line of output
				}
			}
		}
	} else {
		fmt.Println("Input file not found.")
	}

	return htmlOutput
}

// Function to replace the placeholder in the template file and save it as the output file
func generatePost(templateFile, outputFile, htmlContent string) {
	if templateData, err := ioutil.ReadFile(templateFile); err == nil {
		template := string(templateData)
		output := strings.Replace(template, "{{replace:me:here}}", htmlContent, -1)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(output))
		if err == nil {
			formattedOutput, _ := doc.Html()

			if err := ioutil.WriteFile(outputFile, []byte(formattedOutput), 0644); err == nil {
				fmt.Println("Post generated and saved successfully.")
			} else {
				fmt.Println("Error saving the output file:", err)
			}
		} else {
			fmt.Println("Error parsing the output HTML:", err)
		}
	} else {
		fmt.Println("Template file not found.")
	}
}

// Main function
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Insufficient arguments.")
		fmt.Println("Usage: ./postzer path/to/example.pzp path/to/myfirstpost.html [path/to/custom-template.html]")
		return
	}

	aliases := parseAliases()
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	templateFile := "init-template.html"
	if len(os.Args) > 3 {
		templateFile = os.Args[3]
	}
	htmlOutput := generateHTML(inputFile, aliases)
	generatePost(templateFile, outputFile, htmlOutput)
}
