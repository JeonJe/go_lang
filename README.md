d
<details>
<summary>Nomadcoders go lang lecture</summary>
<div markdown="1">       

-------------------
**<링크>** : [쉽고 빠른 Go 시작하기 - 노마드 코더 Nomad Coders](https://nomadcoders.co/go-for-beginners?gclid=CjwKCAiAzp6eBhByEiwA_gGq5KXGdhYB3qXCoJm-uujH6DD1fd-yMnGJrk9fZsk07_FGMFDX8GJUHxoCvScQAvD_BwE)

- python 처럼 원하는 디렉토리에 프로젝트를 만들어서 사용할 수 없음. 무조건 `GO path` 디렉토리에 있어야 한다
- Go에서는 내가 원하는 곳 어디서에서든 코드를 다운로드 받아 사용할 수 있다
    - /go/src/에 도메인별로 분류해서 저장해놓는 것이 좋다

        

- 프로젝트 컴파일이 필요하면 파일이름을 main.go로 만들어야 한다
    - main은 entry point여서 컴파일러는 패키지 이름이 main인 것 부터 찾아낸다
    - 사람들이랑 같이 쓸거면 컴파일이 필요없어서 main.go가 없을 것이다
    - go는 특정 function을 찾게 되는데, `func main(){ }` 이다
    - 실행은 `go run main.go [파일명]`

- print는 `import “fmt”` 후 `fmt.Println(" ~~ ")`
- Go는 function을 expert하고 싶으면 function 이름의 첫 글자를 대문자로 작성하면 된다. 그것이 Println의 첫문자가 대문자인 이유다.
- function안에서 var와 타입을 축약해서 사용할 수도 있다. `name := “jeonje”`
    - function 밖에서는 `:=` 은 동작하지 않는다.

- **인자와 리턴의 타입은 반드시 명시 해야한다**
    
    ```go
    func multiply(a int , b int) int {
    	return a * b 
    }
    
    func multiply(a, b int) int {
    	return a * b 
    }
    
    func lenAndUpper(name string) (int, string){
    	return len(name), strings.ToUpper(name)
    }
    
    func main(){ 
    	totalLength, _ := lenAndUpper("nico")
    	fmt.Println(totalLength)
    }
    ```
    
    - return 의 값은 여러개 가능
- Go의 package [golang.org](http://golang.org) 에서 확인 가능하다
- 가변인자 …
    
    ```go
    func repeatMe(words ...string) {
    	fmt.Println(words)
    }
    
    func main(){ 
    	repeatMe("arg","arg2", "arg3")
    }
    ```
    
- naked return
    - return 할 변수를 꼭 명시 하지 않아도 된다
    
    ```go
    func lenAndUpper(name string) (length int, uppercase string){
    	length = len(name)
    	uppercase = strings.ToUpper(name)
    	return
    }
    ```
    
- defer
    - function이 끝났을 때 추가적인 기능을 넣을 수 있다
    - defer는 function이 값을 return 한 뒤에 실행
- for
    
    ```go
    func superAdd(numbers ...int) int{
    	total := 0
    	for _ ,number := range numbers {
    		total += number
    	}
    	return total
    }
    
    func superAdd2(numbers ...int) int{
    	total := 0
    	for i:=0; i< len(numbers); i++{
    		total += numbers[i]
    	}
    	return total
    }
    
    func main(){ 
    	total := superAdd(1,2,3,4,5)
    	fmt.Println(total)
    	
    }
    ```
    
- if 안에다가 변수 선언 가능 (if에서만 사용하기 위해)
    
    ```
    func canIDrink(age int) bool{
    	if koreaAge := age +2; koreaAge < 18{
    		return false
    	} else {
    		return true
    	}
    }
    ```
    
- switch 안에다가 변수 선언 가능 (switch만을 사용하기 위해)
    
    ```go
    func canIDrink(age int) bool{
    	switch koreanAge := age + 2; koreanAge{
    		case 10:
                return false
            case 18:
                return true
    	}
    	return false
    }
    ```
    
- 슬라이스
    - length가 없는 array라 생각하면 된다
        
        ```go
        names := []string{"a", "b", "c", "d", "e","f"}
        ```
        
    - append 함수는 인자가 추가된 새로운 슬라이스를 return한다
    - 대부분 슬라이스를 사용 한다
- Map
    
    ```go
    func main(){ 
    	nico := map[string] string{"name": "nico", "age": "12" }
    	for key, val := range nico{
    		fmt.Println(key, val)
    	}
    }
    ```
    
- struct
    - go는 class나 object가 없다
    - 파이선 처럼 “__**init__” constructor method가 없다 , 스스로 constructor를 실행해야 한다**
    
    ```go
    type person struct {
    	name string
      age  int
    	favFood []string
    }
    
    func main(){ 
    	favFood := []string{"kimchi","ramen"}
    	whssodi := person{name:"whssodi", age:30, favFood:favFood}
    	fmt.Println(whssodi)
    }
    ```
    

- **bank account proj**
    
    constructor를 사용해서 만드는 법 
    
    ```go
    //main.go
    package main
    
    import (
    	"fmt"
    	"github.com/JeonJe/learngo/accounts"
    )
    
    func main(){ 
    	
    	account := accounts.NewAccount("jeonje")
    	fmt.Println(account)
    
    }
    
    //accounts.go
    package accounts
    
    // Account struct
    type Account struct {
    	owner string
    	balance int
    }
    
    // NewAccount creates a new account
    func NewAccount(owner string) *Account {
    	account := Account{owner : owner, balance : 0}
    	return &account 
    }
    ```
    
    - receiver
        - receiver를 작성하는데 있어서 지켜야 할 사항
            - struct의 첫 글자를 따서 소문자로 지어야 한다
        - receiver의 값을 변경하려면 포인터 리시버로 전달해야 한다. 그냥 리시버로 전달하면 복사본에만 반영이 된다.
    
    ```go
    //Deposit x amount on your account
    //복사본의 balance가 증가하므로 main에서 반영되지 않음 
    func (a Account) Deposit(amount int){
    	a.balance += amount
    }
    
    //Deposit x amount on your account
    //반영시키려면 리시버를 *Account로 받아야 함 
    func (a *Account) Deposit(amount int){
    	a.balance += amount
    }
    
    //Balance of your account
    func (a Account) Balance() int{
    	return a.balance
    }
    ```
    
    - 예외처리
        - Go에는 exception 같은 것이 없다.  try - except, try - catch 도 없다
        - nil은 파이썬의 None 같은 느낌이다.
    
    ```go
    //Withdraw x amount on your account
    func (a *Account) WithBalance(amount int) error {
    
    	if a.balance < amount {
    		return errors.New("can't withdraw")
    	}
    
    	a.balance -= amount
    	return nil
    }
    ```
    
    - 리턴형은 error로 써줘야 하며, return error.Error() 또는 errors.New(”에러문구”)를 써야 한다. 에러문구의 첫 시작은 대문자이면 안된다. 에러가 아닐 때는 return nil을 해줘야 함.
    
    ```go
    //main.go은 에러를 확인해서 처리 해줘야 함 
    err := account.WithBalance(20)
    
    	if err != nil{
    		log.Fatalln(err)
    	}
    ```
    
    ```go
    **var errNoMoney = errors.New("can't withdraw")**
    
    //Withdraw x amount on your account
    func (a *Account) WithBalance(amount int) error {
    
    	**if a.balance < amount {
    		return errNoMoney
    	}**
    
    	a.balance -= amount
    	return nil
    }
    ```
    
    - 이렇게도 작성 가능, 코드 퀄리티를 위해 errors.New를 담는 부분의 변수명의 시작을 err로 붙인다.
- **dictionary proj**
    
    ```go
    package mydict
    
    import "errors"
    
    //Dictionary type
    type Dictionary map[string] string 
    
    var errNotFound = errors.New("not found")
    var errWordExists = errors.New("that word already exists")
    var errCantUpdate = errors.New("cant update non-existing word")
    //Search for a word 
    func (d Dictionary) Search(word string) (string, error) {
    	value, exists := d[word]
    	if exists{
    		return value, nil
    	}
    	return "", errNotFound
    }
    
    //Add a word to the dictionary
    func (d Dictionary) Add(word, def string) error {
    	_, err := d.Search(word)
    	if err == errNotFound{
    		d[word] = def
    	}else if err == nil{
    		return errWordExists
    	}
    	return nil
    }
    
    //Update a word
    //포인터 리시버를 쓰지 않는 이유는 해시맵이 기본적으로 *를 포함하고 있기 때문이다.
    func (d Dictionary) Update(word, def string) error{
    	
    	_, err := d.Search(word)
    	switch err{
    	case nil:
    		d[word] = def
    	case errNotFound:
            return errCantUpdate
    	}
    	return nil
    }
    
    //Delete a word
    //포인터 리시버를 쓰지 않는 이유는 해시맵이 기본적으로 *를 포함하고 있기 때문이다.
    func (d Dictionary) Delete(word string) {
    	delete(d, word)
    }
    ```
    
- **URL Checker proj**
    
    ```go
    package main
    
    import (
    	"errors"
    	"fmt"
    	"net/http"
    )
    
    func main(){ 
    	urls := []string{
    		"https://www.airbnb.com",
    		"https://www.google.com",
    		"https://www.facebook.com",
    		"https://www.amazon.com",
    	}
    	results := make(map[string] string)
    
    	for _, url := range urls {
    		result := "OK"
    		err := hitURL(url)
    		if err != nil{
    			result = "FAILED"
    		}
    
    		results[url] = result
    	}
    	for url, result := range results{
    		fmt.Println(url, result)
    	}
    }
    
    var errRequestFailed = errors.New("request failed")
    
    func hitURL(url string) error { 
    	fmt.Println("checking :", url)
    	resp, err := http.Get(url)
    	if err != nil || resp.StatusCode >= 400{
    		return errRequestFailed
    	}
    	
    	return nil
    }
    ```
    
    - 이렇게 하나씩 처리하면 속도가 느리다. Goroutine으로 동시에 처리 하도록 변경한다.
    - go 키워드 사용
        - 메인 함수가 실행되는 동안만 go 루틴 유지
    - 채널은 Goroutine 사이의 데이터 전달
        - 고루틴으로부터 리턴을 받는 대신 채널을 통해서 값을 전달받음
        
        ```go
        func main(){ 
        	people := [2]string{"nico","whssodi"}
        	c := make(chan bool)
        	for _, person := range people {
        		go isSexy(person,c)
        	}
        // 고루틴 하나를 기다림 
        	result := <- c 
        	fmt.Println(result)
        }
        
        func isSexy(person string, c chan bool){
        	time.Sleep(time.Second*5)
        	c <- true
        }
        ```
        
        `← c` 은 blocking operation이다. 
        
        ```go
        package main
        
        import (
        	"fmt"
        	"net/http"
        )
        type requestResult struct{
        	url string
        	status string
        }
        
        func main(){ 
        	c := make(chan requestResult)
        	results := make(map[string] string)
        
        	urls := []string {
        		"http://www.baidu.com",
                "http://www.qq.com",
                "http://www.163.com",
                "http://www.baidu.com",
                "http://www.qq.com",
        	}
        	for _, url := range urls{ 
        		go hitURL(url, c)
        	}
        	
        	for i:=0; i<len(urls); i++ {
        		result := <-c
        		results[result.url] = result.status
        	}
        	for url, status := range results{
        		fmt.Println(url, status)
        	}
        }
        
         
        // chan<- 은 send only를 명시 
        func hitURL(url string, c chan<- requestResult) {
        	status := "OK"
        	resp, err := http.Get(url)
        	if err != nil  || resp.StatusCode >= 400{
        		status = "FAEILD"
        	}
        	
        	c <- requestResult{url : url, status : status}
        	
        }
        ```
        
- **Job** **scrapper proj (사람인)**
    - go 버전 jquery - goquery
        - `$ go get github.com/PuerkitoBio/goquery`
    - HTML 내부를 들여다 볼 수 있게 해줌 doc.Find(~~~)
    
    ```go
    package main
    
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
    
    var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=python"
    
    func main(){ 
    	var jobs []extractedJob
    	totalPages := getPages()
    	fmt.Println("Total pages: ", totalPages)
    
    	for i := 0; i< totalPages; i++ {
    		extractedJobs := getPage(i)
    		jobs = append(jobs, extractedJobs...)
    	}
    	writeJobs(jobs)
    	fmt.Println("Done, extracted", len(jobs))
    
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
    
    func getPage(page int) []extractedJob{
    	var jobs []extractedJob
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
    		job := extractJob(card)
    		jobs = append(jobs, job)
    	})
    	return jobs	
    
    }
    func extractJob(card *goquery.Selection) extractedJob{
    	id, _ := card.Attr("value")
    	title := cleanString(card.Find(".job_tit>a").Text())
    	location := cleanString(card.Find(".job_condition>span>a").Text())
    	summary := cleanString(card.Find(".job_sector").Text())
    	return extractedJob{
    		id: id,
            title: title,
            location: location,
            summary: summary,
        }
    	
    }
    
    func cleanString(str string) string {
    	return strings.Join(strings.Fields(strings.TrimSpace(str))," ")
    }
    
    func getPages() int{
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
    ```
    
    아래 코드는 2개 변경사항 반영
    
    1) main↔ getPage 와 getPage↔extractJob 간의 채널 생성 
    
    2) 검색어를 넣을 수 있게 변경 
    
    ```go
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
    	title := cleanString(card.Find(".job_tit>a").Text())
    	location := cleanString(card.Find(".job_condition>span>a").Text())
    	summary := cleanString(card.Find(".job_sector").Text())
    	c<- extractedJob{
    		id: id,
            title: title,
            location: location,
            summary: summary,
        }
    }
    
    func cleanString(str string) string {
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
    ```
    
- **web server with echo framework (+ scrapper proj)**
    - 변경사항 발생 시 재시작필요
    
    ```go
    package main
    
    import (
    	"fmt"
    	"os"
    	"strings"
    
    	"github.com/JeonJe/learngo/scrapper"
    	"github.com/labstack/echo"
    )
    
    const fileName string = "jobs.csv"
    // Handler
    func hello(c echo.Context) error {
    	return c.File("home.html")
      }
      
    func handleScrape(c echo.Context) error {
    	defer os.Remove(fileName)
    	term := strings.ToLower( scrapper.CleanString( c.FormValue("term")))
    	fmt.Println(term)
    	scrapper.Scrape(term)
    
    	return c.Attachment("jobs.csv", "job.csv")
      }
      
    func main(){
    	// Echo instance
    	e := echo.New()
      	// Routes
      	e.GET("/", hello)
      	e.POST("/scrape", handleScrape)
    	  // Start server
    	e.Logger.Fatal(e.Start(":1323"))
    }
    ```
    
</div>
</details>
