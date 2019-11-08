// Find and replace script
//
// Taking the logic from the link below...
// http://xahlee.info/python/findreplace_regex.html
// ...and translating everything into Golang

/* 
Questions/learning topics
-------------------------
array vs slices vs lists---it seems slices are most favored
https://stackoverflow.com/questions/21326109/why-are-lists-used-infrequently-in-go

I'll go ahead and use (linked) lists. But the Golang idiomatic way to
do things seems to usually be slicey dicey. Hmm.
*/

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

var do_backup bool = true
var backup_suffix string = "~~" // Feels like emacs

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

	find_replace_list.PushFront(BarToBaz)
	find_replace_list.PushFront(FooToBar)
	// Push more things to find_replace_list as needed
	// 注意！ Order matters! We're modifying files here!!

	// Printing stuff to surpress unused variable errors
	fmt.Println("Will it work?")
	fmt.Println(do_backup, backup_suffix, find_replace_list.Len())
	fmt.Println(min_level, max_level, input_dir, file_extension_regex, file_list)

	// reads input_dir top, specified in top level code
	replace_string_in_file(input_dir) 
}

// Replaces all strings by regex in find_replace_list at fpath
func replace_string_in_file(fpath string) {
	// TODO remove!! temporary, for testing
	new_path := filepath.Join(fpath, "thing.txt")

	// Note: here, we read an *entire* file, rather than going
	// line by line
	input_file, err := ioutil.ReadFile(new_path) // For read access.
	input_text :=  string(input_file)

	if err != nil {
		panic(err) // blow up n quit
	}

	num_replaced := 0

	// iterate through find_replace_list
	for current := find_replace_list.Front();
	current != nil;
	current = current.Next() {
		// TODO parse regex from FindAndReplacePair, which is stored as string
		// foo, _ := regexp.CompilePOSIX("foo") // Second variable is error if any
		find_and_replace_pair := current.Value.(FindAndReplacePair) // Type assertion
		fmt.Printf("Find and replace pair: %V (type: %T)\n",
			find_and_replace_pair, find_and_replace_pair) 

		// TODO Do replacements
		// .RegularExpression.ReplaceAllString(input_text)
	}
	fmt.Print(input_text, num_replaced)
}

/*
Using the Python code below as "pseudo-code"

##################################################

def replace_string_in_file(fpath):
   "Replaces all strings by regex in find_replace_list at fpath."

   input_file = open(fpath, "r", encoding="utf-8")

   try:
      file_content = input_file.read()
   except UnicodeDecodeError:
      print("UnicodeDecodeError:{:s}".format(input_file))
      return

   input_file.close()

   num_replaced = 0
   for a_pair in find_replace_list:
      tem_tuple = re.subn(a_pair[0], a_pair[1], file_content)
      output_text = tem_tuple[0]
      num_replaced += tem_tuple[1]
      file_content = output_text

   if (num_replaced > 0):
      print(("◆ changed %d %s" % (num_replaced, fpath) ))

      if do_backup:
         shutil.copy2(fpath, fpath + backup_suffix)

      output_file = open(fpath, "r+b")
      output_file.read() # we do this way to preserve file creation date
      output_file.seek(0)
      output_file.write(output_text.encode("utf-8"))
      output_file.truncate()
      output_file.close()

##################################################
*/

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


