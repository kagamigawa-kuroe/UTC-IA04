package mypackage

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func IsPalindrome(word string) bool{
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i]{
			return false
		}
	}
	return true
}

func Palindromes(words []string) (l []string){
	for i := range words{
		if(IsPalindrome(words[i])){
			l = append(l,words[i])
		}
	}
	return l
}

func Footprint(s string) (footprint string){
	b := []byte(s)
	sort.Slice(b,func(i,j int) bool{
		return b[i] < b[j]
	})
	footprint = string(b)
	return footprint
}

func Anagrams(words []string) (anagrams map[string][]string){
	for i:= range words{
		a := anagrams[Footprint(words[i])]
		a = append(a,words[i])
	}
	return anagrams
}

func DictFromFile(filename string) (dict []string){
	var all []string

	fi, err:= os.Open(filename)
	br := bufio.NewReader(fi)

	if err!=nil{
		fmt.Println(err) //打开文件错误
		return 
	}

	// fmt.Println("read start")
    for {
        a, _, c := br.ReadLine()
        if c == io.EOF {
            break
        }
		// fmt.Println(string(a))
        all = append(all,string(a))
    }
	// fmt.Println("read end")
	var anagrams map[string][]string = Anagrams(all)
	var words []string = Palindromes(all)

	/// LONGEST Palindrome
	max := 0
	for i:=range words{
		if(len(words[max])<=len(words[i])){
			max = i
		}
	}
	fmt.Println("-----------------------")
	fmt.Println("the LONGEST WORD IS:")
	fmt.Println(words[max])
	fmt.Println("-----------------------")


	/// AGENT
	for i,j := range anagrams{
		if(i=="AGENT"){
			fmt.Println("the AGENT HAS:")
			fmt.Println(j)
			fmt.Println("-----------------------")
		}
	}
	
	var max_v string = Footprint(all[0])
	/// LONGEST anagram
	for i,j := range anagrams{
		if(len(j)>len(anagrams[max_v])){
			max_v = i
		}
	}
	fmt.Println("THE LONGEST ANAGRAM IS:")
	fmt.Println(max_v)
	fmt.Println("-----------------------")

	return all
}