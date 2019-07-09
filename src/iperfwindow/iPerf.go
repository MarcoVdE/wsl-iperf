package iperfwindow

import (
	"../powershell"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func IPerfWindow() {
	iperf_app := app.New()

	w := iperf_app.NewWindow("iPerf 3 WSL - Test")

	iPerf3Console := widget.NewEntry()

	iPerf3AddressWidget := widget.NewEntry()
	iPerf3PortWidget := widget.NewEntry()
	iPerf3SpeedWidget := widget.NewEntry()

	w.SetContent(widget.NewVBox(
		widget.NewLabel("iPerf3 Window"),
		widget.NewLabel("Address:"),
		iPerf3AddressWidget,
		widget.NewLabel("Port"),
		iPerf3PortWidget,
		widget.NewLabel("Speed in Mbps (max 80% of line speed"),
		iPerf3SpeedWidget,
		widget.NewButton("Run iPerf3 Tests", func() {
			go iPerf3Console.SetText(powershell.RunIperf3Output(iPerf3AddressWidget.Text, iPerf3PortWidget.Text, iPerf3SpeedWidget.Text))
		}),
		iPerf3Console,
		widget.NewButton("Quit", func() {
			iperf_app.Quit()
		}),
	))

	w.ShowAndRun()
}
