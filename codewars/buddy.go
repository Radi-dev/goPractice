package kata
func SumDivisors(n int)int{
  s:=0
  for i := 1; i < n; i++ {
    if n%i==0 {
      s+=i
    }
  }
  return s
}
func Buddy(start, limit int) []int {
  for i := start; i <= limit; i++ {
    n :=i
    var m int = SumDivisors(n)-1
    if SumDivisors(m)-1==n && m>n {
    return []int{n,m}}
  }
  return nil
}