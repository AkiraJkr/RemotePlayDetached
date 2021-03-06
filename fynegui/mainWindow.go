package fynegui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var rpd fyne.App
var mainContainer *fyne.Container
var mainWindow fyne.Window

//Run func
func Run() {
	rpd = app.New()
	mainWindow = rpd.NewWindow("Remote Play Detached")
	mainWindow.Resize(fyne.NewSize(600, 350))
	mainWindow.SetContent(buildMainContent())
	mainWindow.SetMaster()
	mainWindow.ShowAndRun()
}

func buildMainContent() *fyne.Container {
	mainWindow.SetMainMenu(buildMainMenu())
	mainContainer = fyne.NewContainerWithLayout(layout.NewVBoxLayout(), buildAppListContainer(), layout.NewSpacer(), buildButtonBar())
	return mainContainer
}

func refreshContent() {
	mainWindow.SetContent(buildMainContent())
}

func buildMainMenu() *fyne.MainMenu {

	mainMenu := fyne.NewMainMenu(fyne.NewMenu("Menu",
		fyne.NewMenuItem("About", func() {
			fmt.Println("clicked: About")
			aboutWindow := rpd.NewWindow("About")
			aboutWindow.Resize(fyne.NewSize(500, 400))

			licenseLabel := widget.NewLabel(LICENSE)
			scrollContainer := widget.NewScrollContainer(licenseLabel)

			okButton := widget.NewButton("OK", func() {

				aboutWindow.Close()

				refreshContent()
			})

			buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), okButton, layout.NewSpacer())
			paragraphContainer := fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(495, 380)), scrollContainer)
			content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), paragraphContainer, layout.NewSpacer(), buttons)
			aboutWindow.SetContent(content)
			aboutWindow.Show()
		}),
	))
	return mainMenu
}

func buildButtonBar() *fyne.Container {

	importButton := widget.NewButton("Import", func() {
		fmt.Println("clicked: Import")
		importApp()
	})

	exitButton := widget.NewButton("Exit", func() {
		fmt.Println("clicked: Exit")
		rpd.Quit()
	})

	buttonBar := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), importButton, layout.NewSpacer(), exitButton)

	return buttonBar
}
