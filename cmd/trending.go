/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ikhsanalatsary/MeowTube/instances"
	"github.com/ikhsanalatsary/MeowTube/interfaces"
	"github.com/spf13/cobra"
)

var region string

// var source instances.FastestInstance
var trendingType string
var trendingTypes = map[string]string{
	"music":  "music",
	"gaming": "gaming",
	"news":   "news",
	"movies": "movies",
}

// trendingCmd represents the trending command
var trendingCmd = &cobra.Command{
	Use:   "trending",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Request trending...")
		fmt.Println(args)
		fmt.Println("Region: " + region)
		var query string
		if len(region) == 2 {
			query = "?region=" + region
		}
		if len(trendingType) > 0 {
			if _, ok := trendingTypes[trendingType]; !ok {
				fmt.Println("Invalid type")
				os.Exit(1)
			}
			if strings.Contains(query, "?") {
				query += "&type=" + trendingType
			} else {
				query += "?type=" + trendingType
			}
		}
		fmt.Println("query: ", query)
		source := instances.FindFastest("/api/v1/trending" + query)
		if source.Error != nil {
			log.Fatal(source.Error)
		}
		fmt.Println("Source: " + source.FastestURL)
		// resp, err := http.Get(source.FastestURL + "/api/v1/trending" + query)
		defer source.Resp.Body.Close()
		data, err := ioutil.ReadAll(source.Resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		res, err := interfaces.UnmarshalVideo(data)
		if err != nil {
			log.Fatal(err)
		}
		m, err := res.Marshal()
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(string(m))
		os.Stdout.Write(m)
	},
}

func init() {
	rootCmd.AddCommand(trendingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trendingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trendingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	trendingCmd.Flags().StringVarP(&region, "region", "r", "", "To see trendings in a specific region in format ISO 3166 country code (default: \"US\")")
	trendingCmd.Flags().StringVarP(&trendingType, "type", "t", "", "To see trendings in a specific type (\"music\", \"gaming\", \"news\", \"movies\")")
}
