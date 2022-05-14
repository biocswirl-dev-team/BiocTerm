/*

Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// export YamlConfig
type YamlConfig struct {
	Connection struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
}

// readlessonCmd represents the readlesson command
var readlessonCmd = &cobra.Command{
	Use:   "readlesson",
	Short: "Parses BiocSwirl lesson.yaml files",
	Long: `This function is designed specifically to parse lesson.yaml files that follow the standard swirl format. Example:

- Class: meta
  Course: RNA-seq analysis
  Lesson: Read_alignment_quantification
  Author: Nicolai von Kügelgen, Julia Philipp
  Type: Standard
  Organization: 
  Version: 2.4.5
- Class: mult_question
  Output: Which of the following statements is true for mapping reads to the transcriptome?
  AnswerChoices: Very fast due to smaller search space;Splice sites are not an issue;Unreliable for species with incomplete annotation;All of the other statements
  CorrectAnswer: All of the other statements
  AnswerTests: omnitest(correctVal= 'All of the other statements')
  Hint: Mapping to the transcriptome uses only the final sequences of all known transcripts from a given genome`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Attempting to read in lesson")

		var fileName string
		flag.StringVar(&fileName, "-f", "", "YAML file to use")
		flag.Parse()

		if fileName == "" {
			fmt.Println("Specify the lesson.yaml")
			return
		}

		// read file
		yamlFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error reading lesson.yaml: %s\n", err)
			return
		}

		// parse using unmarshal
		var yamlConfig YamlConfig
		err = yaml.Unmarshal(yamlFile, &yamlConfig)
		if err != nil {
			fmt.Printf("Error parsing lesson.yaml: %s\n", err)
		}

		// print lesson
		fmt.Printf("Lesson: %v\n", yamlConfig)

	},
}

func init() {
	rootCmd.AddCommand(readlessonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	readlessonCmd.Flags().String("file", "f", "lesson.yaml file to be read in")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readlessonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
