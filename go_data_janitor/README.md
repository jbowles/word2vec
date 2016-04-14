# Sources
Got most of these sources from the original word2vec project

## News crawler

```sh
wget http://www.statmt.org/wmt14/training-monolingual-news-crawl/news.2012.en.shuffled.gz
wget http://www.statmt.org/wmt14/training-monolingual-news-crawl/news.2013.en.shuffled.gz
gzip -d news.2012.en.shuffled.gz
gzip -d news.2013.en.shuffled.gz
```

## U.N. Director Reports translations

```sh
wget http://www.statmt.org/wmt13/training-parallel-un.tgz
tar -xvf training-parallel-un.tgz '*.en'
```

## 1B lang modelling benchmark

```sh
wget http://www.statmt.org/lm-benchmark/1-billion-word-language-modeling-benchmark-r13output.tar.gz
tar -xvf 1-billion-word-language-modeling-benchmark-r13output.tar.gz
```

## UMBC

```sh
wget http://ebiquity.umbc.edu/redirect/to/resource/id/351/UMBC-webbase-corpus
tar -xvf umbc_webbase_corpus.tar.gz '*.txt'
wget http://dumps.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles.xml.bz2
```


## Wikipedia

```
wget http://dumps.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles.xml.bz2
bzip2 -c -d enwiki-latest-pages-articles.xml.bz2
```
