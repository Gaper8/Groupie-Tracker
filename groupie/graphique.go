package groupie

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Graphique() {
	newApp := app.New()
	windows := newApp.NewWindow("Groupie Tracker !")

	windows.Resize(fyne.NewSize(1200, 800))
	pageglobalartist(windows)
	windows.ShowAndRun()
}

func pageglobalartist(mainpage fyne.Window) {
	artists, err := Api()
	if err != nil {
		fmt.Println("Erreur", err)
		return
	}
	listbuttonartist := make([]fyne.CanvasObject, 0, len(artists))
	for _, art2 := range artists {
		art := art2
		button := widget.NewButton(art.Name, func() {
			showdataartist(mainpage, art)
		})
		listbuttonartist = append(listbuttonartist, button)
	}
	scrollableList := container.NewVScroll(container.NewVBox(listbuttonartist...))
	mainpage.SetContent(scrollableList)
}

func showdataartist(mainpage fyne.Window, artist ArtisteElement) {
	locationData, err := LocationApi()
	if err != nil {
		fmt.Println("Erreur", err)
		return
	}

	datesData, err := DatesApi()
	if err != nil {
		fmt.Println("Erreur", err)
		return
	}

	load, err := fyne.LoadResourceFromURLString(artist.Image)
	if err != nil {
		fmt.Print(err)
	}

	var artistLocations []string
	for _, index := range locationData.Index {
		if index.ID == artist.ID {
			artistLocations = index.Locations
			break
		}
	}

	var artistConcertDates []string
	for _, index := range datesData.Index {
		if index.ID == artist.ID {
			artistConcertDates = index.Dates
			break
		}
	}

	testimage := canvas.NewImageFromResource(load)
	testimage.FillMode = canvas.ImageFillOriginal

	artistDetailsLabel := widget.NewLabel(fmt.Sprintf("Nom: %s\nMembres: %s\nDate de création: %d\nPremier album: %s\nRelations: %s", artist.Name, strings.Join(artist.Members, ", "), artist.CreationDate, artist.FirstAlbum, artist.Relations))
	locationetdate := fmt.Sprintf("Lieux des concerts: %s\nDates de concert: %s", strings.Join(artistLocations, ", "), strings.Join(artistConcertDates, ", "))
	locationetdatelabel := widget.NewLabel(locationetdate)

	homeButton := widget.NewButton("Home", func() {
		pageglobalartist(mainpage)
	})
	mainpage.SetContent(container.NewVScroll(container.NewVBox(testimage, artistDetailsLabel, locationetdatelabel, homeButton)))
}
