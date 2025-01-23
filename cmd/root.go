package cmd

import (
	"mydriveuploader/drive"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mydriveuploader",
	Short: "MyDriveUploader is a CLI tool for uploading files to Google Drive",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(drive.UploadCmd)
}