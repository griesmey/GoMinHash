# GoMinHash

## MinHash Background and Algorithm
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

