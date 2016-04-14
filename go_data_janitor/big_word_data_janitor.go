package main

import (
	tkz "github.com/jbowles/nlpt_tkz"
	"sync"
	"time"
)

var bigwordRootfp string = "/Users/jbowles/x/training_data/8_billion_word/"
var corporaRootfp string = "/Users/jbowles/x/training_data/corpora/"
var billionBenchmarkfp string = (bigwordRootfp + "1-billion-word-language-modeling-benchmark-r13output/")
var tkzr string = "unicode"

var BigWordData = map[int][]string{
	0:  {"wikipedia", bigwordRootfp + "enwiki-latest-pages-articles.xml", "datasets/enwiki-latest-pages-articles-tokenized.txt"},
	1:  {"file", bigwordRootfp + "news.2012.en.shuffled", "datasets/news.2012.en.shuffled.tokenized.txt"},
	2:  {"file", bigwordRootfp + "news.2013.en.shuffled", "datasets/news.2013.en.shuffled.tokenized.txt"},
	3:  {"dir", billionBenchmarkfp + "heldout-monolingual.tokenized.shuffled", "datasets/heldout-monolingual.tokenized.txt"},
	4:  {"dir", billionBenchmarkfp + "training-monolingual.tokenized.shuffled", "datasets/training-monolingual.tokenized.txt"},
	5:  {"dir", bigwordRootfp + "webbase_all", "datasets/umbc_webbase_tokenized.txt"},
	6:  {"dir", corporaRootfp + "big_text", "datasets/big_text_tokenized.txt"},
	7:  {"dir", corporaRootfp + "bible", "datasets/king_james_bible_tokenized.txt"},
	8:  {"dir", corporaRootfp + "reuters/training", "datasets/reuters_tokenized.txt"},
	9:  {"dir", corporaRootfp + "abc", "datasets/rural_science_news_australiabc_tokenized.txt"},
	10: {"dir", corporaRootfp + "europarl_raw", "datasets/euro_parliamnet_tokenized.txt"},
	11: {"dir", corporaRootfp + "gutenberg", "datasets/gutenberg_books_tokenized.txt"},
	12: {"dir", bigwordRootfp + "un", "datasets/united_nations_report_tokenized.txt"},
	13: {"file", bigwordRootfp + "hotels_all_name_description_addr.txt", "datasets/hotel_name_description_tokenized.txt"},
}

func streamDataOps() {
	var wg sync.WaitGroup
	for _, value := range BigWordData {
		switch value[0] {
		case "dir":
			wg.Add(1)
			go tkz.StreamTokenizedDirectory(&wg, time.Minute*90, value[1], value[2], tkzr)
		case "file":
			wg.Add(1)
			go tkz.StreamTokenizedFile(&wg, time.Minute*90, value[1], value[2], tkzr)
		case "wikipedia":
			wg.Add(1)
			go tkz.StreamTokenizedWikipediaDump(&wg, value[1], value[2], tkzr)
		}
	}
	wg.Wait()
}

func main() {
	streamDataOps()
}
