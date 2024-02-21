package NFOGenerator

import "encoding/xml"

type Movie struct {
	XMLName       xml.Name `xml:"movie"`
	Text          string   `xml:",chardata"`
	Title         string   `xml:"title"`
	Originaltitle string   `xml:"originaltitle"`
	Sorttitle     string   `xml:"sorttitle"`
	Customrating  string   `xml:"customrating"`
	Mpaa          string   `xml:"mpaa"`
	Set           string   `xml:"set"`
	Studio        string   `xml:"studio"`
	Year          string   `xml:"year"`
	Outline       string   `xml:"outline"`
	Plot          string   `xml:"plot"`
	Runtime       string   `xml:"runtime"`
	Director      string   `xml:"director"`
	Poster        string   `xml:"poster"`
	Thumb         string   `xml:"thumb"`
	Fanart        string   `xml:"fanart"`
	Actor         struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"name"`
		Thumb string `xml:"thumb"`
	} `xml:"actor"`
	Maker       string   `xml:"maker"`
	Label       string   `xml:"label"`
	Tag         []string `xml:"tag"`
	Genre       []string `xml:"genre"`
	Num         string   `xml:"num"`
	Premiered   string   `xml:"premiered"`
	Releasedate string   `xml:"releasedate"`
	Release     string   `xml:"release"`
	Cover       string   `xml:"cover"`
	Website     string   `xml:"website"`
}
