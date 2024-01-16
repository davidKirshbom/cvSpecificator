package handlers

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
) 
 
var CvFormats = map[string]string{
	"PDF": "PDF",
}

type CreateCvOptions struct {
	Format string
}

func getSeleniumDriver() (selenium.WebDriver,error){
	_, err := selenium.NewChromeDriverService("C:\\Program Files\\chromedriver.exe", 4444)

	if err != nil {
		log.Fatal("Error:", err)
		return nil,err
	   }
	   caps := selenium.Capabilities{}
	   caps.AddChrome(chrome.Capabilities{Args: []string{
	   "--window-size=800,1200",
		  "--disable-infobars",
		  "--headless",

	   },
	  })
	driver, err := selenium.NewRemote(caps, "")
	
 	if err != nil {
	log.Fatal("Error:", err)
	return nil,err
	}
   return driver,nil
}
func createScreenshot(driver selenium.WebDriver,savePath string)error{
	screenshotBytes, _ := driver.Screenshot()
	img, _, _ := image.Decode(bytes.NewReader(screenshotBytes))
	out, err := os.Create(savePath)
	if err != nil {
	  log.Fatal("Error on create writer for screenshot png:", err)
	  return err
	}
	err = png.Encode(out, img)
	return err
}

func getPdfWebPageScreenshot(driver selenium.WebDriver) (string,error) {
	screenShootFilePath := "./screenshot" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"
	createScreenshot(driver, screenShootFilePath)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.ImageOptions(
		screenShootFilePath,
		0, 0,
		0, 0,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0, "",
	)
	pdfFilePath := "./cv" + strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
	err := pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		log.Fatal("Error on create pdf file:", err)
		return "",err
	}
	return pdfFilePath,nil
}

func CreateCv(options CreateCvOptions)(string,error) {

	driver,err := getSeleniumDriver()
	if err != nil {
		log.Fatal("Error:", err)
		return  "",err
	}
	defer driver.Quit()
 err = driver.Get("http://localhost:8080/cv/")

 if err != nil {
  log.Fatal("Error on get cv web page:", err)
 }
cvPdfPath,err := getPdfWebPageScreenshot(driver)
if(err != nil){
	log.Fatal("Error on get pdf web page screenshot:", err)
	return "",err
}

 return cvPdfPath,nil
}
