package mux

import (	
	 "github.com/davidKirshbom/cvSpecificator/models"
	"github.com/davidKirshbom/cvSpecificator/handlers"
	 "github.com/bwmarrin/discordgo"
) 
var user models.User
var userQuestions []models.Question
var lastQuestionIndex = 0

func init(){
	user =models.User{}
	userQuestions =  handlers.CreateQuestionsGetDetails(user)
	lastQuestionIndex =0
}


func (m *Mux) GetUserDetails(session *discordgo.Session, userMessage *discordgo.Message, context *Context) {
	isResetMessage := context.IsDirected 
	if(isResetMessage){
		user.Reset()
		m.OpenDialog("display this message",m.GetUserDetails,context,userMessage)
		
	}else{ 
		var lastQuestion = userQuestions[lastQuestionIndex]
		lastQuestion.Answer(userMessage.Content)
		if lastQuestion.AnswerAction!=nil{
			lastQuestion.AnswerAction(userMessage.Content)
		}
		if lastQuestion.After!=nil{
			lastQuestion.After()
		}
	}

	for index,question := range userQuestions{	
		if question.Check(){
			if(!isResetMessage){lastQuestionIndex = index}
			if question.QuestionFunc!=nil{
				question.Question = question.QuestionFunc() 
				
			}
			_, err := session.ChannelMessageSend(userMessage.ChannelID, question.Question)
			if err != nil {
				// Handle error
			}
			return
		}
	}
	if !user.State.StartedForm{
		_, err := session.ChannelMessageSend(userMessage.ChannelID, "hey let's not waste time, what's your first name?")
		if err != nil {
			// Handle error
		}
		user.State.StartedForm=true
	}
	
}
