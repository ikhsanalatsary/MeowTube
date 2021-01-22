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

	url "net/url"

	"github.com/ikhsanalatsary/MeowTube/instances"
	"github.com/ikhsanalatsary/MeowTube/interfaces"
	"github.com/spf13/cobra"
)

var sortBy = map[string]string{
	"relevance":   "relevance",
	"rating":      "rating",
	"upload_date": "upload_date",
	"view_count":  "view_count",
}

var durations = map[string]string{
	"long":  "long",
	"short": "short",
}

var searchTypes = map[string]string{
	"video":    "video",
	"playlist": "playlist",
	"channel":  "channel",
	"all":      "all",
}

var searchDate = map[string]string{
	"hour":  "hour",
	"today": "today",
	"week":  "week",
	"month": "month",
	"year":  "year",
}

// var q string
var searchRegion string
var searchType string
var page int32
var sort string
var date string
var duration string
var features string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search \"search criteria\"",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.
q: String
page: Int32
sort_by: "relevance", "rating", "upload_date", "view_count"
date: "hour", "today", "week", "month", "year"
duration: "short", "long"
type: "video", "playlist", "channel", "all", (default: video)
features: "hd", "subtitles", "creative_commons", "3d", "live", "purchased", "4k", "360", "location", "hdr" (comma separated: e.g. "&features=hd,subtitles,3d,live")
region: ISO 3166 country code (default: "US")`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var query string = "?q=" + url.QueryEscape(args[0]) + "&fields=type,title,author,videoId,playlistId,authorId,publishedText"
		if len(region) == 2 {
			query += "&region=" + region
		}
		if len(searchType) > 0 {
			if _, ok := searchTypes[searchType]; !ok {
				fmt.Println("Invalid type")
				os.Exit(1)
			}
			query += "&type=" + searchType
		}
		if page > 0 {
			query += "&page=" + string(page)
		}
		if len(sort) > 0 {
			if _, ok := sortBy[sort]; !ok {
				fmt.Println("Invalid sort criteria")
				os.Exit(1)
			}
			query += "&sort_by=" + sort
		}
		if len(date) > 0 {
			if _, ok := searchDate[date]; !ok {
				fmt.Println("Invalid date criteria")
				os.Exit(1)
			}
			query += "&date=" + date
		}
		if len(duration) > 0 {
			if _, ok := durations[duration]; !ok {
				fmt.Println("Invalid duration criteria")
				os.Exit(1)
			}
			query += "&duration=" + duration
		}
		if len(features) > 0 {
			if _, ok := searchDate[date]; !ok {
				fmt.Println("Invalid date criteria")
				os.Exit(1)
			}
			query += "&features=" + features
		}
		fmt.Println("query: ", query)
		source, err := instances.FindFastest(&instances.InstanceList, "/api/v1/search"+query)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Source: " + source.FastestURL)
		defer source.Resp.Body.Close()
		data, err := ioutil.ReadAll(source.Resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		res, err := interfaces.UnmarshalSearch(data)
		if err != nil {
			log.Fatal(err)
		}
		m, err := res.Marshal()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(m))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// searchCmd.Flags().StringVar(&q, "q", "", "Search criteria")
	searchCmd.Flags().Int32VarP(&page, "page", "p", 0, "Request on specific page")
	searchCmd.Flags().StringVarP(&sort, "sortby", "s", "", "Sort criteria, e.g: \"relevance\", \"rating\", \"upload_date\", \"view_count\"")
	searchCmd.Flags().StringVarP(&date, "date", "D", "", "date criteria, e.g: \"hour\", \"today\", \"week\", \"month\", \"year\"")
	searchCmd.Flags().StringVarP(&duration, "duration", "d", "", "duration criteria, e.g: \"short\", \"long\"")
	searchCmd.Flags().StringVarP(&searchType, "type", "t", "", "type criteria, e.g: \"video\", \"playlist\", \"channel\", \"all\", (default: video)")
	searchCmd.Flags().StringVarP(&region, "region", "r", "", "To see search results in a specific region in format ISO 3166 country code (default: \"US\")")
	searchCmd.Flags().StringVarP(&features, "features", "f", "", "\"hd\", \"subtitles\", \"creative_commons\", \"3d\", \"live\", \"purchased\", \"4k\", \"360\", \"location\", \"hdr\" (comma separated: e.g. \"&features=hd,subtitles,3d,live\")")
	// searchCmd.MarkFlagRequired("q")
}
