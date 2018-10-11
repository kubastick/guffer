package main

import "github.com/andlabs/ui"

func setupUI() {
	// Setup window
	mainWindow := ui.NewWindow("Guffer", 640, 480, true)
	// Close callback
	mainWindow.OnClosing(func(window *ui.Window) bool {
		mainWindow.Destroy()
		ui.Quit()
		return false
	})
	// Also close callback
	ui.OnShouldQuit(func() bool {
		mainWindow.Destroy()
		return true
	})
	// Layout
	form := ui.NewForm()
	apiKeys := ui.NewGroup("Api keys")
	apiKeys.SetChild(form)
	mainWindow.SetChild(apiKeys)
	mainWindow.SetMargined(true)
	defer mainWindow.Show()

	// Consumer key entry
	consumerKeyEntry := ui.NewEntry()
	form.Append("", ui.NewLabel(""), false)
	form.Append("Consumer key:  ", consumerKeyEntry, false)

	// Consumer password entry
	consumerSecretEntry := ui.NewPasswordEntry()
	form.Append("", ui.NewLabel(""), false)
	form.Append("Consumer secret:  ", consumerSecretEntry, false)

}
