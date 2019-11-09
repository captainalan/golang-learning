// Find and replace script
// =======================
// 
// Taking some of the logic from the link below...
// http://xahlee.info/python/findreplace_regex.html
// ...and translating everything into Golang
// Some features of the above script are not yet implemented...

package main

import (
	"fmt"
	"container/list"
 	"io/ioutil"
 	"path/filepath"
	"regexp"
)

// Declare an array of filenames
var file_list list.List 
// Input directory
var input_dir string = filepath.Join("foo", "bar")
// Extension(s) of files to do find and replace in
var file_extension_regex = regexp.MustCompile(`\.txt$`)

var min_level uint = 1 // Files and dirs inside input_dir are level 1.
var max_level uint = 9 // inclusive

// Things will be added to this list in main()
var find_replace_list list.List

// Find and replace pairs here. Each is a 2-tuple, represented
// as a struct. The first element is regex object, the second
// is replace string.
type FindAndReplacePair struct {
	RegularExpression string // gets compiled later
	Replacement string
}

func main() {
	// Define find and replace pairs (here, I just do one)
	FooToBar := FindAndReplacePair{ "foo", "bar" }
	BarToBaz := FindAndReplacePair{ "bar", "baz" }
	BazToFoo := FindAndReplacePair{ "baz", "foo" }

	find_replace_list.PushFront(FooToBar)
	find_replace_list.PushBack(BarToBaz)
	find_replace_list.PushBack(BazToFoo)
	// Push more things to find_replace_list as needed
	// 注意！ Order matters! We're modifying files here!!

	// Printing stuff to surpress unused variable errors
	fmt.Println("Will it work?")
	fmt.Println(find_replace_list.Len())
	fmt.Println(min_level, max_level, input_dir, file_extension_regex, file_list)

	// reads input_dir top, specified in top level code

	// TODO remove!! temporary, for testing
	new_path := filepath.Join(input_dir, "thing.txt")
	replace_string_in_file(new_path) 
}

// Replaces all strings by regex in find_replace_list at fpath
func replace_string_in_file(fpath string) {

	// Note: here, we read an *entire* file, rather than going
	// line by line
	input_file, err := ioutil.ReadFile(fpath) // For read access.
	input_text :=  string(input_file)

	if err != nil {
		panic(err) // blow up n quit
	}

	// iterate through find_replace_list
	for current := find_replace_list.Front();
	current != nil;
	current = current.Next() {
		// Type assertion used here!
		find_and_replace_pair := current.Value.(FindAndReplacePair) 

		// Compile regexp to use for find and replace task
		current_regexp, _ := regexp.CompilePOSIX(
			find_and_replace_pair.RegularExpression)

		// Perform find and replace and store result
		new_text := current_regexp.ReplaceAllString(input_text,
			find_and_replace_pair.Replacement)

		// Write the result to file
		// TODO find what appropriate permissions should be 
		ioutil.WriteFile(fpath, []byte(new_text), 0755)

		// TODO Write some more useful output
		fmt.Println("Text with replacements: ", new_text)
	}
}

/*
print(datetime.datetime.now())
print("Input Dir:", input_dir)
for x in find_replace_list:
   print("Find regex:「{}」".format(x[0]))
   print("Replace pattern:「{}」".format(x[1]))
   print("\n")

if (len(file_list) != 0):
   for ff in file_list: replace_string_in_file(os.path.normpath(ff) )
else:
    for dirPath, subdirList, fileList in os.walk(input_dir):
        curDirLevel = dirPath.count( os.sep) - input_dir.count( os.sep)
        curFileLevel = curDirLevel + 1
        if min_level <= curFileLevel <= max_level:
            for fName in fileList:
                if (re.search(file_extension_regex, fName, re.U)):
                    replace_string_in_file(dirPath + os.sep + fName)

print("Done.")
*/


/* 
Questions/learning topics
-------------------------
array vs slices vs lists---it seems slices are most favored
https://stackoverflow.com/questions/21326109/why-are-lists-used-infrequently-in-go

I'll go ahead and use (linked) lists. But the Golang idiomatic way to
do things seems to usually be slicey dicey. Hmm.
*/
