package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Project struct {
	Project       string
	Company       string
	Role          string
	ProjectPeriod string
}

type Contact struct {
	Name, Email, Phone string
	WillAttend         bool
}

var responces = make([]*Contact, 0, 10)
var templates = make(map[string]*template.Template, 4)

func LoadTemplates() {
	templateNames := [4]string{"welcome", "form", "list", "thanks"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	templates["welcome"].Execute(w, nil)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	templates["list"].Execute(w, Projects)
}

type formData struct {
	*Contact
	Errors []string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates["form"].Execute(w, formData{
			Contact: &Contact{}, Errors: []string{},
		})

	} else if r.Method == http.MethodPost {
		r.ParseForm()
		responceData := Contact{
			Name:  r.FormValue("name"),
			Email: r.FormValue("email"),
			Phone: r.FormValue("phone"),
		}

		errors := []string{}
		// Проверка значений полей формы. Пустые поля недопустимы.
		if responceData.Name == "" {
			errors = append(errors, "Please, enter your name!")
		}
		if responceData.Email == "" {
			errors = append(errors, "Please, enter your email!")
		}
		if responceData.Phone == "" {
			errors = append(errors, "Please, enter your phone!")
		}
		if len(errors) > 0 {
			templates["form"].Execute(w, formData{
				Contact: &responceData, Errors: errors,
			})
		} else {
			responces = append(responces, &responceData)
			templates["thanks"].Execute(w, responceData.Name)
		}
	}

}

var Projects = []Project{
	{
		Project:       "Identity Threat Detection and Response",
		Company:       "Компания Индид",
		Role:          "QA Tech Lead",
		ProjectPeriod: "2023-По настоящее время",
	},
	{
		Project:       "Гипервизор Горизонт-ВС",
		Company:       "ИЦ Баррикады",
		Role:          "Head of QA",
		ProjectPeriod: "2022-2023",
	},
	{
		Project:       "MSP360 Backup for macOS and Linux",
		Company:       "MSP360",
		Role:          "QA Engineer/QA Lead",
		ProjectPeriod: "2019-2022",
	},
	{
		Project:       "СЗИ ВИ Dallas Lock",
		Company:       "Конфидент",
		Role:          "QA Engineer",
		ProjectPeriod: "2018-2019",
	},
}

func main() {
	LoadTemplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
