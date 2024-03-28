package groupie

import (
    "fmt"
	"fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "strconv"
)

var suggestions []string // Déclarer la variable suggestions 

func Recherche() {

    //check box widget
    checkbox1 := widget.NewCheck("1", func(b bool) { fmt.Println("Checkbox 1:", b) })
    checkbox2 := widget.NewCheck("2", func(b bool) { fmt.Println("Checkbox 2:", b) })
    checkbox3 := widget.NewCheck("3", func(b bool) { fmt.Println("Checkbox 3:", b) })
    checkbox4 := widget.NewCheck("4", func(b bool) { fmt.Println("Checkbox 4:", b) })
    checkbox5 := widget.NewCheck("5", func(b bool) { fmt.Println("Checkbox 5:", b) })
    checkbox6 := widget.NewCheck("6", func(b bool) { fmt.Println("Checkbox 6:", b) })
    checkbox7 := widget.NewCheck("7", func(b bool) { fmt.Println("Checkbox 7:", b) })

    //un widget pour la barre de recherche
    searchBar := widget.NewEntry()
    searchBar.SetPlaceHolder("Entrez votre recherche")

    //un widget pour afficher les suggestions  
    suggestionList := widget.NewList(
        func() int {
            return len(suggestions)// Retourne a longueur des suggestions
        },
        func() fyne.CanvasObject {
            return widget.NewLabel("") //un label vide pour chaque élément de la liste
        },
		func(i int, obj fyne.CanvasObject) {
            // Mettre à jour le label avec le texte de la suggestion
            obj.(*widget.Label).SetText(suggestions[i])
        },
		
    )

    //Les valeurs de la plage
    minValueEntry := widget.NewEntry()
	maxValueEntry := widget.NewEntry()
    minValueEntry1 := widget.NewEntry()
	maxValueEntry2 := widget.NewEntry()

    // Bouton pour appliquer le filtre
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

    //un conteneur pour organiser les widgets
    container.NewVBox(
        searchBar,
        suggestionList,
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
    )
}