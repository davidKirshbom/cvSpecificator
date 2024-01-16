package mux

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/davidKirshbom/cvSpecificator/handlers"
) 
func init(){
	
	userQuestions =  handlers.CreateQuestionsGetDetails(user)
	lastQuestionIndex =0
}

func (m *Mux) Optimize(session *discordgo.Session, userMessage *discordgo.Message, context *Context) {
	if(!user.State.FinishForm){ 
		err := user.Load()
		if err!=nil{
		session.ChannelMessageSend(userMessage.ChannelID, "Don't have your data please fill it first by typing 'fill'")
		return
		}
	
	}
	jobDescription := userMessage.Content
	// Remove the first two words from job description removing bot id and command
	words := strings.Fields(jobDescription)
	if len(words) > 2 {
		jobDescription = strings.Join(words[2:], " ")
		jobDescription = strings.ReplaceAll(jobDescription, "\n\n", "")

	} else {
		jobDescription = ""
	}

	user.OptimizeExperience =  handlers.HandleGetChatgptOptimization(user.Experience,jobDescription)
	user.Save()
	pdfFilePath ,err:=handlers.CreateCv(handlers.CreateCvOptions{Format: handlers.CvFormats["PDF"]})
	if err!=nil{
		session.ChannelMessageSend(userMessage.ChannelID, "Error creating cv")
		return
	}
	pdfReader, err := os.Open(pdfFilePath)
	session.ChannelFileSend(userMessage.ChannelID, "cv.pdf",pdfReader)



}

