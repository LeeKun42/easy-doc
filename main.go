package main

import "easy-doc/app/cmd"

func main() {
	cmd.Main.AddCommand(&cmd.GormDto)
	cmd.Main.Execute()
}
