package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func solve(A *treeNode) int {
	if A == nil {
		return 0
	}
	ans := 0
	if A.left != nil && A.left.left == nil && A.left.right == nil {
		ans += A.left.value
	} else {
		ans += solve(A.left)
	}
	ans += solve(A.right)
	return ans
}

// func main() {
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Go Away")
// 	})

// 	err := app.Listen("0.0.0.0:808")
// 	if err != nil {
// 		panic(err)
// 	}
// }

// [{()}]
type LB struct {
	a int
}

const (
	TotalEmails     = 1000000
	NumWorkers      = 100
	EmailsPerWorker = TotalEmails / NumWorkers
)

type Email struct {
	ID   int
	Body string
}

func sendEmails(emails <-chan Email, wg *sync.WaitGroup) {
	defer wg.Done()

	for email := range emails {
		// Simulate sending the email
		fmt.Printf("Sending email %d...\n", email.ID)
		// time.Sleep(time.Millisecond) // Simulated sending delay

		// Do additional processing or logging if needed
		// ...
	}
}
func single() {
	for i := 0; i < 1000000; i++ {
		// time.Sleep(time.Millisecond)
		fmt.Printf("Sending email %d...\n", i)

	}
}
func dNums(A int) int {
	count := 0
	for A > 0 {
		if A%2 != 0 {
			count++
		}
		A = A / 2
	}
	return count
}

// func main() {
// 	fmt.Println(dNums(4))
// Start the clock
// startTime := time.Now()
// // single()

// // Create a channel to send emails to workers
// emails := make(chan Email, TotalEmails)

// // Create a wait group to synchronize the goroutines
// var wg sync.WaitGroup
// wg.Add(NumWorkers)

// // Start worker goroutines
// for i := 0; i < NumWorkers; i++ {
// 	go sendEmails(emails, &wg)
// }

// // Generate emails and send them to workers
// for i := 0; i < TotalEmails; i++ {
// 	email := Email{
// 		ID:   i,
// 		Body: fmt.Sprintf("Hello, this is email #%d", i),
// 	}
// 	emails <- email
// }

// // Close the emails channel to indicate no more emails will be sent
// close(emails)

// // Wait for all workers to finish
// wg.Wait()

// // Calculate and print the elapsed time
// elapsed := time.Since(startTime)
// fmt.Printf("All emails sent. Elapsed time: %s\n", elapsed)
// }
func rev(x int) int {
	s := []int{}
	for x > 0 {
		s = append(s, x%10)
		x = x / 10
	}
	for x < 0 {
		// x = x * -1
		s = append(s, x%10)
		x = x / 10
	}

	ans := 0
	for i := 0; i < len(s); i++ {
		ans = ans*10 + s[i]
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}
	return ans
}

func hashs() {
	// Replace these values with your actual data
	razorpayPaymentID := "pay_MINFlmx5F8G0me"
	razorpayOrderID := "order_MINFZKUs3hIHy4"
	secret := "zXPPIue7Kv2sz5tCy12rVrem"

	message := fmt.Sprintf("%s|%s", razorpayOrderID, razorpayPaymentID)
	secretKey := []byte(secret)

	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(message))

	// Get the HMAC-SHA256 signature as a byte slice
	signature := h.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	generatedSignature := hex.EncodeToString(signature)

	fmt.Println(generatedSignature)
}
func x(A [][]int, B int) int {
	x := math.MaxInt
	i := 0
	j := len(A[0]) - 1
	for i < len(A)-1 && j >= 0 {
		if A[i][j] == B {
			x = min(x, (i+1)*1009+(j+1))
		}
		if A[i][j] > B {
			j--
		} else {
			i++
		}

	}
	if x == math.MaxInt {
		return -1
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func selsort(A []int, k int) int {
	for i := 0; i < k; i++ {
		k := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[k] {
				k = j
			}
		}
		A[i], A[k] = A[k], A[i]
		fmt.Println(A)
	}
	return A[k-1]
}
func solveds(A [][]int, B []int, C []int, D []int, E []int) []int {
	for i := 0; i < len(A); i++ {
		for j := 1; j < len(A[0]); j++ {
			A[i][j] += A[i][j-1]
		}
	}
	for i := 0; i < len(A); i++ {
		for j := 1; j < len(A[0]); j++ {
			A[j][i] += A[j-1][i]
		}
	}
	mod := int(1e9) + 7
	out := make([]int, len(B))
	for i := 0; i < len(B); i++ {
		a, b, c, d := B[i]-1, C[i]-1, D[i]-1, E[i]-1
		pf := A[c][d]
		if a > 0 {
			pf -= A[a-1][d]
		}
		if b > 0 {
			pf -= A[c][b-1]
		}
		if a > 0 && b > 0 {
			pf += A[a-1][b-1]
		}
		out[i] = (pf + mod) % mod
	}

	return out
}
func solvesd(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		for j := 1; j < len(A[i]); j++ {
			A[i][j] += A[i][j-1]
		}
	}
	for i := 0; i < len(A); i++ {
		for j := 1; j < len(A[i]); j++ {
			A[j][i] += A[j-1][i]
		}
	}
	return A
}
func flip(A string) []int {
	max, s, k, j := 0, 0, 0, 0
	for i := 0; i < len(A); i++ {
		if A[i] == '0' {
			s++
			j = i
			continue
		}
		if s > max {
			max = s
			k = j
		}
		s = 0
	}
	if max == 0 {
		return []int{}
	}
	return []int{k - max + 2, k + 1}
}
func plusOne(A []int) []int {
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == 9 {
			A[i] = 0
			continue
		}
		A[i]++
		return removeZero(A)
	}
	A[0] = 1
	for i := 1; i < len(A); i++ {
		A[i] = 0
	}
	A = append(A, 0)
	return A
}
func removeZero(a []int) []int {
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			break
		}
		a = a[i+1:]
		i--
	}
	return a
}
func maxSubArray(A []int) int {
	max := math.MinInt
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
func (l LB) print() {
	fmt.Println(l.a)
}

var lb LB = LB{a: 1}

func s() {
	j := struct {
		a int
	}{
		a: 1,
	}
	c := &j
	fmt.Println(c)
	c.a = 2
	fmt.Println(c)

	fmt.Println(j)

}
func solved(A []string, B string) int {
	// math.Pow()
	hash := map[int]string{}
	num := 0
	for i, v := range B {
		hash[i] = string(v)
	}
	for _, k := range A {
		if string(k[0]) == hash[num] {
			continue
		}
		num++
	}
	return 1
}

func helperz(A int) {
	if A == 0 {
		return
	}
	helperz(A - 1)
	fmt.Print(A, " ")
}
func add(A int) int {
	if A == 0 {
		return 0
	}
	return add(A-1) + A

}
func printn(A string) bool {
	return helper(A, 0, len(A)-1)
}
func helper(A string, i, j int) bool {
	if i >= j {
		return true
	}
	if A[i] != A[j] {
		return false
	}
	return helper(A, i+1, j-1)
}

func solve1(A []int, B int) int {
	count := 0
	hash := map[int]int{}
	for _, v := range A {
		need := B - v
		_, ok := hash[need]
		if !ok {
			hash[v] = 1
		} else {
			hash[v]++
			count += hash[need]
		}
	}
	return count
}
func solvess(A []int, B int) int {
	pre, suf := make([]int, len(A)), make([]int, len(A))
	pre[0] = A[0]
	for i := 1; i < len(A); i++ {
		pre[i] = pre[i-1] + A[i]
	}
	suf[len(A)-1] = A[len(A)-1]
	for i := len(A) - 2; i >= 0; i-- {
		suf[i] = suf[i+1] + A[i]
	}
	sum := Maxx(suf[len(A)-1], pre[0])
	for i := 1; i < B; i++ {
		s := Maxx(pre[i], suf[len(A)-(B+i)])
		sum = Maxx(sum, s)
	}
	return sum
}

func Maxx(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func solveX(A []int) int {
	hash := make(map[int]int)
	for _, v := range A {
		_, ok := hash[v]
		if !ok {
			hash[v] = 1
		} else {
			hash[v]++
		}
	}
	for _, v := range A {
		count, ok := hash[v]
		if ok {
			if count >= 2 {
				return v
			}
		}
	}
	return -1
}
func solvee(A []int, B []int) []int {
	hash1, hash2 := make(map[int]int), make(map[int]int)
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		_, ok := hash1[A[i]]
		if !ok {
			hash1[A[i]] = 1
		} else {
			hash1[A[i]]++
		}
		i++
		_, ok = hash2[B[j]]
		if !ok {
			hash2[B[j]] = 1
		} else {
			hash2[B[j]]++
		}
		j++
	}
	for j < len(B) {
		_, ok := hash2[B[j]]
		if !ok {
			hash2[B[j]] = 1
		} else {
			hash2[B[j]]++
		}
		j++
	}
	for i < len(A) {
		_, ok := hash1[A[i]]
		if !ok {
			hash1[A[i]] = 1
		} else {
			hash1[A[i]]++
		}
		i++
	}
	out := []int{}
	if len(hash1) > len(hash2) {
		for i, v := range hash1 {
			o, ok := hash2[i]
			if ok {
				if o >= v {
					for j := 0; j < v; j++ {
						out = append(out, i)
					}
				} else {
					for j := 0; j < o; j++ {
						out = append(out, i)
					}
				}
			}
		}
	} else {
		for i, v := range hash2 {
			o, ok := hash1[i]
			if ok {
				if o >= v {
					for j := 0; j < v; j++ {
						out = append(out, i)
					}
				} else {
					for j := 0; j < o; j++ {
						out = append(out, i)
					}
				}
			}
		}
	}

	return out
}

func addBinary(A []int) int {
	count := 0
	if len(A) == 1 {
		return 1
	}
	pre := make([]int64, len(A))
	if A[0] == 0 {
		count++
	}
	pre[0] = int64(A[0])
	for i := 1; i < len(A); i++ {
		pre[i] = int64(A[i] + int(pre[i-1]))
		if pre[i] == 0 {
			count++
		}
	}
	hash := map[int64]int{}
	for i := 0; i < len(A); i++ {
		_, ok := hash[int64(pre[i])]
		if !ok {
			hash[int64(pre[i])] = 1
		} else {
			hash[int64(pre[i])]++
		}
	}
	for _, v := range hash {
		if v > 1 {
			count = ((count % (1000000000 + 7)) + ((v*(v-1))/2)%(1000000000+7)) % (1000000000 + 7)
		}
	}
	return count
}

func find(a, b, c byte) (string, string) {
	if a == '1' && b == '1' && c == '1' {
		return "1", "1"
	} else if a == '0' && b == '0' && c == '0' {
		return "0", "0"
	} else {
		if (a == '1' && b == '1') || (b == '1' && c == '1') || (a == '1' && c == '1') {
			return "0", "1"
		}
	}
	// delete()
	return "1", "0"
}

func Contains(s string) bool {
	w := []string{"a", "e", "i", "o", "u"}
	for _, v := range w {
		if v == s {
			return true
		}
	}
	return false
}
func longestCommonPrefix(A []string) string {
	l := 9999999999999
	for i := 0; i < len(A); i++ {
		if len(A[i]) < l {
			l = len(A[i])
		}
	}
	c := A[0]
	for i := 0; i < len(A)-1; i++ {
		d := findCommonPrefix(A[i], A[i+1])
		if len(c) > len(d) {
			c = d
		}
	}
	return c
}

func findCommonPrefix(str1, str2 string) string {
	var prefix string
	// Compare characters until a mismatch is found
	i := 0
	for i < len(str1) && i < len(str2) && str1[i] == str2[i] {
		prefix += string(str1[i])
		i++
	}

	return prefix
}
func sorts(A []int) []int {
	mz := A[0]
	for i := 0; i < len(A); i++ {
		if A[i] > mz {
			mz = A[i]
		}
	}
	L := make([]int, mz+1)
	for i := 0; i < len(A); i++ {
		L[A[i]]++
	}
	j := 0
	for i := 0; i < len(L); i++ {
		if L[i] > 0 {
			for k := 0; k < L[i]; k++ {
				A[j] = i
				j++
			}
		}
	}
	return A
}
func longestPalindrome(A string) string {
	if len(A) == 0 {
		return ""
	}
	s := A[0:1]
	for i := 0; i < len(A)-1; i++ {
		d := checkPalindrom(A, i, i)
		if len(s) <= len(d) {
			s = d
		}
		d = checkPalindrom(A, i, i+1)
		if len(s) <= len(d) {
			s = d
		}
	}
	return s
}

func checkPalindrom(s string, a, b int) string {
	for a >= 0 && b < len(s) && s[a] == s[b] {
		a--
		b++
	}
	return s[a+1 : b]
}

// var a rune = 'a'
func alter(s string) string {
	x := []rune(s)
	for i := 0; i < len(s); i++ {
		if s[i] <= 97 {
			x[i] = rune(s[i] + 32)

		} else {
			x[i] = rune(s[i] - 32)
		}
	}
	return string(x)
}

func solvex(A string) int {
	// prefix arrayy of sum of 1's
	prefix := make([]int, len(A))
	prefix[0] = 0
	if string(A[0]) == "1" {
		prefix[0] = 1
	}
	for i := 1; i < len(A); i++ {
		if string(A[i]) == "1" {
			prefix[i] = prefix[i-1] + 1
		} else {
			prefix[i] = 0
		}
	}
	// suffix array of sum of 1's
	suffix := make([]int, len(A))
	suffix[len(A)-1] = 0
	if string(A[len(A)-1]) == "1" {
		suffix[len(A)-1] = 1
	}
	for i := len(A) - 2; i >= 0; i-- {
		if string(A[i]) == "1" {
			suffix[i] = suffix[i+1] + 1
		} else {
			suffix[i] = 0
		}
	}

	maxOne, tempSum := 0, 0
	for i := 0; i < len(A); i++ {
		if string(A[i]) == "0" {
			if i == 0 {
				tempSum = suffix[i]
			} else if i == len(A)-1 {
				tempSum = prefix[i]
			} else {
				tempSum = prefix[i-1] + suffix[i+1]
			}
			if tempSum > 0 {
				if tempSum+1 > maxOne {
					maxOne = tempSum + 1
				}
			}
		}
	}
	if strings.Contains(A, "1") && maxOne == 0 {
		return len(A)
	}
	return maxOne
}

func Maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countFactors(n int) int {
	count := 0
	sqrtN := int(math.Sqrt(float64(n)))

	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			// Increment the count for factors i and n/i
			count += 2
		}
	}

	// If the number is a perfect square, decrement the count by 1
	if sqrtN*sqrtN == n {
		count--
	}

	return count
}

func solveb(A int, B int, C int) int {
	num, _ := strconv.Atoi(Sort(A, B, C))
	return num
}

func Sort(a, b, c int) string {
	if a >= b && a >= c {
		if b >= c {
			return fmt.Sprintf("%d%d%d", c, b, a)
		} else {
			return fmt.Sprintf("%d%d%d", b, c, a)
		}
	} else if b >= a && b >= c {
		if a >= c {
			return fmt.Sprintf("%d%d%d", c, a, b)
		} else {
			return fmt.Sprintf("%d%d%d", a, c, b)
		}
	} else {
		if a >= b {
			return fmt.Sprintf("%d%d%d", b, a, c)
		} else {
			return fmt.Sprintf("%d%d%d", a, b, c)
		}
	}
}

func solvea(A int, B int) int {
	A, B = max(A, B)
	d := A - B
	c := 1
	c = c << (d + 1)
	c = c | 1
	c = c << (B - 1)
	return c
}
func max(a, b int) (int, int) {
	temp, sec := 0, 0
	if a >= b {
		temp = a
		sec = b
	} else {
		temp = b
		sec = a
	}
	return temp, sec
}
func numSetBits(A int) (count int) {
	for A > 0 {
		if A&1 == 1 {
			count++
		}
		A = A >> 1
	}
	return
}

func decimalToBinary(num int) string {
	if num == 0 {
		return "0"
	}

	binary := ""
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		binary += strconv.Itoa(bit)
	}

	return binary
}

func solves(s string, B [][]int) bool {
	val := false
	for i := 0; i < len(s); i++ {
		var e string
		if string(s[i]) == "(" {
			e = ")"
		} else if string(s[i]) == "{" {
			e = "}"
		} else {
			e = "]"
		}
		F, S, P := 0, 0, 0
		for j := i + 1; j < len(s); j++ {
			if string(s[j]) == e && F == 0 && S == 0 && P == 0 {
				val = true
			} else if string(s[j]) == "(" {
				P++
			} else if string(s[j]) == "{" {
				F++
			} else if string(s[j]) == "[" {
				S++
			} else if string(s[j]) == ")" && P >= 1 {
				P--
			} else if string(s[j]) == "}" && F >= 1 {
				F--
			} else if string(s[j]) == "]" && S >= 1 {
				S--
			} else {
				return false
			}
		}
	}
	return val
}
