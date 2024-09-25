mod interpreter;
mod language_server;

use clap::Subcommand;
use clap::Parser;
use std::path::PathBuf;

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
}

fn main() {
    let cli = Cli::parse();

    match &cli.command {
        Some(Commands::Run { filepath }) => { println!("Running file: {}", filepath) }
        Some(Commands::Ls { root }) => { start_server(root.clone())}
        None => { println!("Command not recognized") }
    }
}
