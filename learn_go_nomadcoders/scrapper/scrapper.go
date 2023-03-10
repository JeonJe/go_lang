package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct{
	id string
	title string
	location string 
	summary string
}

//Scrape 
func Scrape(term string){ 
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=" + term 
	var jobs []extractedJob
	c := make (chan []extractedJob)

	totalPages := getPages(baseURL)
	fmt.Println("Total pages: ", totalPages)

	for i := 0; i< totalPages; i++ {
		go getPage(i, baseURL, c)
	}
	for i := 0; i< totalPages; i++ {
		extractedJobs := <-c 
		jobs = append(jobs, extractedJobs...)
	}
	
	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}


func getPage(page int, baseURL string, mainC chan<- []extractedJob){
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("requesting : " + pageURL)

    res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close();

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs

}


func writeJobs(jobs []extractedJob){
	file,err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID","Title","Location","Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx="+ job.id, job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}

} 

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := CleanString(card.Find(".job_tit>a").Text())
	location := CleanString(card.Find(".job_condition>span>a").Text())
	summary := CleanString(card.Find(".job_sector").Text())
	c<- extractedJob{
		id: id,
        title: title,
        location: location,
        summary: summary,
    }
}
//Clean String cleans the string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str))," ")
}

func getPages(baseURL string) int{
	pages := 0
	res, err := http.Get(baseURL)

	checkErr(err)
	checkCode(res)
	
	defer res.Body.Close();

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}


func checkErr(err error){
	if err != nil{
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response){
	if res.StatusCode != 200{
		log.Fatalln("request status : ",res.Status)
    }
	
}