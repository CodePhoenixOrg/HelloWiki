package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	title string
	body []byte
}

func CreatePage() *Page {
	return &Page{title: "", body: []byte("")}
}

func NewPage() *Page {
	return new(Page)
}

func (p *Page) Title() string {
	return p.title
}

func (p *Page) SetTitle(title string) {
	p.title = title
}

//func (p *Page) Body() []byte {
//	return p.body
//}

//func (p *Page) SetBody(body []byte) {
//	p.body = body
//}

func (p *Page) save() error {
	filename := p.title + ".txt"

	return ioutil.WriteFile(filename, p.body, 0600)
}

func (p *Page) clear() {
	p.title = ""
	p.body = []byte("")
}

//func (p *Page) load(title string) (*Page, error) {
func (p *Page) load(title string) error {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	p.title = title
	p.body = body

	//return &Page{title: title, body: body}, nil
	return nil
}

func (p *Page) viewHandler (w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p.load(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.title, p.body)
	fmt.Println(string(p.body))
}

func main() {
	fmt.Println("Create a new page")

	p1 := NewPage()

//	p1.SetTitle("TestPage")
//	p1.SetBody([]byte("This is a sample Page.\n"))

	p1.title = "TestPage"
	p1.body = []byte("This is a sample Page.\n")

	fmt.Println("Save the page")
	p1.save()

	fmt.Println("Clear the object")
	p1.clear()

	p2 := CreatePage()
	p2.title = "Dummy"

	//p2, _ := p1.load("TestPage")
	fmt.Println("Load the page")
	//p2.load("TestPage")

	//p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.body))

	http.HandleFunc("/view/", p2.viewHandler)
	http.ListenAndServe(":8000", nil)

}


