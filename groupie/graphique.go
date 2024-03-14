package groupie

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	newapp := app.New()
	windows := newapp.NewWindow("Groupie Tracker !")

	windows.SetContent(widget.NewLabel("Hello !"))

	windows.Resize(fyne.NewSize(500, 500))
	artists := api()
	for _, art := range artists {

		label1 := canvas.NewText("Groupie Tracker", color.White)
		label2 := canvas.NewText(fmt.Sprint("Name", art.Name), color.White)
		label3 := canvas.NewText(fmt.Sprint("Members", art.Members), color.White)
		label4 := canvas.NewText(fmt.Sprint("CreationDate", art.CreationDate), color.White)
		label5 := canvas.NewText(fmt.Sprint("FirstAlbum", art.FirstAlbum), color.White)
		label6 := canvas.NewText(fmt.Sprint("Locations", art.Locations), color.White)
		label7 := canvas.NewText(fmt.Sprint("ConcertDates", art.ConcertDates), color.White)
		label8 := canvas.NewText(fmt.Sprint("Relations", art.Relations), color.White)

		windows.SetContent(
			container.NewVBox(
				label1,
				label2,
				label3,
				label4,
				label5,
				label6,
				label7,
				label8,
			),
		)
	}

	windows.ShowAndRun()
}
