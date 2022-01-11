package main

import (
	"log"

	"github.com/wdevore/RISCV-Meta-Assembler/src"
)

func main() {
	run_assembler()
}

func run_assembler() {
	assembler, err := src.NewAssembler()

	if err != nil {
		log.Fatalln(err)
	}

	err = assembler.Configure(".")
	if err != nil {
		log.Fatalln(err)
	}

	props := assembler.Properties()
	log.Println("Generating output: " + props.BinaryName())

	log.Println("Assembling...")

	for _, source := range props.Files() {
		err = assembler.Run(source)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// assembler.Print()

	log.Println("Assembly done.")
}

// func test_expression() {
// 	// (* (- 123) (group 45.67)) = -123 * (45.67)
// 	// (* (- 123) (group 45.669998))

// 	// expression := interpreter.NewBinaryExpression(
// 	// 	interpreter.NewUnaryExpression(
// 	// 		scanner.NewToken(api.MINUS, "-", nil, 1),
// 	// 		interpreter.NewLiteralExpression(
// 	// 			literals.NewIntegerLiteral("123"),
// 	// 		),
// 	// 	),
// 	// 	scanner.NewToken(api.STAR, "*", nil, 1),
// 	// 	interpreter.NewGroupingExpression(
// 	// 		interpreter.NewLiteralExpression(
// 	// 			literals.NewNumberLiteral("45.67"),
// 	// 		),
// 	// 	),
// 	// )

// 	// astPrinter := interpreter.NewAstPrinter().(*interpreter.AstPrinter)
// 	// pretty := astPrinter.Print(expression)
// 	// fmt.Println(pretty)
// }
