package groupie

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Graphique() {
	newapp := app.New()
	windows := newapp.NewWindow("Groupie Tracker !")

	windows.SetContent(widget.NewLabel("Hello !"))
	test := container.NewVBox()

	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Entrez votre recherche !")

	windows.Resize(fyne.NewSize(500, 500))
	artists := Api()

	for _, art := range artists {
		name := widget.NewLabel(art.Name)
		firstalbum := widget.NewLabel(art.FirstAlbum)
		locations := widget.NewLabel(art.Locations)
		concertsdates := widget.NewLabel(art.ConcertDates)
		relations := widget.NewLabel(art.Relations)

		var membersString string
		for _, member := range art.Members {
			membersString += member
		}

		creationdatestring := strconv.Itoa(int(art.CreationDate))

		members := widget.NewLabel(membersString)
		creationDate := widget.NewLabel(creationdatestring)

		test.Add(
			container.NewVBox(
				name,
				members,
				creationDate,
				firstalbum,
				locations,
				concertsdates,
				relations,
			),
		)
	}

	c := container.NewVScroll(test)
	windows.SetContent(c)

	windows.ShowAndRun()
}
