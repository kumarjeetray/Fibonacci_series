package main
import "fmt"
 
func main() {
	var n int
	t1:= 0
	t2:= 1
	var nextTerm int
    fmt.Print("Enter the number of terms: ");
    fmt.Scan(&n);
    fmt.Print("Fibonacci Series: ");
    for i:= 1; i <= n; i++{
		fmt.Print(t1)
		if i==n{
			break
		}
		fmt.Print(",")
        nextTerm=t1+t2;
        t1 = t2;
        t2 = nextTerm;
    }
}