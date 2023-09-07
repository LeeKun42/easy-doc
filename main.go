package main

import "api-doc/app/cmd"

func main() {
	cmd.Main.AddCommand(&cmd.GormDto)
	cmd.Main.Execute()
}
