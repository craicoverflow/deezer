package deezer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetArtist(t *testing.T) {
	result, err := GetArtist(27)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetArtistTop(t *testing.T) {
	result, err := GetArtistTop(27, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetArtistAlbums(t *testing.T) {
	result, err := GetArtistAlbums(27, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetArtistComments(t *testing.T) {
	result, err := GetArtistComments(27, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetArtistFans(t *testing.T) {
	result, err := GetArtistFans(27, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetArtistRelated(t *testing.T) {
	result, err := GetArtistRelated(27, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetArtistRadio(t *testing.T) {
	result, err := GetArtistRadio(27, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
