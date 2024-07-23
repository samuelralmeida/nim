# Nim Game AI

This software implements a command-line Nim game with an AI opponent that learns through reinforcement learning. The AI uses Q-learning to improve its strategy over time.

## Usage

To get started with the Nim Game AI, you'll need to have [Go](https://go.dev/) installed. Then, you can clone this repository and build the project:

```bash
git clone https://github.com/samuelralmeida/nim.git
cd nim
go build
```

You can run the Nim game with the AI opponent by using the following command:

```bash
./nim -alpha=0.5 -epsilon=0.1 -n=100000
```

### Available flags:
- alpha: The learning rate of the AI (default: 0.5)

- epsilon: The exploration rate of the AI (default: 0.1)

- n: The number of training iterations for the AI (default: 10000)

## Features
- Reinforcement Learning: The AI uses Q-learning to improve its strategy.

- Command-line Interface: Play the Nim game directly from the command line.

- Customizable Parameters: Adjust the learning rate, exploration rate, and number of training iterations to see how they affect the AI's performance.

## How it Works

The game is set up with the following initial configuration:

- Four piles of objects: [1, 3, 5, 7]
- Two players: human and AI

The AI learns by playing multiple games and updating its strategy based on the outcomes. The Q-learning algorithm helps the AI to balance between exploring new moves and exploiting known successful moves.

## Contributing

If you have any suggestions or find any issues, please feel free to create a pull request or open an issue.

---

Enjoy playing and improving the AI for the Nim game!