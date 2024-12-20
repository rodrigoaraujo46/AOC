package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Registers struct {
	A   int
	B   int
	C   int
	IP  int
	Out string
}

type Instruction struct {
	Opcode  int
	Operand int
}

type Program []int

func part2(program []int, ans int) int {
	if len(program) == 0 {
		return ans
	}
	for b := range 8 {
		a := ans<<3 | b
		b = b ^ 6
		c := a >> b
		b = b ^ c
		b = b ^ 7
		if b%8 == program[len(program)-1] {
			sub := part2(program[:len(program)-1], a)
			if sub == 0 {
				continue
			}
			return sub
		}
	}
	return 0
}

func (r Registers) GetComboFromOperand(operand int) int {
	if operand == 4 {
		return r.A
	}
	if operand == 5 {
		return r.B
	}
	if operand == 6 {
		return r.C
	}

	if operand >= 0 && operand <= 3 {
		return operand
	}

	panic(fmt.Sprintf("Error: Operand %d is not recognized.", operand))
}

func adv(operand int, r *Registers) {
	exponent := operand
	if operand < 0 || operand > 3 {
		exponent = r.GetComboFromOperand(operand)
	}

	denominator := int(math.Pow(2, float64(exponent)))

	r.A = r.A / denominator
	r.IP += 2
}

func bxl(operand int, r *Registers) {

	bitwiseXOR := r.B ^ operand

	r.B = bitwiseXOR
	r.IP += 2
}

func bst(operand int, r *Registers) {
	dividend := operand
	if operand < 0 || operand > 3 {
		dividend = r.GetComboFromOperand(operand)
	}

	r.B = dividend % 8
	r.IP += 2
}

func jnz(operand int, r *Registers) {
	if r.A == 0 {
		r.IP += 2
		return
	}

	r.IP = operand
}

func bxc(r *Registers) {
	bitwiseXOR := r.B ^ r.C
	r.B = bitwiseXOR
	r.IP += 2
}

func out(operand int, r *Registers) {
	dividend := operand
	if operand < 0 || operand > 3 {
		dividend = r.GetComboFromOperand(operand)
	}

	quotient := dividend % 8

	if r.Out != "" {
		r.Out += ","
	}
	r.Out += strconv.Itoa(quotient)

	r.IP += 2
}

func bdv(operand int, r *Registers) {
	exponent := operand
	if operand < 0 || operand > 3 {
		exponent = r.GetComboFromOperand(operand)
	}

	denominator := int(math.Pow(2, float64(exponent)))

	r.B = r.A / denominator
	r.IP += 2
}

func cdv(operand int, r *Registers) {
	exponent := operand
	if operand < 0 || operand > 3 {
		exponent = r.GetComboFromOperand(operand)
	}

	denominator := int(math.Pow(2, float64(exponent)))

	r.C = r.A / denominator
	r.IP += 2
}

func (i Instruction) DispatchInstruction(r *Registers) {
	switch i.Opcode {
	case 0:
		adv(i.Operand, r)
	case 1:
		bxl(i.Operand, r)
	case 2:
		bst(i.Operand, r)
	case 3:
		jnz(i.Operand, r)
	case 4:
		bxc(r)
	case 5:
		out(i.Operand, r)
	case 6:
		bdv(i.Operand, r)
	case 7:
		cdv(i.Operand, r)
	default:
		panic(fmt.Sprintf("Error: Opcode %d is not recognized.", i.Opcode))
	}
}

func part1(registers *Registers, program Program) string {
	for registers.IP >= 0 && registers.IP < len(program) {
		instruction := Instruction{
			Opcode:  program[registers.IP],
			Operand: program[registers.IP+1],
		}
		instruction.DispatchInstruction(registers)
	}

	return registers.Out
}

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	//	lines = lines[:len(lines)-1]

	a, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	b, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	c, _ := strconv.Atoi(strings.Split(lines[2], ": ")[1])
	registers := &Registers{A: a, B: b, C: c}

	nums := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	program := make(Program, len(nums))
	for i := range nums {
		bit, _ := strconv.Atoi(nums[i])
		program[i] = bit
	}

	if len(os.Args) < 3 {
		fmt.Printf("PART 1: %v\n", part1(registers, program))
		fmt.Printf("PART 2: %v\n", part2(program, 0))
		return
	}
	if os.Args[2] == "2" {
		fmt.Printf("PART 2: %v", part2(program, 0))
		return
	}
	fmt.Printf("PART 1: %v", part1(registers, program))
}
