package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:  bufio.NewScanner(in),
		out: out,
		game: game,	
    }
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

    numberOfPlayersInput := cli.readline()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

    if err != nil {
        fmt.Fprint(cli.out, BadPlayerInputErrMsg)
        return
    }

    cli.game.Start(numberOfPlayers)

    winnerInput := cli.readline()
    winner := extractWinner(winnerInput)

    cli.game.Finish(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readline() string {
	cli.in.Scan()
	return cli.in.Text()
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}
