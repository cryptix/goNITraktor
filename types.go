package goNITraktor

import (
	"encoding/xml"
	"time"
)

type NmlRoot struct {
	XMLName xml.Name
	Head    struct {
		Company string `xml:"COMPANY,attr"`
		Program string `xml:"PROGRAM,attr"`
	} `xml:"HEAD"`

	Collection Collection `xml:"COLLECTION"`
	Playlists  Playlist   `xml:"PLAYLISTS"`
}

type Collection struct {
	EntryCnt int               `xml:"ENTRIES,attr"`
	Entries  []CollectionEntry `xml:"ENTRY"`
}

type CollectionEntry struct {
	Titel    string                  `xml:"TITLE,attr"`
	Artist   string                  `xml:"ARTIST,attr"`
	Location CollectionEntryLocation `xml:"LOCATION"`
	Info     CollectionEntryInfo     `xml:"INFO"`
}

type CollectionEntryLocation struct {
	Dir    string `xml:"DIR,attr"`
	File   string `xml:"FILE,attr"`
	Volume string `xml:"VOLUME,attr"`
}

type CollectionEntryInfo struct {
	Genre       string      `xml:"GENRE,attr"`
	Playtime    int         `xml:"PLAYTIME,attr"`
	Bitrate     int         `xml:"BITRATE,attr"`
	ImportDate  TraktorDate `xml:"IMPORT_DATE,attr"`
	ReleaseDate TraktorDate `xml:"RELEASE_DATE,attr"`
	LastPlayed  TraktorDate `xml:"LAST_PLAYED,attr"`
}

// Embedded version
// type TraktorDate struct {
// 	time.Time
// }
// func (t *TraktorDate) UnmarshalXMLAttr(attr xml.Attr) (err error) {
// 	t.Time, err = time.Parse("2006/1/2", attr.Value)
// 	return
// }
// func (t TraktorDate) String() string {
// 	return t.Time.Format("02.01.2006")
// }

type TraktorDate time.Time

func (t *TraktorDate) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	newT, err := time.Parse("2006/1/2", attr.Value)
	if err != nil {
		return err
	}

	*t = TraktorDate(newT)
	return nil
}

func (t TraktorDate) String() string {
	return time.Time(t).Format("02.01.2006")
}

type Playlist struct {
	Nodes []PlaylistNode `xml:"NODE"`
}

type PlaylistNode struct {
	Name     string           `xml:"NAME,attr"`
	Subnodes []PlaylistNode   `xml:"SUBNODES>NODE"`
	Playlist PlaylistPlaylist `xml:"PLAYLIST"`
}

type PlaylistPlaylist struct {
	EntryCnt int             `xml:"ENTRIES,attr"`
	Type     string          `xml:"TYPE,attr"`
	UUID     string          `xml:"UUID,attr"`
	Entries  []PlaylistEntry `xml:"ENTRY"`
}

type PlaylistEntry struct {
	PrimaryKey   PlaylistPrimaryKey   `xml:"PRIMARYKEY"`
	ExtendedData PlaylistExtendedData `xml:"EXTENDEDDATA"`
}

type PlaylistPrimaryKey struct {
	Type string `xml:"TYPE,attr"`
	Key  string `xml:"KEY,attr"`
}

type PlaylistExtendedData struct {
	Deck      int     `xml:"DECK,attr"`
	Duration  float64 `xml:"DURATION,attr"`
	Type      string  `xml:"EXTENDEDTYPE,attr"`
	StartDate int     `xml:"STARTDATE,attr"`
	StartTime int     `xml:"STARTTIME,attr"`
}
