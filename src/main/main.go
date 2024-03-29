package main

import (
	"../powershell"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"log"
	"strconv"
	"time"
)

func main() {
	fmt.Println("starting the project")

	main_app := app.New()

	w := main_app.NewWindow("The Windows Linux Network Tool")

	//iPerf Tab labels and widgets.
	iPerf3AddressLabel := widget.NewLabelWithStyle("Address: ", fyne.TextAlignLeading, fyne.TextStyle{})
	//iPerf3AddressLabel := widget.NewLabel("Address: ")
	iPerf3AddressWidget := widget.NewEntry()
	iPerf3AddressWidget.SetPlaceHolder("ping.online.net")

	iPerf3PortLabel := widget.NewLabel("Port: ")
	iPerf3PortWidget := widget.NewEntry()
	iPerf3PortWidget.SetPlaceHolder("5001")

	iPerf3SpeedLabel := widget.NewLabel("Speed (Mbps): ")
	iPerf3SpeedWidget := widget.NewEntry()
	iPerf3SpeedWidget.SetPlaceHolder("80")

	iPerf3VerboseWidget := widget.NewCheck("Verbose: ", nil) //Check fyne-io/widget/check_test.go:14

	iPerf3ReverseWidget := widget.NewCheck("Reverse (download): ", nil) //Check fyne-io/widget/check_test.go:14

	iPerf3TimeLabel := widget.NewLabel("Time (seconds): ")
	iPerf3TimeWidget := widget.NewEntry()
	iPerf3TimeWidget.SetPlaceHolder("10")

	iPerf3OmitLabel := widget.NewLabel("Omit (seconds): ")
	iPerf3OmitWidget := widget.NewEntry()
	iPerf3OmitWidget.SetPlaceHolder("3")

	iPerf3isUDPWidget := widget.NewCheck("UDP: ", nil) //Check fyne-io/widget/check_test.go:14

	iPerf3Output := widget.NewMultiLineEntry()
	iPerf3Output.SetReadOnly(true)
	iPerf3Output.SetPlaceHolder("Output Window. \n\n\nPress Test to start test")

	iPerfLayout := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		iPerf3AddressLabel, iPerf3AddressWidget,
		iPerf3PortLabel, iPerf3PortWidget,
		iPerf3SpeedLabel, iPerf3SpeedWidget,
		iPerf3VerboseWidget, iPerf3ReverseWidget,
		iPerf3TimeLabel, iPerf3TimeWidget,
		iPerf3OmitLabel, iPerf3OmitWidget,
		iPerf3isUDPWidget,
	)

	w.SetContent(
		widget.NewVBox(

			widget.NewTabContainer(
				widget.NewTabItem("Home",
					widget.NewVBox(
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
					),
				),
				widget.NewTabItem("iPerf Test",
					widget.NewVBox(
						iPerfLayout,
						iPerf3Output,
						widget.NewButton("Run Test", func() {
							//get ints and test valid.
							bandwidth, err := strconv.Atoi(iPerf3SpeedWidget.Text)
							if err != nil {
								log.Printf("Bandiwdth Error: %s", err)
							}
							port, err := strconv.Atoi(iPerf3PortWidget.Text)
							if err != nil {
								log.Printf("Port Conversion Error: %s", err)
							}
							iperfTime, err := strconv.Atoi(iPerf3TimeWidget.Text)
							if err != nil {
								log.Printf("Time Conversion Error: %s", err)
							}
							omit, err := strconv.Atoi(iPerf3OmitWidget.Text)
							if err != nil {
								log.Printf("Time Conversion Error: %s", err)
							} else if omit > iperfTime {
								log.Printf("Omit(%is) longer than time(%is)", omit, iperfTime)
							}

							//TODO: rewrite the result as channel bringing back info and ticking the update every 0.3s
							powershellChannel := make(chan string)
							go iPerfTest(iPerf3AddressWidget.Text, port,
								bandwidth, iPerf3VerboseWidget.Checked, iPerf3ReverseWidget.Checked,
								iperfTime, omit, iPerf3isUDPWidget.Checked, powershellChannel)

							go setTextFromChannel(iPerf3Output, powershellChannel)

						}),
						widget.NewButton("Copy Result", func() {
							//get the result from output window.
							if iPerf3Output.Text != "" {
								clipboard := fyne.Clipboard(w.Clipboard())
								clipboard.SetContent(iPerf3Output.Text)
							} else {
								dialog.ShowInformation("No results", "No text to copy.", w)
							}
						}),
					),
				),

				widget.NewTabItem("ISP",
					widget.NewVBox(
						widget.NewLabelWithStyle("Default ISP tests", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
						widget.NewButton("Atomic CT FF", func() {}),
						widget.NewButton("Cool Ideas - JHB", func() {}),
						widget.NewButton("Cool Ideas - CT", func() {}),
					),
				),
			),

			widget.NewButton("Quit", func() {
				main_app.Quit()
			}),
		)) //end of Main VBox and window content.
	w.ShowAndRun()
}

func iPerfTest(address string, port int, bandwidth int, verbose bool, reverse bool, time int, omit int, isUDP bool, channel chan string) {
	channel <- powershell.RunIPerf3Test(powershell.NewIPerfObject(address, port,
		bandwidth, verbose, reverse,
		time, omit, isUDP))
}

func setTextFromChannel(output *widget.Entry, channel chan string) {
	output.SetText(<-channel)
	time.Sleep(1000)
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
