package Json

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie /** 只有可导出的成员变量才会被序列化到json里面，所以里面的项全部Capital Letter */
type Movie struct {
	Title  string
	Year   int  `json:"released"`        // 成员标签: 指定Json中字段名称为 released
	Color  bool `json:"color,omitempty"` // 成员标签： 指定Json中字段名称为color， 且 omitempty 忽略0值或空值
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main() {
	data0, err0 := json.Marshal(movies)
	if err0 != nil {
		log.Fatalf("JSON marshaling failed: %s", err0)
	}
	fmt.Println("Json Marshal function output is : ")
	fmt.Printf("%s\n", data0)

	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Println("Json MarshalIndent function output is : ")
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON Unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // [{Casablanca} {Cool Hand Luke} {Bullitt}]
}
