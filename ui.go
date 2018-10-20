package main

import "github.com/andlabs/ui"

func setupUI() {
	// Setup window
	mainWindow := ui.NewWindow("Guffer", 640, 480, false)
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

	// Load and save buttons
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
			ui.MsgBox(mainWindow, "Success", "The file has been saved successfully")
		}
	})
	apiKeysButtons.Append(openButton, false)
	apiKeysButtons.Append(ui.NewLabel(""), true)
	apiKeysButtons.Append(saveButton, false)
	vBox.Append(ui.NewLabel(""), false)
	vBox.Append(apiKeysButtons, false)
	// Tweets group
	vBox.Append(ui.NewLabel(""), false)
	tweets := ui.NewGroup("Tweets")
	vBox.Append(tweets, true)
	vBox.Append(ui.NewLabel(""), false)
	addTweetButton := ui.NewButton("Add tweet")
	addTweetButton.OnClicked(func(button *ui.Button) {
		openAddTweetWindow()
	})
	vBox.Append(addTweetButton, false)
	vBox.Append(ui.NewLabel(""), false)
	startButton := ui.NewButton("Start guffer")
	vBox.Append(startButton, false)
	vBox.Append(ui.NewLabel(""), false)
	// About text
	vBox.Append(ui.NewLabel(""), false)
	about := ui.NewLabel("Guffer is open-source project! Contribute at https://github.com/mrichman/guffer")
	vBox.Append(about, false)
}

func openAddTweetWindow() {
	// Setup window
	addTweetWindow := ui.NewWindow("Add scheduled tweet", 400, 200, false)
	// Close callback
	addTweetWindow.OnClosing(func(window *ui.Window) bool {
		addTweetWindow.Destroy()
		ui.Quit()
		return false
	})
	// Also close callback
	ui.OnShouldQuit(func() bool {
		addTweetWindow.Destroy()
		return true
	})
	addTweetWindow.Show()
}
