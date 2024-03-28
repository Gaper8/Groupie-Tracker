package groupie

import (
	"fmt"
	"strconv"
	"strings"

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

	artists, err := Api()
	if err != nil {
		fmt.Println("Erreur", err)
		return
	}

	locationsData, err := LocationApi()
	if err != nil {
		fmt.Println("Erreur", err)
		return
	}

	for _, art := range artists {
		name := widget.NewLabel("Nom de l'artiste : " + art.Name)
		firstalbum := widget.NewLabel("Album : " + art.FirstAlbum)
		locations := widget.NewLabel("Lieu concert : " + art.Locations)
		concertsdates := widget.NewLabel("Dates de concert : " + art.ConcertDates)
		relations := widget.NewLabel(art.Relations)

		var membersString string
		for i, member := range art.Members {
			membersString += member
			if i < len(art.Members)-1 {
				membersString += ", "
			}
		}

		creationdatestring := strconv.Itoa(int(art.CreationDate))

		members := widget.NewLabel("Membres : " + membersString)
		creationDate := widget.NewLabel("Date de création : " + creationdatestring)

		for _, location := range locationsData.Index {

			locationsLabel := widget.NewLabel("Locations: " + strings.Join(location.Locations, ", "))
			datesLabel := widget.NewLabel("Dates: " + location.Dates)

			test.Add(
				container.NewVBox(
					name,
					members,
					creationDate,
					firstalbum,
					locationsLabel,
					locations,
					concertsdates,
					relations,
					datesLabel,
				),
			)
		}

		c := container.NewVScroll(test)
		windows.SetContent(c)

		windows.ShowAndRun()
	}
}
