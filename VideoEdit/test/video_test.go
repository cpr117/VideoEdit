// @User CPR
package test

import (
	"VideoEdit/utils"
	"fmt"
	"testing"
)

func TestVideoCut(t *testing.T) {
	testlist := []string{"mp4", "flv", "mov", "wmv", "mkv", "avi"}
	for _, format := range testlist {
		filename := fmt.Sprintf("./video/upload/%s_test.%s", format, format)
		_result := utils.VM.CutVideo(filename, "0:10", "5")
		fmt.Println(_result)
	}
}
