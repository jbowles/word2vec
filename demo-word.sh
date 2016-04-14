make
if [ ! -e text8 ]; then
  wget http://mattmahoney.net/dc/text8.zip -O text8.gz
  gzip -d text8.gz -f
fi
time ./word2vec -train text8 -output vectors.bin -cbow 1 -size 200 -window 8 -negative 25 -hs 0 -sample 1e-4 -threads 20 -binary 1 -iter 15
./distance vectors.bin


#time ./word2vec -train /Users/jbowles/x/training_data/corpora/words/en-basic -output /Users/jbowles/x/training_data/models/word2vec/en_basic_bow.bin -cbow 1 -size 200 -window 8 -negative 25 -hs 0 -sample 1e-4 -threads 20 -binary 1 -iter 15


#time ./word2vec -train /Users/jbowles/x/training_data/hotels/hotel_description.txt -output /Users/jbowles/x/training_data/models/word2vec/hotel_description_bow.bin -cbow 1 -size 200 -window 8 -negative 25 -hs 0 -sample 1e-4 -threads 20 -binary 1 -iter 15

#time ./word2vec -train /Users/jbowles/x/training_data/corpora/big_text/big_text.txt -output /Users/jbowles/x/training_data/models/word2vec/big_text_bow.bin -cbow 1 -size 200 -window 8 -negative 25 -hs 0 -sample 1e-4 -threads 20 -binary 1 -iter 15


################# skip-gram take aloooooot loooonger
#time ./word2vec -train text8 -output vectors.bin -cbow 0 -size 200 -window 8 -negative 25 -hs 0 -sample 1e-4 -threads 20 -binary 1 -iter 15

