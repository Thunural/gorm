package gen

import (
	"gormui/config"
	"os"
	"path/filepath"
)

func Check(c config.Param) error {
	//path := c.OutFiles + "/" + c.ProjectName
	currDir, _ := filepath.Abs(c.OutFiles)
	checkDir := filepath.Join(currDir, c.ProjectName)
	_, err := os.Stat(checkDir)
	if err != nil {
		return err
	} else {
		return nil
	}
}
