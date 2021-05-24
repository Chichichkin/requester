package internal

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type Requester struct {
	NumOfRoutines int
	Urls          []string
}

func (r *Requester) Run() ([]string, error) {
	err := r.ArgsParser(os.Args[1:])
	if err != nil {
		return nil, err
	}
	log.Printf("curent number of gorotines: %d", r.NumOfRoutines)

	var wg sync.WaitGroup
	urls := make(chan string, len(r.Urls))
	results := make(chan string, len(r.Urls))

	for i := 0; i < r.NumOfRoutines; i++ {
		wg.Add(1)
		go requestWorker(&wg, urls, results)
	}
	// sending tasks for workers
	for _, url := range r.Urls {
		urls <- url
	}
	// closing tasks channel
	close(urls)

	// wait until all workers done their job
	wg.Wait()
	var listResults []string
	for i := 0; i < len(r.Urls); i++ {
		listResults = append(listResults, <-results)
		log.Println(listResults[i])
	}
	return listResults, nil
}

func (r *Requester) ArgsParser(args []string) (err error) {
	//check if there any args
	if args == nil {
		return errors.New("no arguments given")
	}
	//check if in args specified number of max goroutines
	i, found := Find(args, "-parallel")
	if found {
		if numOfRoutines, err := strconv.Atoi(args[i+1]); err == nil {
			r.NumOfRoutines = numOfRoutines
			args = append(args[:i], args[i+2:]...)
		} else {
			return errors.New("missing number of routines")
		}
	}
	i, found = Find(args, "-p")
	if found {
		if numOfRoutines, err := strconv.Atoi(args[i+1]); err == nil {
			r.NumOfRoutines = numOfRoutines
			args = append(args[:i], args[i+2:]...)
		} else {
			return errors.New("missing number of routines")
		}
	}
	//check if there is path to file with sites to visit
	i, found = Find(args, "-file")
	if found {
		return r.ReadFromFile(args[i+1])
	}
	i, found = Find(args, "-f")
	if found {
		return r.ReadFromFile(args[i+1])
	}
	//parse args for urls
	r.Urls = args
	return nil
}
func (r *Requester) ReadFromFile(filePath string) (err error) {
	if !filepath.IsAbs(filePath) {
		currPath, err := os.Getwd()
		if err != nil {
			return err
		}
		filePath = filepath.Join(currPath, filePath)
		if !filepath.IsAbs(filePath) {
			return errors.New("wrong path")
		}
	}
	if _, err := os.Stat(filePath); err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			return errors.New("failed opening file")
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var urls []string

		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
		file.Close()
		r.Urls = urls
		return nil

	} else if os.IsNotExist(err) {
		return errors.New("path to file doesn't exist")
	} else {
		return errors.New("problems with file")
	}
}
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
func requestWorker(wg *sync.WaitGroup, tasks <-chan string, results chan<- string) {
	for url := range tasks {
		res, err := MakeRequest(url)
		if err != nil {
			results <- err.Error()
		} else {
			results <- fmt.Sprintf("%s %s", url, res)
		}

	}

	// done with worker
	wg.Done()
}
