package main

import "fmt"

//struct embedding
//playlist pgm

type song struct {
	title, artist string
}

type justdummy struct {
	dumm   string
	dumint int
}

type playlist struct {
	genre     string
	songs     []song
	justdummy //this is struct embedding.. This is a composition because i include another struct in the playlist struct.. Compared to oops, GO prefers compostion(struct embedding) instead of inheritance
	//justdummy justdummy  this is same as above, we need to write only just dummy and is called anonymous field. A struct field without a name is called anonymous field. An anonymous field gets its name automatically from its type
	string //this is also another anonymous field,, the struct field gets its name automatically from its type(name will be same as type, here name is string)
}

func main() {

	song1 := song{title: "wonder-wall", artist: "oasis"}
	song2 := song{title: "super-sonic", artist: "oasis"}
	song3 := song{title: "super-sdfsd", artist: "oasdffsis"}
	rock := playlist{
		genre:     "indie-Rock",
		songs:     []song{song1, song2, song3},
		justdummy: justdummy{dumm: "struct embedding , also anonymous filed", dumint: 52},
		string:    "Anonymous field",
	}

	fmt.Println(rock)
	fmt.Println()

	//change a title
	song := rock.songs[2] //this wont change it, because here song is a copy of rock.songs[2]
	song.artist = "peterrrr NOT CHANGES"
	fmt.Println("Here song variable is", song, "rock.song[2] is", rock.songs[2])

	//instead change it directly
	rock.songs[2].title = "CHANGED"

	fmt.Println()
	//pring the playlist
	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}

	//for anonymous field we can omit the name while calling also
	rock.justdummy.dumm = "anonymous field calling"
	rock.dumm = "anonymous field calling- this is same as above, the name can omit while calling"

	//which means instead of calling like this rock.justdummy.dumm, we need to just call rock.dumm
	fmt.Println("")
	fmt.Println(rock)
}
