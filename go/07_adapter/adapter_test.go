package adapter

import "testing"

// go test -v ./...

func TestAdapter(t *testing.T) {
	client := &Client{}

	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windows := &Windows{}
	windowsAdapter := &WindowsAdapter{
		windowsMachine: windows,
	}

	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}