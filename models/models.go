package models

import (
	"encoding/base64"
	"html/template"

	"github.com/davidKirshbom/cvSpecificator/utils"
	"github.com/skip2/go-qrcode"
)
type Role struct {
	Title   string `json:"title"` // Explain string
	Explain string `json:"explain"`
}
func (r *Role) Clone()Role{
	return Role{
		Title: r.Title, 
		Explain: r.Explain,
	}
}

type Experience struct{
	StartYear string 		`json:"startYear"`
	EndYear string 			`json:"endYear"`
	CompanyName string 		`json:"companyName"`
	Location string 		`json:"location"`
	Role Role				`json:"role"`
}

func (e *Experience) Clone()Experience{
	return Experience{
		StartYear: e.StartYear,
		EndYear: e.EndYear,
		CompanyName: e.CompanyName,
		Location: e.Location,
		Role:e.Role.Clone(),
	}
}
func (e *Experience) HTML() template.HTML{
	return  template.HTML(utils.Htmlize(e.Role.Explain))
}

type Education struct{
	StartYear string `json:"startYear"`
	EndYear string `json:"endYear"`
	SchoolName string `json:"schoolName"`
	Location string `json:"location"`
	DegreeTitle string `json:"degreeTitle"`
	Learned string `json:"learned"`
} 

type Link struct{
	Label string `json:"label"`
	Url string `json:"url"`

}
func (l *Link) QR() string {
	result, err := qrcode.Encode(l.Url, qrcode.Medium, 200)
	if err != nil {
		panic(err)
	}

	return  (base64.StdEncoding.EncodeToString(result)) 
}



type State struct{
	
	StartedForm bool	`default:"false" json:"startedForm"`
	MoreExperienceAsked bool `json:"moreExperienceAsked" default:"false"`
	FinishExperienceAsked bool `default:"false" json:"finishExperienceAsked"`
	FinishForm bool `default:"false" json:"finishForm"`
	MoreEducationAsked bool `default:"false" json:"moreEducationAsked"`
	FinishEducationAsked bool `default:"false" json:"finishEducationAsked"`
	

}

type Question struct{
	Check func()bool
	Question string	
	QuestionFunc func()string
	Answer func(answer string) 
	AnswerAction func(answer string)
	After func()
}