package router

import (
	"fmt"
	"log"
	"os"
	

	"github.com/bwmarrin/discordgo"
	"github.com/davidKirshbom/cvSpecificator/x/mux"
)

var Router = mux.New()
var Session *discordgo.Session

func Init(){
	

	Session,_=discordgo.New("Bot "+ os.Getenv("DISCORD_BOT_TOKEN"));
	if Session==nil{ 
		log.Fatalf("Error creating discord session")
	}
	
	if Session.Token==""{
		Session.Token=os.Getenv("DISCORD_BOT_TOKEN");
	} 
	

	fmt.Println("Bot is running")
	Session.AddHandler(Router.OnMessageCreate)

	Router.Route("help","Display this message",Router.Help)
	Router.Route("fill","fill cv data",Router.GetUserDetails)
	Router.Route("optimize","optimize cv",Router.Optimize)
	err :=Session.Open();
	if err!=nil{
		fmt.Println("Error opening connection")
		os.Exit(1)
		return
	}
}