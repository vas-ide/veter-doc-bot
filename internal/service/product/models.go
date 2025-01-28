package product

import "github.com/google/uuid"


var allProducts = []Product{
	{Title: "Call a doctor"},
	{Title: "Examination"},
	{Title: "Recommendations and treatment"},
	{Title: "Sterilization"},
	{Title: "Consultation"},

}



type Product struct {
	
	
	Title string

	
}


type Animal struct {

	ID uuid.UUID
	Owner Owner
	Name string
	

}


type Owner struct {

	ID uuid.UUID
	Name string

}


















