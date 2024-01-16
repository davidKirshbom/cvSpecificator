package main

import (
	"context"
	"fmt"

	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/davidKirshbom/cvSpecificator/handlers"
	"github.com/davidKirshbom/cvSpecificator/models"
	router "github.com/davidKirshbom/cvSpecificator/router"
	"github.com/davidKirshbom/cvSpecificator/utils"
	"github.com/joho/godotenv"
)
func init(){
	dir,_ :=os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s\\%s",dir ,"..\\.env") ,);
	log.Print(os.Getenv("DISCORD_BOT_TOKEN"))
	
	if err != nil {
		log.Fatalf("Error loading .env file")
	  }
	  log.Printf("env file loaded")

	

}
func main()  {

	//create the discordgo session
	router.Init()
	sc:=make(chan os.Signal,1)

	//create cv server
	mux := http.NewServeMux()
	mux.HandleFunc("/cv/", cvHandler)
	workingDir ,_ := os.Getwd()
	cssDir :=http.Dir(workingDir+"\\resources\\static\\.");

	log.Print(cssDir)
	mux.Handle("/css/",http.StripPrefix("/css/", http.FileServer(cssDir)));
    server := &http.Server{Addr: ":8080", Handler: mux}
	defer server.Shutdown( context.Background());
	
	if err := server.ListenAndServe(); err != nil {
        log.Fatal(err) 
    }

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc 
	
	
} 
func cvHandler(w http.ResponseWriter, r *http.Request) {
	workingsDir, _ := os.Getwd()
	templatePath :=workingsDir+"\\resources\\templates";
	template , err:=template.New("cv.html").Funcs(template.FuncMap{
		"Titleize": utils.Titleize,
		"NewLineToBr": utils.NewLineToBr,
		"WithComData": utils.WithComData,
		"Htmlize":utils.Htmlize,
		"HtmlEmphasisWords":utils.HtmlEmphasisWords,
	}).ParseFiles(templatePath+"\\cv.html",templatePath+"\\experience_block.html",templatePath+"\\education_block.html")
	if err!=nil{
		log.Print(err)
	}
	r.ParseForm() 
	
	var user models.User
	handlers.HandleLoadUserData(&user)
	 err=template.Execute(w,&user) ;if err!=nil{
		 log.Print(err)
	 }
	
}	