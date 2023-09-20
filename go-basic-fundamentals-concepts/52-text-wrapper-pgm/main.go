package main

import "fmt"

func main() {
	const text = `Türkiye'nin başkenti Ankara'dr. Translation: The capital of Turkey is Ankara. Benim adm Ahmet, memnun oldum.Translation: My name is Ahmet, pleased to meet you.
	
Türk mutfağ dünya çapnda ünlüdür.Translation: Turkish cuisine is famous worldwide. İstanbul Boğaz'nda gezmek harika bir deneyimdir. Translation: Exploring the Bosphorus in Istanbul is a fantastic experience.
	
Türkçe öğrenmek zorlu olabilir, ancak keyifli bir deneyimdir.Translation: Learning Turkish can be challenging, but it's an enjoyable experience.I hope you find these Turkish sentences helpful! If you have any specific phrases or sentences you'd like to see, please feel free to ask.`

	const maxwidth = 40
	var lw int //line width

	for _, r := range text {
		fmt.Printf("%c", r)
		if lw++; lw > maxwidth {
			//fmt.Printf("<[%d>", lw)
			fmt.Println()
			lw = 0
		} else if r == '\n' {
			lw = 0
		}

	}
	fmt.Println()
}
