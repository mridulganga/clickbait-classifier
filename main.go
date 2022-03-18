package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/navossoc/bayesian"
)

const (
	CB    bayesian.Class = "cb"
	NONCB bayesian.Class = "noncb"
)

var classifier *bayesian.Classifier

func learn() {

	// read csv file
	dat, err := os.ReadFile("clickbait_data.csv")
	r := csv.NewReader(strings.NewReader(string(dat)))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// class samples
	cb := []string{}
	noncb := []string{}

	// parse csv
	for _, record := range records[1:] {

		// remove stopwords
		words := stopwords.CleanString(record[0], "en", true)
		wordsList := strings.Split(words, " ")

		// put the data in right set
		if record[1] == "1" {
			cb = append(cb, wordsList...)
		} else {
			noncb = append(noncb, wordsList...)
		}
	}

	// create model
	classifier = bayesian.NewClassifier(CB, NONCB)
	classifier.Learn(cb, CB)
	classifier.Learn(noncb, NONCB)

	// save model to file?
	err = classifier.WriteClassesToFile("clickbait")
	if err != nil {
		fmt.Println(err)
	}

}

func check(input string) (string, float64) {
	// input text
	text := input

	text = stopwords.CleanString(text, "en", true)
	textWordsList := strings.Split(text, " ")

	// log scores
	// scores, likely, _ := classifier.LogScores(
	// 	textWordsList,
	// )
	// fmt.Println(scores, likely)

	// prob. scores
	probs, likely, _ := classifier.ProbScores(
		textWordsList,
	)
	_ = likely
	// fmt.Println(probs, likely)

	// check where its classified
	if probs[0] > probs[1] {
		return "CB", probs[0]
	} else {
		return "non CB", probs[1]
	}
}

func main() {
	learn()
	// check these titles
	titles := []string{
		"Whitehouse says climate change is not real",
		"You’ll Never Believe this Simple Method to Rank High in Google",
		"8 Content Marketing Fails That You Need to Know",
		"21 resons why minecraft will never die",
		"This Weird Trick Increased the Conversion Rate by 110%",
		"This Is What Happens if You Stop Worrying Too Much about SEO",
		"Why did a ten-time All-Star sports hero go to jail?",
		"10 things Spotify does not want you to know",
	}
	for _, t := range titles {
		ty, s := check(t)
		fmt.Println(t, "- (", ty, s, ")")
	}
}
