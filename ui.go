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
	vBox := ui.NewVerticalBox()
	apiKeys := ui.NewGroup("Api keys")
	apiKeys.SetChild(form)
	vBox.Append(apiKeys, false)
	mainWindow.SetChild(vBox)
	mainWindow.SetMargined(true)
	defer mainWindow.Show()

	// Consumer key entry
	consumerKey := ui.NewEntry()
	form.Append("", ui.NewLabel(""), false)
	form.Append("Consumer key:  ", consumerKey, false)

	// Consumer password entry
	consumerSecretEntry := ui.NewPasswordEntry()
	form.Append("", ui.NewLabel(""), false)
	form.Append("Consumer secret:  ", consumerSecretEntry, false)

	// Access token
	accessToken := ui.NewEntry()
	form.Append("", ui.NewLabel(""), false)
	form.Append("Acces token:  ", accessToken, false)

	// Access token secret
	accessTokenSecret := ui.NewPasswordEntry()
	form.Append("", ui.NewLabel(""), false)
	form.Append("Access token:  ", accessTokenSecret, false)
	form.Append("", ui.NewLabel(""), false)

	// Load and save button
	apiKeysButtons := ui.NewHorizontalBox()
	openButton := ui.NewButton("Open saved auth file")
	openButton.OnClicked(func(button *ui.Button) {
		ui.OpenFile(mainWindow)
	})
	saveButton := ui.NewButton("Save auth file")
	saveButton.OnClicked(func(button *ui.Button) {
		fileLocation := ui.SaveFile(mainWindow)
		keys := TwitterAuthKeys{
			ConsumerKey:       consumerKey.Text(),
			ConsumerSecret:    consumerSecretEntry.Text(),
			AccessToken:       accessToken.Text(),
			AccessTokenSecret: accessTokenSecret.Text(),
		}
		err := keys.saveToTomlFile(fileLocation)
		if err != nil {
			ui.MsgBoxError(mainWindow, "Error!", "Failed to save file:"+err.Error())
		} else {
			ui.MsgBox(mainWindow, "Succes", "The file has been saved successfully")
		}
	})
	apiKeysButtons.Append(openButton, false)
	apiKeysButtons.Append(saveButton, false)
	vBox.Append(apiKeysButtons, false)
}
