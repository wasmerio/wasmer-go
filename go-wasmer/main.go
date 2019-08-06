package main

import (
	"fmt"
	"github.com/urfave/cli"
	wasm "github.com/spacemeshos/go-ext-wasm/wasmer"
	"log"
	"os"
	"strconv"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-wasmer"
	app.Usage = "Run WebAssembly binaries."
	app.Version = "0.2.0"

	app.Commands = []cli.Command{
		{
			Name:    "call",
			Aliases: []string{},
			Usage:   "Call a WebAssembly exported function with arguments: <wasm-file> <function-name> <arguments...>",
			Action: func(c *cli.Context) error {
				if c.NArg() < 2 {
					return cli.NewExitError(fmt.Sprintf("`call` expects at least 2 arguments: <wasm-file> and <function-name>; given %d.", c.NArg()), 1)
				}

				arguments := c.Args()
				wasmFile := arguments[0]
				functionName := arguments[1]
				functionArguments := arguments[2:]

				bytes, err := wasm.ReadBytes(wasmFile)

				if err != nil {
					return cli.NewExitError(fmt.Sprintf("`call` — Failed to read bytes from `%s`: %s.", wasmFile, err), 2)
				}

				instance, err := wasm.NewInstance(bytes)

				if err != nil {
					return cli.NewExitError(fmt.Sprintf("`call` — Failed to instantiate `%s`: %s.", wasmFile, err), 3)
				}

				exportedFunction, exists := instance.Exports[functionName]

				if !exists {
					return cli.NewExitError(fmt.Sprintf("`call` — Exported function `%s` does not exist on the WebAssembly module `%s`.", functionName, wasmFile), 4)
				}

				exportedFunctionArguments := make([]interface{}, len(functionArguments))

				for nth, functionArgument := range functionArguments {
					parsed, err := strconv.ParseInt(functionArgument, 10, 32)

					if err != nil {
						return cli.NewExitError(fmt.Sprintf("`call` — Failed to parse the #%d argument `%s` of the exported function `%s` of the WebAssembly module `%s`: %s.", nth+1, functionArgument, functionName, wasmFile, err), 5)
					}

					exportedFunctionArguments[nth] = (int32)(parsed)
				}

				result, err := exportedFunction(exportedFunctionArguments...)

				if err != nil {
					return cli.NewExitError(fmt.Sprintf("`call` — Exported function `%s` of the WebAssembly module `%s` failed: %s.", functionName, wasmFile, err), 6)
				}

				fmt.Println(result)

				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
