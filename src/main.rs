#![allow(dead_code)]
mod interpreter;
mod language_server;

use clap::Subcommand;
use clap::Parser;

use language_server::start_server;

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Cli {
    #[command(subcommand)]
    command: Option<Commands>,
}

#[derive(Subcommand, Debug)]
enum Commands {
    Run {
        filepath: String,
    },
    Ls {
        root: Option<String>,
    },
    Repl {},
}

fn main() {
    let cli = Cli::parse();

    match &cli.command {
        Some(Commands::Run { filepath }) => { println!("Running file: {}", filepath) }
        Some(Commands::Ls { root }) => { start_server(root.clone())}
        Some(Commands::Repl {}) => { println!("Starting REPL") }
        None => { println!("Command not recognized") }
    }
}
