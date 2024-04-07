# Capital Quiz Game

Welcome to the Capital Quiz Game! This is a simple command-line quiz game where players are quizzed on country capitals.

## Usage

To play the game, run the executable and follow the prompts. By default, the game will start with a set of questions. You can also specify the number of questions using the `-q` option.

```
./quizgame            # Start the game with default settings
./quizgame -q 10      # Set the number of questions to 10
```

## Options

- `-h`: Display help message.
- `-q <length>`: Set the number of questions to `<length>`.

## Gameplay

- You will be presented with a country name.
- Choose the correct capital from the provided options.
- Answer as many questions as you can within the specified length.

## Exiting the Game

- You can exit the game at any time by typing `exit`, `q`, or `quit`.

## Requirements

- Go 1.13 or higher
- CountryCapital.csv file containing country-capital pairs

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, feel free to open an issue or create a pull request.
