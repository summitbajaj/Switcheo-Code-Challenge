// Exercise 4: Sum to n
// 1. sum_to_n_a() uses a for loop to sum to n. The time complexity for this is O(N).
// 2. sum_to_n_b() uses a recursion to sum to n. The time complexity for this is O(N).
// 3. sum_to_n_c() uses the mathematical approach to sum to n. The time complexity for this is O(1).

// Looking at the time complexity, approach 3 is the most efficient. 
// While approach 1 and 2 have the same time complexity, approach 1 might be better. 
// Approach 2 uses recursion, hence the stack memory might get filled up 
// and cause a stack overflow error. Nonetheless, 
// for this to happen a large amount of recursion calls is required.

// Assuming this input will always produce a result lesser than Number.MAX_SAFE_INTEGER.

func sum_to_n_a(n int) int {
	// your code here
	// running a for loop to sum to n
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func sum_to_n_b(n int) int {
	// your code here
	// using a recusrion to sum to n
	if n == 1 {
		return 1
	}
	return n + sum_to_n_b(n-1)
}

func sum_to_n_c(n int) int {
	// your code here
	// using the mathematical approach to summing to n: sum = n/2 *(n+1)
	return (n * (n + 1)) / 2
}
