package cmd

import (
    "log"
    "github.com/spf13/cobra"
)

var cmd = &cobra.Command{
    Use: "Volumetric Cloud",
    Short: "Cloud",
}

func Execute() {
    if err := cmd.Execute(); err != nil {
        log.Fatalln(err.Error())
    }
}
