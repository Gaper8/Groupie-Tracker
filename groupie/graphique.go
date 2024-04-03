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
	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Entrez votre recherche")

	checkbox1 := widget.NewCheck("1", func(b bool) { fmt.Println("Checkbox 1:", b) })
	checkbox2 := widget.NewCheck("2", func(b bool) { fmt.Println("Checkbox 2:", b) })
	checkbox3 := widget.NewCheck("3", func(b bool) { fmt.Println("Checkbox 3:", b) })
	checkbox4 := widget.NewCheck("4", func(b bool) { fmt.Println("Checkbox 4:", b) })
	checkbox5 := widget.NewCheck("5", func(b bool) { fmt.Println("Checkbox 5:", b) })
	checkbox6 := widget.NewCheck("6", func(b bool) { fmt.Println("Checkbox 6:", b) })
	checkbox7 := widget.NewCheck("7", func(b bool) { fmt.Println("Checkbox 7:", b) })

	minValueSlider := widget.NewSlider(0, 50)
	sliderContainer := container.NewMax(minValueSlider)
    sliderContainer.Resize(fyne.NewSize(100, 100)) 

	maxValueSlider := widget.NewSlider(51, 100)
	maxSliderContainer := container.NewMax(maxValueSlider)
	maxSliderContainer.Resize(fyne.NewSize(100, 100))
	

	minValueContainer := container.NewHBox(widget.NewLabel("Min"), maxValueSlider, widget.NewLabel("Max"))
	maxValueContainer := container.NewHBox(widget.NewLabel("Min"), maxValueSlider, widget.NewLabel("Max"))
	
	

	applyButton := widget.NewButton("Appliquer", func() {
	})

	searchAndFiltersContainer := container.NewVBox(
		searchBar,
		widget.NewLabel("Nombre de membres :"),
		checkbox1, checkbox2, checkbox3, checkbox4, checkbox5, checkbox6, checkbox7,
		widget.NewLabel("Date du premier album :"),
		minValueSlider,
		minValueContainer,
		
		widget.NewLabel("Date de Création :"),
		maxValueSlider,
		maxValueContainer,
		

		applyButton,
	)

	artists, err := Api()
	if err != nil {
		fmt.Println("Erreur", err)
		return
	}

	listButtonArtist := make([]fyne.CanvasObject, 0, len(artists))
	for _, art2 := range artists {
		art := art2
		button := widget.NewButton(art.Name, func() {
			showdataartist(mainpage, art)
		})
		listButtonArtist = append(listButtonArtist, button)
	}

	artistsListContainer := container.NewVScroll(container.NewVBox(listButtonArtist...))

	mainpage.SetContent(container.NewBorder(nil, nil, searchAndFiltersContainer, nil, artistsListContainer))
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
