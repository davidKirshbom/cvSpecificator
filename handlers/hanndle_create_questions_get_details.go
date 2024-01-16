package handlers

import (
"fmt"

"github.com/davidKirshbom/cvSpecificator/models"
"log"
	"strings"
)

	func CreateQuestionsGetDetails(user models.User) []models.Question {

		var questions=[]models.Question {
	{	
		Check: func() bool {
			return user.FirstName==""	
		},
		Question:"hey let's not waste time, what's your first name?",	
		Answer: func(answer string){
			user.FirstName = answer
		},
		
		
	},
	{
		Check: func() bool {
			return user.LastName==""	
		},
		Question:fmt.Sprintf("hey %s, what's your last name?",user.FirstName),	
		Answer: func(answer string){
			user.LastName = answer
			},
	},
	{	
		Check: func() bool {
			return user.Title==""	
		},
		Question:fmt.Sprintf("hey %s %s, what's your role title?",user.FirstName,user.LastName),
		Answer: func(answer string){
			user.Title = answer
			},
	},
	{
		Check: func() bool {
			return user.About==""	
		},
		Question:"what's your about?",
		Answer:	func(answer string){
			user.About = answer
			},
		After:func(){
			user.Experience = append(user.Experience,models.Experience{})
		},
	},
	{
		Check: func() bool {
			return user.Experience[len(user.Experience)-1].CompanyName == ""
		},
		QuestionFunc: func() string {
			if len(user.Experience) == 1 {
				return "tell me about the first relevant job you had, what was the company name?"
			} else {
				return "tell me more, what was the company name?"
			}
		},
		Answer: func(answer string) {
			user.Experience[len(user.Experience)-1].CompanyName = answer
		},
	},
	{
		Check:func() bool{
			return user.Experience[len(user.Experience)-1].Location==""
		},
		Question:"what was the location of the company?",
		Answer:func (answer string){
			user.Experience[len(user.Experience)-1].Location = answer
		},
	},
	{
		Check:func()bool{
			return user.Experience[len(user.Experience)-1].Role.Title==""
		},

		Question:"what was your role title?",
		Answer:func (answer string){
			user.Experience[len(user.Experience)-1].Role.Title = answer
		},
	},
	{
		Check:func()bool{
			return user.Experience[len(user.Experience)-1].Role.Explain==""
		},
		Question:"Explain your role in a few words",
		Answer:func (answer string){
			user.Experience[len(user.Experience)-1].Role.Explain = answer
		},
	},
	{
		Check:func()bool{
			return user.Experience[len(user.Experience)-1].StartYear==""
		},
		Question:"what year did you start working there?",
		Answer:func (answer string){
			user.Experience[len(user.Experience)-1].StartYear = answer
		},
	},
	{
		Check:func()bool{
			return user.Experience[len(user.Experience)-1].EndYear==""
		},
		Question:"what year did you end working there?",	
		Answer:func (answer string){
			user.Experience[len(user.Experience)-1].EndYear = answer
			},
	},

	{	
		Check:func()bool{
			return !user.State.FinishExperienceAsked
		},
		Question:"do you have another relevant job experience?(yes/no)",	
		Answer:func(answer string){
			if(answer=="yes"){
				user.Experience = append(user.Experience,models.Experience{})
				user.State.MoreExperienceAsked=false
			}
			if(answer=="no"){
				user.State.MoreExperienceAsked=false
				user.State.FinishExperienceAsked=true
				user.Education = append(user.Education,models.Education{})
			}
		},	
	},
	{
		Check:func()bool{
			return user.Education[len(user.Education)-1].SchoolName==""
		},
			QuestionFunc: func() string {
				if len(user.Education)==1{
					return "tell me about the first relevant Education you had,what was the school name?"
				}else
					{return "tell me more,what was the school name?"}
			},
			Answer: func (answer string){
			user.Education[len(user.Education)-1].SchoolName = answer
		},
	},
	{
		Check:func()bool{
			return user.Education[len(user.Education)-1].Location==""
		},
		Question:"what was the location of the school?",
		Answer:func (answer string){
			user.Education[len(user.Education)-1].Location = answer
		},
	},
	{
		Check:func()bool{
			return user.Education[len(user.Education)-1].DegreeTitle==""
		},
		Question:"what was your degree title?",
		Answer:func (answer string){
			user.Education[len(user.Education)-1].DegreeTitle = answer
		},
	},
	{
		Check:func()bool{
			return user.Education[len(user.Education)-1].Learned==""
		},
		Question:"what did you learn?",
		Answer:func (answer string){
			user.Education[len(user.Education)-1].Learned = answer
			},
	},
	{
		Check:func()bool{
			return user.Education[len(user.Education)-1].StartYear==""
		},
		Question:"what year did you start learning there?",
		Answer:func (answer string){
			user.Education[len(user.Education)-1].StartYear = answer
			},
	},
	{
		Check:func()bool{
			return user.Education[len(user.Education)-1].EndYear==""
		},
		Question:"what year did you end learning there?",
		Answer:func (answer string){
			user.Education[len(user.Education)-1].EndYear = answer
			},
	},
	{
		Check:func()bool{
			return !user.State.FinishEducationAsked
		},
		Question:"do you have another relevant education experience?",
		Answer:func(answer string){
			if answer=="yes"{
				user.Education = append(user.Education,models.Education{})
				user.State.MoreEducationAsked=true

			}
			if answer=="no"{
				user.State.FinishEducationAsked=true
			}

		},
	},
	
	{
		Check:func()bool{
			return len(user.Skills)==0 
		},
		Question:"what are your skills?(nodejs,react,python,etc) (separate with commas)",
		Answer:func(answer string) {
			user.Skills = strings.Split(answer,",")
		},
	},
	{
		Check:func()bool{
			return len(user.Languages)==0
		},
		Question:"what are your languages?(english,hebrew,etc) (separate with commas)",
		Answer:func(answer string) {
			user.Languages = strings.Split(answer,",")
		},
	},
	{
		Check:func()bool{
			return user.Link.Label==""
		},
		Question:"what is your link label?",
		Answer: func(answer string) {
			user.Link.Label = answer
			},
	},
	{
		Check:func()bool{
			return user.Link.Url==""
		},
		Question:"what is your link url?",
		Answer: func(answer string) {
			user.Link.Url = answer
			user.State.FinishForm=true
		},
	},
	{	
		Check: func() bool {
			return user.State.FinishForm	
		},
		QuestionFunc:func() string {
			return "great, i have your data,next time you see a job post just throw me the description and il make a good cv for that job"
		},
		Answer: func(answer string) {
			user.State.FinishForm=true
			err :=user.Save()
			
			if err!=nil{
				log.Println(err)
				
			}
			},
		After: func() {
			user.State.FinishForm=true
			
			},
	},
	
}
return questions


	}