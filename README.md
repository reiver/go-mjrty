# go-mjrty

**go-mjrty** is a Go library which implements of the **mjrty** algorithm.

The **mjrty** algorithm figures out what is the **majority** of a sequence of items in a **single pass**.

And thus does this in linear time complexity (and constant space complexity) --
in O(n) time (and O(1) space).

This makes the **mjrty** suitable for some types of data streaming.

(Although there are some limitations due to the max size of the `uint64` type in Golang.
So in cases where the data stream is expected to be *infinite*, special care must be taken.)

**NOTE that if a sequence does NOT have a majority item, then what the mjrty algorithm returns is nonsense.**
This is just how the algorithm works.
(The is a trade off for having an algorithm that can calculate this in a single pass)

## Description

For example, consider the following sequence of strings:

1. "Apple",
2. "Apple",
3. "Apple",
4. "Cherry",
5. "Cherry",
6. "Banana",
7. "Banana",
8. "Cherry",
9. "Cherry",
10. "Cherry",
11. "Banana",
12. "Cherry",
13. "Cherry".

In this example, our sequence has three different elements in the sequence:
`"Apple"`, `"Banana"` and `"Cherry"`.

The counts for each of these is:

| Element    | Count         |
| ---------- | -------------:|
| `"Apple"`  | 3             |
| `"Banana"` | 3             |
| `"Cherry"` | 7             |


Also, our **total** number of items is: 13.

A majority is what is colloquially thought of as `50% + 1`.
(Although more accurately it is `⌊(total × 50%) + 1⌋`.)

In our example, which has **a total of 13 items**, for something to get the majority
it must have a count of **at least**:

```
⌊(total × 50%) + 1⌋ = ⌊(13 × 50%) + 1⌋ = ⌊(13 × 0.5) + 1⌋ = ⌊(6.5) + 1⌋ = ⌊7.5⌋ = 7
```

Thus, we need something to have a count ≥ 7.
(I.e., in this example, if something has a count of 7, 8, 9, 10, 11, 12 or 13, then it is the majority item.)

Thus, in our example, `"Cherry"` has the majority.


## Usage

```
	
	// Our sequence.
	//
	// NOTE that we do NOT have to put the sequence into a slice.
	// We are just doing that in this example.
	//
	sequence := []string{"Apple", "Banana", "Chery", "Apple", "Apple"}


	// Initialize.
		mjrty := New()

	// Insert data.
	//
	// While this is happening (if it actually exists for this sequence) it
	// will find the majority item.
	//
		for _, datum := range test.stream {

			mjrty.Insert(datum)
		}

	// Get the majority item.
	//
	// In the case of this example, this will be: "Apple".
	//
		majority := mjrty.Majority()

```

## See Also

For more information on the **mjrty** algorithm see:

* http://www.cs.utexas.edu/~moore/best-ideas/mjrty/
* http://blogs.law.harvard.edu/pamphlet/2011/09/23/tales-of-peer-review-episode-1-boyer-and-moores-mjrty-algorithm/
