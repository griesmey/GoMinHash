# GoMinHash

## MinHash Background and Algorithm

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


https://en.wikipedia.org/wiki/MinHash

## How to test
$ go test

## Example Usage
```golang
left := "I made this."
right := "You made this? I made this."
sim := minHashSimilarity(GenerateMinHash(left), GenerateMinHash(right))
```
`sim` will contain the Jaccard Similarity between the `left` and `right` strings.  By generating the min hash signatures we can quickly obtain similarities between the two string segments.  

