/*
Copyright © 2021 Abdul Fattah Ikhsan <ikhsannetwork@gmail.com>

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
	"os"

	"github.com/ikhsanalatsary/MeowTube/instances"
	"github.com/ikhsanalatsary/MeowTube/interfaces"
	"github.com/ikhsanalatsary/MeowTube/logger"
	"github.com/spf13/cobra"
)

// popularCmd represents the popular command
var popularCmd = &cobra.Command{
	Use:   "popular [no option]",
	Short: "To see popular videos on YouTube",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Request popular...")
		source := instances.FindFastest("/api/v1/popular")
		if source.Error != nil {
			logger.ThrowError(source.Error)
		}
		// resp, err := http.Get(source.FastestURL + "/api/v1/trending" + query)
		defer source.Resp.Body.Close()
		data, err := ioutil.ReadAll(source.Resp.Body)
		if err != nil {
			logger.ThrowError(err)
		}
		res, err := interfaces.UnmarshalVideo(data)
		if err != nil {
			logger.ThrowError(err)
		}
		m, err := res.Marshal()
		if err != nil {
			logger.ThrowError(err)
		}
		// fmt.Println(string(m))
		os.Stdout.Write(m)
	},
}

func init() {
	rootCmd.AddCommand(popularCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// popularCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// popularCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
