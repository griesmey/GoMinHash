package algo

import (
	"log"
	"regexp"
	"strings"
)

// We say that two jobs are too similar if they
// contain 79.9% of the same unique words
const (
	SimilarityThreshold = 0.799
)

type WordSet struct {
	jid        int64
	membership map[string]bool
	length     int
}

func NewWordSet() *WordSet {
	var wordSet WordSet
	wordSet.membership = map[string]bool{}
	wordSet.length = 0

	return &wordSet
}

func NewWordSetFromText(text string) *WordSet {
	ws := NewWordSet()

	tokens := strings.Split(text, " ")
	for i := 0; i < len(tokens)-2; i++ {
		w := strings.ToLower(tokens[i] + " " + tokens[i+1] + " " + tokens[i+2])

		ws.membership[w] = true
	}

	ws.length = len(ws.membership)

	return ws
}

func (this *WordSet) Add(word string) {
	lower_word := strings.ToLower(word)
	if _, ok := this.membership[lower_word]; !ok {
		this.length += 1
	}
	this.membership[lower_word] = true
}

func (this *WordSet) Remove(word string) {
	lower_word := strings.ToLower(word)

	if _, ok := this.membership[lower_word]; ok {
		this.length -= 1
	}
	this.membership[word] = false
}

func (this *WordSet) Len() int {
	return this.length
}

func (this *WordSet) Contains(word string) bool {

	if _, ok := this.membership[strings.ToLower(word)]; ok {
		return true
	}
	return false
}

func (this *WordSet) Intersection(other *WordSet) int {
	// Compute set intersection
	intersection := 0

	for key := range this.membership {
		if other.Contains(key) {
			intersection += 1
		}
	}
	return intersection
}

// Similar is based on Jaccard index
// If their similarity thresholds are too high then we return true
func Similar(left, right string) bool {
	return JaccardSimilarity(left, right) >= SimilarityThreshold
}

// JaccardDistance calculate the similarity between two wordSets
// Useful for determining if two block of texts are similar
//
// Explanation: https://en.wikipedia.org/wiki/Jaccard_index
func JaccardDistance(left, right *WordSet) float64 {
	intersection := float64(left.Intersection(right))
	union := float64(left.Len()+right.Len()) - intersection

	return intersection / union
}

// Same as Similarity but for word sets
func SimilarWordSets(left, right *WordSet) bool {
	return JaccardDistance(left, right) >= SimilarityThreshold
}

// Calculates similarity between two arbitrary strings
func JaccardSimilarity(left, right string) float64 {
	l := NewWordSetFromText(left)
	r := NewWordSetFromText(right)

	return JaccardDistance(l, r)
}

func removeNonAlphanumeric(text string) string {
	// We only want alpha numeric characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(text, " ")

	return processedString
}
