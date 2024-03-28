package groupie

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
	artistDetailsLabel := widget.NewLabel(fmt.Sprintf("Nom: %s\nMembres: %s\nDate de création: %d\nPremier album: %s\nLieux: %s\nDates de concert: %s\nRelations: %s", artist.Name, strings.Join(artist.Members, ", "), artist.CreationDate, artist.FirstAlbum, artist.Locations, artist.ConcertDates, artist.Relations))

	homeButton := widget.NewButton("Home", func() {
		pageglobalartist(mainpage)
	})
	mainpage.SetContent(container.NewVScroll(container.NewVBox(artistDetailsLabel, homeButton)))
}
