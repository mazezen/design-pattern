package adapter

import "fmt"

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (wd *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	wd.windowsMachine.insertIntoUSBPort()
}
