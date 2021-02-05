// Compiled to Wasm as follows:
//
// ```sh
// $ rustc --target wasm32-wasi -O wasi.rs -o wasi.raw.wasm
// $ wasm-strip wasi.raw.wasm
// $ wasm-opt -O4 -Oz wasi.raw.wasm -o wasi.wasm
// ```

use std::{env, fs};

fn main() {
    // Arguments
    {
        let mut arguments = env::args().collect::<Vec<String>>();

        println!("Found program name: `{}`", arguments[0]);

        arguments = arguments[1..].to_vec();
        println!(
            "Found {} arguments: {}",
            arguments.len(),
            arguments.join(", ")
        );
    }

    // Environment variables
    {
        let mut environment_variables = env::vars()
            .map(|(arg, val)| format!("{}={}", arg, val))
            .collect::<Vec<String>>();
        environment_variables.sort();

        println!(
            "Found {} environment variables: {}",
            environment_variables.len(),
            environment_variables.join(", ")
        );
    }

    // Directories.
    {
        let root = fs::read_dir("/")
            .unwrap()
            .map(|e| e.map(|inner| format!("{:?}", inner)))
            .collect::<Result<Vec<String>, _>>()
            .unwrap();

        println!(
            "Found {} preopened directories: {}",
            root.len(),
            root.join(", ")
        );
    }
}
