package main

// Make sure to uncomment the packages you are using below
import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type API interface {
	GetNumBills() int
	GetBill(int) int
	GetPrice() int
	GetPayment() int
}

type api struct {
	bills    []int
	MAXPRICE int
}

func newapi(nbills, maxface, maxprice int) *api {
	bills := make([]int, nbills)
	mem := make(map[int]bool)
	n := 0
	fmt.Print("**")
	for n < nbills {
		face := rand.Intn(maxface) + 1
		println("face", face)
		_, ok := mem[face]
		if !ok {
			mem[face] = true
			bills[n] = face
			n += 1
		}
	}
	fmt.Println(bills)
	return &api{bills, maxprice}
}

func newAPI(nbills, maxface, maxprice int) API {
	return newapi(nbills, maxface, maxprice)
}

func (a *api) GetNumBills() int {
	return len(a.bills)
}

func (a *api) GetBill(index int) int {
	return a.bills[index]
}

func (a *api) GetPrice() int {
	return rand.Intn(a.MAXPRICE)
}

func (a *api) GetPayment() int {
	payment := 0
	for {
		payment = rand.Intn(a.MAXPRICE)
		if payment > 0 && payment < a.MAXPRICE {
			break
		}
	}
	return payment
}

type solution struct {
	API
}
type SolutionInterface interface {
	MinBills() int
	SmallestUnreachable() int
}

func NewSolution() SolutionInterface {
	// You can initiate and calculate things here
	return &solution{newAPI(5, 7, 200)}
}

/*
Task 1: Returns the minimum number of bills needed to give a customer their
change when purchasing an item.
*/
func (s *solution) MinBills() int {

	nbills := s.GetNumBills()
	fmt.Println("number of bills", nbills)
	bills := make([]int, nbills)
	for i := 0; i < nbills; i++ {
		bills[i] = s.GetBill(i)
	}
	sort.Ints(bills)
	fmt.Println("Bills", bills)
	price := s.GetPrice()
	fmt.Println("Price", price)
	payment := s.GetPayment()
	fmt.Println("payment", payment)
	diff := payment - price
	fmt.Println("diff", diff)

	res := 0
	for i := nbills - 1; i >= 0; i-- {
		fmt.Println("+", bills[i])
		for diff > bills[i] {
			diff -= bills[i]
			res += 1
		}
	}
	fmt.Println("res", res)
	return res

}

/*
Task 2: Returns the minimum amount unreachable using at most one bill of each
denomination.
*/
func (s *solution) SmallestUnreachable() int {
	// Write your code here
	powerSetHelper := func (original []int) [][]int {
		powerSetSize := int(math.Pow(2, float64(len(original))))
		result := make([][]int, 0, powerSetSize)

		var index int
		for index < powerSetSize {
		var subSet []int

		for j, elem := range original {
		if index& (1 << uint(j)) > 0 {
		subSet = append(subSet, elem)
	}
	}
		result = append(result, subSet)
		index++
	}
		return result
	}


	nbills := s.GetNumBills()
	fmt.Println("number of bills", nbills)
	bills := make([]int, nbills)
	for i := 0; i < nbills ; i++ {
		bills[i] = s.GetBill(i)
	}
	fmt.Println("Bills", bills)
	price := s.GetPrice()
	fmt.Println("Price", price)
	payment := s.GetPayment()
	fmt.Println("payment", payment)
	diff := payment - price
	fmt.Println("diff", diff)

	p := powerSetHelper(bills)
	max := 0
	maxIndex := -1
	for i,v := range p {
		sum := 0
		for _,b := range v {
			sum += b
		}
		if sum > max {
			max = sum
			maxIndex = i
		}
	}
	fmt.Println(p[maxIndex])

	return diff - max
}

func main() {
	s := NewSolution()
	minBills := s.MinBills()
	//Time complexity exponential becasue of PowerSet probably exists much better
	// and less naive solution then 2Powx, but I am weak in combinatoric optimization for now...
	smallestUn := s.SmallestUnreachable()
	fmt.Print("Minimum number of bills in return", minBills,smallestUn)
}
