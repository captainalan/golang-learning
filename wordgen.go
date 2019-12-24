/*
Program generates new (phonotactically acceptable) words

    $ ./wordgen
    noligaloo
    $ ./wordgen
    flabtastic
    $ ./wordgen
    hooban

Algorithm is up to you.
*/

package main
import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/* Using 'initial + rhyme" strategy means longer lists of syllables, but offloads logic
   of composing syllables to *data* rather than calculation. A similar strategy is used
   in Chinese Pīnyīn.
*/
type Syllable struct {
	Initial string // e.g. 'b' in 'bat'
	Rhyme   string // e.g. 'at' in 'bat'
}
// Words are composed of syllables
// TODO implement type for words

const MAX_LENGTH = 20 // TODO increase value; add moar data

// TODO get syllable info from reading in files rather than hardcoding them here.
var initials = [MAX_LENGTH]string{ "b", "s", "d", "f", "g", "h", "j", "k", "l", "m", "st", "sl", "sc", "sm", "br", "kr", "fr", "cl", "gr", "p" } 
var rhymes = [MAX_LENGTH]string{ "at", "et", "it", "ut", "ought", "ood", "oop", "ilk", "urk", "uck", "unk", "iff", "urf", "onk", "ack", "erm", "ee","ing" ,"ar" , "oo" } 

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// rand.Seed(101) // Pick a number u lyke

	argsWithoutProg := os.Args[1:]

	if (len(argsWithoutProg) == 0) {
		fmt.Println(GenerateWord())
	} else if (len(argsWithoutProg) == 1) {

		num, err := strconv.Atoi(argsWithoutProg[0]);

		if(err != nil) {
			fmt.Println("Bad first argument. Should be an int")
		} else {
			for i := 0; i < num; i++ {
				fmt.Println(GenerateWord())
		}
		
	    }
	}
}

func GenerateSyllable () Syllable {
	// WARNING depends on global data declared outside function
	initial := initials[rand.Intn(MAX_LENGTH)]
	rhyme   := rhymes[rand.Intn(MAX_LENGTH)]
	return Syllable{
		initial,
		rhyme,
	}
}

func GenerateWord () string {
	// Decide number of syllables
	ran_dumb := rand.Float64()
	num_syllables := -1;  // Initialize

	// ...arbitrarily setting some weights...
	// TODO make this statistically informed
	if (ran_dumb <= 0.05) {
		num_syllables = 4
	} else if (ran_dumb <= 0.2) {
		num_syllables = 3
	} else if (ran_dumb <= 0.5) {
		num_syllables = 2
	} else {
		num_syllables = 1
	}

	// Generate word to return
	word := "" // Word to return
	for i := 0; i < num_syllables; i++ {
		syllable := GenerateSyllable()
		word += (syllable.Initial + syllable.Rhyme) 
	}

	return word
}
