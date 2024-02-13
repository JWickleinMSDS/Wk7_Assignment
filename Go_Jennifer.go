package main

import (
	"github.com/dave/jennifer/jen"
)

func main() {
	file := jen.NewFile("main")
	file.Func().Id("main").Params().Block(
		jen.Var().Id("goTime").Qual("time", "Duration"),
		jen.Var().Id("rTime").Qual("time", "Duration"),
		jen.Var().Id("pythonTime").Qual("time", "Duration"),
		jen.List(jen.Id("goTime"), jen.Err()).Op("=").Id("readExecutionTime").Call(jen.Lit("Go_Execution_Time.txt")),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Qual("fmt", "Println").Call(jen.Lit("Error reading Go execution time:"), jen.Err()),
			jen.Return(),
		),
		jen.List(jen.Id("rTime"), jen.Err()).Op("=").Id("readExecutionTime").Call(jen.Lit("R_execution_time.txt")),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Qual("fmt", "Println").Call(jen.Lit("Error reading R execution time:"), jen.Err()),
			jen.Return(),
		),
		jen.List(jen.Id("pythonTime"), jen.Err()).Op("=").Id("readExecutionTime").Call(jen.Lit("Python_execution_time.txt")),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Qual("fmt", "Println").Call(jen.Lit("Error reading Python execution time:"), jen.Err()),
			jen.Return(),
		),
		jen.Qual("fmt", "Printf").Call(jen.Lit("Go Execution Time: %s\nR Execution Time: %s\nPython Execution Time: %s\n"), jen.Id("goTime"), jen.Id("rTime"), jen.Id("pythonTime")),
	)
	file.Func().Id("readExecutionTime").Params(jen.Id("filename").String()).Params(jen.Qual("time", "Duration"), jen.Error()).Block(
		jen.Var().Id("content").Index().Byte(),
		jen.Var().Id("err").Error(),
		jen.List(jen.Id("content"), jen.Id("err")).Op("=").Qual("io/ioutil", "ReadFile").Call(jen.Id("filename")),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Qual("time", "Duration").Call(jen.Lit(0)), jen.Id("err")),
		),
		jen.Var().Id("executionTimeStr").String().Op("=").Qual("strings", "TrimSpace").Call(jen.Qual("string", "content")),
		jen.If(jen.Op("!").Qual("strings", "ContainsAny").Call(jen.Id("executionTimeStr"), jen.Lit("smh"))).Block(
			jen.Id("executionTimeStr").Op("+=").Lit("ms"),
		),
		jen.Var().Id("executionTime").Qual("time", "Duration"),
		jen.List(jen.Id("executionTime"), jen.Id("err")).Op("=").Qual("time", "ParseDuration").Call(jen.Id("executionTimeStr")),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Qual("time", "Duration").Call(jen.Lit(0)), jen.Id("err")),
		),
		jen.Return(jen.Id("executionTime"), jen.Nil()),
	)

	// Generate the file
	err := file.Save("generated_program.go")
	if err != nil {
		panic(err)
	}
}
