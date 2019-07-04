package main

import (
	"../powershell"
	"fmt"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	fmt.Println("starting the project")

	main_app := app.New()

	w := main_app.NewWindow("iPerf 3 - WSL")

	iPerf3AddressWidget := widget.NewEntry()
	iPerf3PortWidget := widget.NewEntry()
	iPerf3SpeedWidget := widget.NewEntry()

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Welcome to the automated iPerf tool for Windows WSL1"),
		widget.NewButton("Enable WSL", func() {
			EnableWSL()
		}),
		widget.NewButton("Install Ubuntu on WSL", func() {
			InstallUbuntu()
		}),
		widget.NewButton("Install iPerf3 on WSL", func() {
			InstallIPerf3()
		}),
		iPerf3AddressWidget,
		iPerf3PortWidget,
		iPerf3SpeedWidget,
		widget.NewButton("Run iPerf3 Tests", func() {
			go powershell.RunIperf3(iPerf3AddressWidget.Text, iPerf3PortWidget.Text, iPerf3SpeedWidget.Text)
		}),
		widget.NewButton("Quit", func() {
			main_app.Quit()
		}),
	))

	w.ShowAndRun()
}

func InstallIPerf3() {
	powershell.InstallIPerf3WSL()
}

func EnableWSL() {
	powershell.EnableWSL()
}

func InstallUbuntu() {
	powershell.InstallUbuntuWSL()
}