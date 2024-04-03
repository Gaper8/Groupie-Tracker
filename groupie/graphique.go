package groupie

import (
	"fmt"
	"strings"
	"strconv"

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
	// Barre de recherche
	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Entrez votre recherche")

	//check box widget
    checkbox1 := widget.NewCheck("1", func(b bool) { fmt.Println("Checkbox 1:", b) })
    checkbox2 := widget.NewCheck("2", func(b bool) { fmt.Println("Checkbox 2:", b) })
    checkbox3 := widget.NewCheck("3", func(b bool) { fmt.Println("Checkbox 3:", b) })
    checkbox4 := widget.NewCheck("4", func(b bool) { fmt.Println("Checkbox 4:", b) })
    checkbox5 := widget.NewCheck("5", func(b bool) { fmt.Println("Checkbox 5:", b) })
    checkbox6 := widget.NewCheck("6", func(b bool) { fmt.Println("Checkbox 6:", b) })
    checkbox7 := widget.NewCheck("7", func(b bool) { fmt.Println("Checkbox 7:", b) })

	minValueEntry := widget.NewEntry()
	maxValueEntry := widget.NewEntry()
	minValueEntry1 := widget.NewEntry()
	maxValueEntry2 := widget.NewEntry()

	applyButton := widget.NewButton("Appliquer", func() {
		minValueStr := minValueEntry.Text
		maxValueStr := maxValueEntry.Text

		minValue, err := strconv.ParseFloat(minValueStr, 4)
		if err != nil {
			fmt.Println("Erreur lors de la conversion de la valeur minimale:", err)
			return
		}

		maxValue, err := strconv.ParseFloat(maxValueStr, 4)
		if err != nil {
			fmt.Println("Erreur lors de la conversion de la valeur maximale:", err)
			return
		}

		fmt.Printf("Filtrer les résultats entre %.0f et %.0f\n", minValue, maxValue)
	})
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
		button.Resize(fyne.NewSize(200, 50))
	}
	listContainer := container.NewVBox(listbuttonartist...)
	scrollableList := container.NewVScroll(container.NewVBox(listbuttonartist...))
	scrollableList.Resize(fyne.NewSize(1200, 600))
	mainpage.SetContent(scrollableList)

	mainpage.SetContent(container.NewVBox(
		searchBar,
		widget.NewLabel("Nombre de membres :"),
		checkbox1,
        checkbox2,
        checkbox3,
        checkbox4,
        checkbox5,
        checkbox6,
        checkbox7,
		widget.NewLabel("Date du premier album :"),
		container.NewGridWithColumns(2,
			widget.NewLabel("Valeur min"),
			minValueEntry,
			widget.NewLabel("Valeur max"),
			maxValueEntry,
		),
		widget.NewLabel("Date de Creation :"),
		container.NewGridWithColumns(2,
			widget.NewLabel("Valeur mine"),
			minValueEntry1,
			widget.NewLabel("Valeur max"),
			maxValueEntry2,
			widget.NewLabel(""),
			applyButton,
		),
		scrollableList,
		listContainer,
	))
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

	relationData, err := RelationApi()
	if err != nil {
		fmt.Println("Erreur", err)
		return
	}

	var artistRelations []string
	for _, index := range relationData.Index {
		if index.ID == artist.ID {
			for date, locations := range index.DatesLocations {
				artistRelations = append(artistRelations, fmt.Sprintf("%s: %s", date, strings.Join(locations, ", ")))
			}
			break
		}
	}

	testimage := canvas.NewImageFromResource(load)
	testimage.FillMode = canvas.ImageFillOriginal

	artistDetailsLabel := widget.NewLabel(fmt.Sprintf("Nom: %s\nMembres: %s\nDate de création: %d\nPremier album: %s", artist.Name, strings.Join(artist.Members, ", "), artist.CreationDate, artist.FirstAlbum))
	locationetdate := fmt.Sprintf("Lieux des concerts: %s\nDates de concert: %s", strings.Join(artistLocations, ", "), strings.Join(artistConcertDates, ", "))
	locationetdatelabel := widget.NewLabel(locationetdate)

	relationsLabel := widget.NewLabel("Dernières relations:")
	relationsDataLabel := widget.NewLabel(strings.Join(artistRelations, "\n"))

	homeButton := widget.NewButton("Home", func() {
		pageglobalartist(mainpage)
	})
	mainpage.SetContent(container.NewVScroll(container.NewVBox(testimage, artistDetailsLabel, locationetdatelabel, relationsLabel, relationsDataLabel, homeButton)))
}
