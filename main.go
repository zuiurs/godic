package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zuiurs/godic/local"
	"github.com/zuiurs/godic/thesaurus"
	"github.com/zuiurs/godic/weblio"
)

func main() {
	var syn, ant, local bool
	flag.BoolVar(&syn, "s", false, "search synonym words (short)")
	flag.BoolVar(&syn, "syn", false, "search synonym words")

	flag.BoolVar(&ant, "a", false, "search antonym words (short)")
	flag.BoolVar(&ant, "ant", false, "search antonym words")

	flag.BoolVar(&local, "l", false, "search from local dictionary (short)")
	flag.BoolVar(&local, "local", false, "search from local dictionary")

	flag.Parse()

	if (syn && ant) || (ant && local) || (local && syn) {
		fmt.Fprintln(os.Stderr, "Error: These option is exclusive")
		os.Exit(1)
	}

	args := flag.Args()

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Required at least an arguments")
		os.Exit(1)
	}

	if syn || ant {
		useThesaurus(args, syn)
	} else if local {
		useLocal(args)
	} else {
		useWeblio(args)
	}

}

func useThesaurus(args []string, syn bool) {
	for _, v := range args {
		fmt.Printf("--<%s>--\n", v)

		words, err := thesaurus.Search(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
		if syn {
			words = thesaurus.SynSort(words)
		} else {
			words = thesaurus.AntSort(words)
		}

		for i, w := range words {
			// show 5 words
			if i > 4 {
				break
			}

			if syn {
				if w.Class == thesaurus.ANTONYM {
					break
				}
			} else {
				if w.Class == thesaurus.SYNONYM {
					break
				}
			}

			fmt.Println(w)
		}
	}
}

func useWeblio(args []string) {
	for _, v := range args {
		fmt.Printf("--<%s>--\n", v)

		words, err := weblio.Search(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func useLocal(args []string) {
	for _, v := range args {
		fmt.Printf("--<%s>--\n", v)

		state, err := local.Search(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		fmt.Println(state)
	}
}
