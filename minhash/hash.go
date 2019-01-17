package minhash

/*
 ------------------ MinHash ----------------------
 MinHash is a locality sensitive hashing technique for hashing
 documents to determine similarity.  This approach is approx. equal
 to the Jaccard Index of tokenized word sets.

 For comparison, we will use the MinHash algorithm to calculate short
 signature vectors to represent the documents. These MinHash
 signatures can then be compared quickly by counting the ratio
 of components in which the signatures agree over the total number of components.

 MinHash is approx. equal to Jaccard.  We show this in the hash_test unittest file.

 Definitions:
    Shingle - unsigned 32 bit integer which represents the id of
              the string hashed

			You can think of shingles as being unique ids for all possible strings
			before they are hashed
			Str = x, f(x) -> g where f'(g) -> x

			The same valued string will always produce the same shingle value
			Str = x = y, f(x) = f(y) = g, where x = y

	MinHash - Hashed signature of a document; It's computed by taking the
			  minimum hashed value for all k hash functions in the
			  hash family, H.  Once the min hash value is found
			  then it is appended to the signature.

	SingleSet - Set of shingles; A single shingle set represents one document

*/

import (
	"hash/crc32"
	"math/rand"
	"strconv"
	"strings"
)

const (
	// 2**32-1
	maxShingleID = 4294967295

	//largest prime number above maxSingleID
	nextPrime = 4294967311

	// provides us with 5% granularity
	numHashes = 20
)

type MinHash []int
type shingle uint32
type shingleSet map[shingle]bool

var coeffA []int
var coeffB []int

func init() {
	coeffA = generateCoeffs()
	coeffB = generateCoeffs()
}

// Str returns the string representation
// of a MinHash.
// Iterate though the size of the MinHash signature
// Returns a list of MinHash values delimited by " "
func (this MinHash) Str() string {
	ss := []string{}

	for _, h := range this {
		ss = append(ss, strconv.FormatInt(int64(h), 10))
	}

	return strings.Join(ss, " ")
}

// Length of MinHash signature
func (this MinHash) Len() int {
	return len(this)
}

// GenerateCoeffs generates a list of random
// coefficients. These coefficients are used in
// subsequent calls to the Min Hash calculator
// and the same coefficients are used for all
// documents in the analysis group
func generateCoeffs() []int {

	random := []int{}
	seen := map[int]bool{}

	for i := 0; i < numHashes; i++ {
		randIndex := rand.Intn(maxShingleID)

		for {
			_, found := seen[randIndex]
			if !found {
				break
			}

			randIndex = rand.Intn(maxShingleID)
		}

		random = append(random, randIndex)
		seen[randIndex] = true
	}
	return random
}

// Hash a string to a 32 bit int
func hash(s string) uint32 {
	h := crc32.New(crc32.IEEETable)
	h.Write([]byte(s))

	return h.Sum32()
}

// string2Shingle converts a string into
// a shingle which is a hashed 32bit unsigned int
func string2Shingle(s string) shingle {
	return shingle(hash(s))
}

// MinHash algorithm
//
// 1. For each shingle compute a hash code
//
// 2. For each shingle set, find the minimum hash code
//    and append to the signature
//
// 	  	- For each hash we have a set of coefficients,
// 		  (A and B), respectively.
//
// 	 	    - We use thse to compute the hash code.
//
// 3. Return the signature to the caller which is the calculated MinHash for that ShingleSet
func calculateMinHash(ss shingleSet) MinHash {
	signature := MinHash{}

	var minHashCode int
	for i := 0; i < numHashes; i++ {
		// make min hash code to be more than
		// the max possible value output by hash
		minHashCode = nextPrime + 1

		for shingle := range ss {
			hashCode := (coeffA[i]*int(shingle) + coeffB[i]) % nextPrime

			if hashCode < minHashCode {
				minHashCode = hashCode
			}
		}

		signature = append(signature, minHashCode)
	}

	return signature
}

// doc2ShingleSet generates a shingle set from a
// string document. Each shingle contain three
// tokens from the document.  We scan per every 3 words
// in the order they appear in the document.
//  Each one of these triples is a shingle
func doc2ShingleSet(d string) shingleSet {
	shingles := shingleSet{}

	tokens := strings.Split(d, " ")
	for i := 0; i < len(tokens)-2; i++ {
		words := strings.ToLower(tokens[i] + " " + tokens[i+1] + " " + tokens[i+2])

		shingles[string2Shingle(words)] = true
	}

	return shingles
}

// Calculate MinHash similarity with strings
// Generate the min hash for both strings
// then calculate the similarity between both
// min hash signatures
func stringSimilarity(s0, s1 string) float64 {
	ss0 := GenerateMinHash(s0)
	ss1 := GenerateMinHash(s1)

	return minHashSimilarity(ss0, ss1)
}

// MinHashFromStr takes a min hash string
// and converts it into a MinHash object
func MinHashFromStr(mh string) (MinHash, error) {
	minHash := MinHash{}

	for _, s := range strings.Split(mh, " ") {
		if i, err := strconv.ParseInt(s, 10, 64); err != nil {
			return minHash, err
		} else {
			minHash = append(minHash, int(i))
		}
	}
	return minHash, nil
}

// Compute number of signature matches / num hashes
// to get approx. Jaccard distance
func minHashSimilarity(m1, m2 MinHash) float64 {
	matches := 0.0

	for i := range m1 {
		if m1[i] == m2[i] {
			matches += 1
		}
	}

	return matches / float64(numHashes)
}

// MinHashSimilar takes two min hash strings as input
// and returns whether they are similar.
func MinHashSimilar(left, right string) bool {
	l, _ := MinHashFromStr(left)
	r, _ := MinHashFromStr(right)
	return minHashSimilarity(l, r) > SimilarityThreshold
}

// compare two strings to see if they are similar
func StringsSimilar(left, right string) bool {
	return minHashSimilarity(GenerateMinHash(left), GenerateMinHash(right)) > SimilarityThreshold
}

// GenerateMinHash generates a minhash from a document string
// This is used to crate a MinHash from scratch
func GenerateMinHash(d string) MinHash {
	ss := doc2ShingleSet(d)

	return calculateMinHash(ss)
}
