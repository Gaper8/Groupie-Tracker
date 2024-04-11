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

// Here we launch the user interface of the application.
func Graphique() {
	newApp := app.New()
	// We create a new Fyne application.
	windows := newApp.NewWindow("Groupie Tracker !")

	windows.Resize(fyne.NewSize(1200, 800))
	pageglobalartist(windows)
	windows.ShowAndRun()
}

// Then, we create a new window, I define its size. I define the content of my global page.
func pageglobalartist(mainpage fyne.Window) {
	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Entrez votre recherche")

	searchAndFiltersContainer := container.NewVBox(searchBar)
	artistsListScroll := container.NewVScroll(container.NewVBox())
	artistsListContainer, ok := artistsListScroll.Content.(*fyne.Container)
	if !ok {
		fmt.Println("Erreur")
		return
	}
	searchBar.OnChanged = func(query string) {
		filteredArtists := filterArtists(query)
		updateArtistsList(mainpage, filteredArtists, searchAndFiltersContainer, artistsListContainer, artistsListScroll)
	}

	checkbox1 := widget.NewCheck("1", func(b bool) { fmt.Println("Checkbox 1:", b) })
	checkbox2 := widget.NewCheck("2", func(b bool) { fmt.Println("Checkbox 2:", b) })
	checkbox3 := widget.NewCheck("3", func(b bool) { fmt.Println("Checkbox 3:", b) })
	checkbox4 := widget.NewCheck("4", func(b bool) { fmt.Println("Checkbox 4:", b) })
	checkbox5 := widget.NewCheck("5", func(b bool) { fmt.Println("Checkbox 5:", b) })
	checkbox6 := widget.NewCheck("6", func(b bool) { fmt.Println("Checkbox 6:", b) })
	checkbox7 := widget.NewCheck("7", func(b bool) { fmt.Println("Checkbox 7:", b) })

	minValueSlider := widget.NewSlider(0, 50)
	sliderContainer := container.NewStack(minValueSlider)
	sliderContainer.Resize(fyne.NewSize(100, 100))

	maxValueSlider := widget.NewSlider(51, 100)
	maxSliderContainer := container.NewStack(maxValueSlider)
	maxSliderContainer.Resize(fyne.NewSize(100, 100))

	minValueContainer := container.NewHBox(widget.NewLabel("Min"), maxValueSlider, widget.NewLabel("Max"))
	maxValueContainer := container.NewHBox(widget.NewLabel("Min"), maxValueSlider, widget.NewLabel("Max"))

	applyButton := widget.NewButton("Appliquer", func() {
	})

	searchAndFiltersContainer = container.NewVBox(
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

	searchBar.OnChanged("")

	artists, err := Api()
	if err != nil {
		fmt.Println("Erreur", err)
		return
	}

	updateArtistsList(mainpage, artists, searchAndFiltersContainer, artistsListContainer, artistsListScroll)
}

func filterArtists(query string) []ArtisteElement {
	artists, err := Api()
	if err != nil {
		fmt.Println("Erreur", err)
		return nil
	}

	query = strings.ToLower(query)

	var filteredArtists []ArtisteElement
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) ||
			strings.Contains(strings.ToLower(strings.Join(artist.Members, " ")), query) ||
			strings.Contains(strings.ToLower(artist.Locations), query) ||
			strings.Contains(strings.ToLower(artist.FirstAlbum), query) ||
			strings.Contains(strings.ToLower(fmt.Sprint(artist.CreationDate)), query) {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

func updateArtistsList(mainpage fyne.Window, artists []ArtisteElement, searchAndFiltersContainer *fyne.Container, artistsListContainer *fyne.Container, artistsListScroll *container.Scroll) {
	listButtonArtist := make([]fyne.CanvasObject, 0)
	for _, artist := range artists {
		artist := artist
		button := widget.NewButton(artist.Name, func() func() {
			return func() {
				showdataartist(mainpage, artist)
			}
		}())
		listButtonArtist = append(listButtonArtist, button)
	}

	artistsListContainer.Objects = nil
	for _, button := range listButtonArtist {
		artistsListContainer.Add(button)
	}
	artistsListContainer.Refresh()
	mainpage.SetContent(container.NewBorder(nil, nil, searchAndFiltersContainer, nil, artistsListScroll))
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
