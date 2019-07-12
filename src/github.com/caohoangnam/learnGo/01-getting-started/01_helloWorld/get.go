package main

/*

for(i=1;i<=15;i++){
    01_0{i}_00
}


Url: https://learning.oreilly.com/api/v1/videoclips/{URL}/?format=json

https://learning.oreilly.com/api/v1/videoclips/01_01_00/?format=json

4

"01_02_00","01_02_01","01_02_02","01_02_03_01","01_02_03_02","01_02_03_03","01_02_03_04","01_02_03_05","01_02_04","01_03_00","01_03_01","01_03_02_01","01_03_02_02","01_03_03_01","01_03_03_02","01_03_03_03","01_03_03_04","01_03_03_05","01_03_03_06","01_03_04","01_04_00","01_04_01_01","01_04_01_02","01_04_02_01","01_04_02_02","01_04_02_03","01_04_03","01_04_04","01_05_00","01_05_01","01_05_02_01","01_05_02_02","01_05_02_03","01_05_03","01_05_04","01_05_05","01_05_06","01_06_00","01_06_01","01_06_02","01_06_03","01_06_04","01_06_05","01_06_06","01_07_00","01_07_01","01_07_02","01_07_03","01_08_00","01_08_01","01_08_02","01_08_03","01_09_00","01_09_01","01_09_02","01_09_03","01_09_04","01_09_05","01_09_06","01_10_00,"01_10_01,"01_10_02_01","01_10_02_02","01_10_02_03","01_10_03","01_10_04_01","01_10_04_02","01_10_05","01_10_06","01_11_00","01_11_01_01","01_11_01_02","01_11_02","01_12_00","01_12_01","01_12_02","01_12_03,"01_12_04","01_12_05,"01_12_06","01_12_07","01_13_00","01_13_01","01_13_02","01_13_03","01_14_00","01_14_01","01_14_02","01_14_03","01_14_04_01","01_14_00_02","01_14_00_03","01_14_00_04","01_14_05"


*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type TimeVideo struct {
	End   string
	Text  string
	Begin string
}

type Lines struct {
	Lines []TimeVideo
}

type Transcription struct {
	Id                    string
	MagpieIranscriptionId int
	Videoclip             string
	Language              string
	Transcription         Lines
}

type VideoOreilly struct {
	Id             string
	ReferenceId    string
	KalturaEntryId string
	MagpieClipId   int
	VideoPlaylist  string
	Transcriptions []Transcription
	WorkId         int
	ReferenceIdUrl string
	ContentFormat  string
	Duration       int
	Description    string
	Title          string
}

var (
	link = []string{"01_01_00", "01_01_01", "01_01_02", "01_01_03", "01_01_04", "01_02_00", "01_02_01", "01_02_02", "01_02_03_01", "01_02_03_02", "01_02_03_03", "01_02_03_04", "01_02_03_05", "01_02_04", "01_03_00", "01_03_01", "01_03_02_01", "01_03_02_02", "01_03_03_01", "01_03_03_02", "01_03_03_03", "01_03_03_04", "01_03_03_05", "01_03_03_06", "01_03_04", "01_04_00", "01_04_01_01", "01_04_01_02", "01_04_01_03", "01_04_02_01", "01_04_02_02", "01_04_02_03", "01_04_03", "01_04_04", "01_05_00", "01_05_01", "01_05_02_01", "01_05_02_02", "01_05_02_03", "01_05_03", "01_05_04", "01_05_05", "01_05_06", "01_06_00", "01_06_01", "01_06_02", "01_06_03", "01_06_04", "01_06_05", "01_06_06", "01_07_00", "01_07_01", "01_07_02", "01_07_03", "01_08_00", "01_08_01", "01_08_02", "01_08_03", "01_09_00", "01_09_01", "01_09_02", "01_09_03", "01_09_04", "01_09_05", "01_09_06", "01_10_00", "01_10_01", "01_10_02_01", "01_10_02_02", "01_10_02_03", "01_10_03", "01_10_04_01", "01_10_04_02", "01_10_05", "01_10_06", "01_11_00", "01_11_01_01", "01_11_01_02", "01_11_02", "01_12_00", "01_12_01", "01_12_02", "01_12_03", "01_12_04", "01_12_05", "01_12_06", "01_12_07", "01_13_00", "01_13_01", "01_13_02", "01_13_03", "01_14_00", "01_14_01", "01_14_02", "01_14_03", "01_14_04_01", "01_14_04_02", "01_14_04_03", "01_14_04_04", "01_14_05"}
)

func main() {
	for i := 0; i < 15; i++ {
		linkAPI := "https://learning.oreilly.com/api/v1/videoclips/9780135261651-UGP2_01_01_" + link[i] + "/?format=json"
		response, err := http.Get(linkAPI)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}

			var videoOreilly VideoOreilly
			json.Unmarshal([]byte(contents), &videoOreilly)

			fileName := videoOreilly.Title + ".vtt"
			f, err := os.Create(fileName)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprintln(f, "WEBVTT")
			fmt.Fprintln(f, "")
			var listVideo = videoOreilly.Transcriptions[0].Transcription.Lines
			fmt.Printf("%s\n", videoOreilly.Title)
			for i := 0; i < len(listVideo); i++ {
				strTimeVideo := listVideo[i].Begin + " --> " + listVideo[i].End
				d := []string{strTimeVideo, listVideo[i].Text, ""}
				fmt.Fprintln(f, d[0])
				fmt.Fprintln(f, d[1])
				fmt.Fprintln(f, d[2])
				fmt.Printf("%s\n", "Success")
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("file written successfully" + videoOreilly.Title)
		}
	}
}
