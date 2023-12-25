package version

import "fmt"

var Version = "1.0"
var BuildTime = "2022-01-01 00:00"
var CommitHash = ""

func GetVersion() {

	fmt.Println("Version:\t" + Version)
	fmt.Println("Built @:\t" + BuildTime)
	fmt.Println("Commit:\t" + CommitHash)
}
