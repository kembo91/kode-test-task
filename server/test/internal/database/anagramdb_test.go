package database_test

import (
	"testing"
)

func TestRetrieveAllAnagrams(t *testing.T) {
	test := []string{"fisha", "sahfi", "amail", "alima"}
	exp := []string{"fisha", "sahfi", "amail", "alima"}
	for _, v := range test {
		err := db.InsertAnagram(v)
		if err != nil {
			t.Error(err)
		}
	}
	got, err := db.RetrieveAllAnagrams()
	if err != nil {
		t.Error(err)
	}
	if !sameStringSlice(got, exp) {
		t.Errorf("Retrieve all anagrams failed")
	}
}

func TestInsertAnagramCorrect(t *testing.T) {
	err := db.InsertAnagram("aabbcc")
	if err != nil {
		t.Error(err)
	}
}

func TestInsertAnagramIncorrect(t *testing.T) {
	err := db.InsertAnagram("aabbcc aab")
	if err == nil {
		t.Error(err)
	}
}

func TestInsertExistingAnagram(t *testing.T) {
	err := db.InsertAnagram("anagram")
	if err != nil {
		t.Error(err)
	}
	err = db.InsertAnagram("anagram")
	if err == nil {
		t.Errorf("Existing anagram check failed")
	}
}

func TestRetrieveAnagram(t *testing.T) {
	var tests = []struct {
		test []string
		exp  []string
	}{
		{[]string{"dog", "god", "odg", "gdo"}, []string{"dog", "god", "odg", "gdo"}},
		{[]string{"mail", "lima", "mila", "alim"}, []string{"mail", "lima", "mila", "alim"}},
		{[]string{"fish", "shif", "fihs", "shfi"}, []string{"fish", "shif", "fihs", "shfi"}},
	}
	for _, tv := range tests {
		for _, teststr := range tv.test {
			err := db.InsertAnagram(teststr)
			if err != nil {
				t.Error(err)
			}
		}
		got, err := db.RetrieveQueryAnagram(tv.test[0])
		if err != nil {
			t.Error(err)
		}
		if !sameStringSlice(tv.exp, got) {
			t.Errorf("Anagram retrieval is wrong")
		}
	}
}

func TestRetrieveAnagramIncorrect(t *testing.T) {
	err := db.InsertAnagram("hello")
	if err != nil {
		t.Error(err)
	}
	_, err = db.RetrieveQueryAnagram("wrong")
	if err == nil {
		t.Error("Found inexistent anagram")
	}
}

func sameStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	amap := make(map[string]int)
	bmap := make(map[string]int)
	for _, s := range a {
		amap[s]++
	}
	for _, s := range b {
		bmap[s]++
	}
	for aMapKey, aMapV := range amap {
		if bmap[aMapKey] != aMapV {
			return false
		}
	}
	return true
}
