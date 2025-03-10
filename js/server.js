const net = require("net")
const fs = require("fs")
const readline = require("readline-sync")

const Pile = require("./classes/Pile")
const Card = require("./classes/Card")
const Word = require("./classes/Word")
const Game = require("./classes/Game")
const Player = require("./classes/Player")

const constants = require("./constant")

const game = new Game(constants.NUMBER_OF_PLAYERS)
const pile = new Pile()

// read file and return an array of Word objects
const wordList = fs.readFileSync(constants.DATA_PATH, "utf-8").split("\n").slice(2).map((value) => new Word((value.split(" ")[1]).split("\r")[0]))
// Take every n words to form a card
for (let i = 0; i < wordList.length / constants.NUMBER_OF_WORDS_PER_CARD; i++) {
    const new_card = new Card(constants.NUMBER_OF_WORDS_PER_CARD) // create a new card
    new_card.updateCard(wordList.slice(i * 5, i * 5 + 5)) // Add words to card
    pile.addCard(new_card) // Add card to main pile
}
// Pick a playing Pile
const playingPile = new Pile()
playingPile.updateCards(pile.randomPick(constants.NUMBER_OF_CARDS_PER_GAME))

game.addPile(playingPile) // Add the playing pile to our game

// Global variables
var connections = 0
const sockets = []
var hints = []
var hint_string = ""
var result = ""
var picked = null
var pickedWord = null

// TCP
const tcpServer = net.createServer((socket) => {
    connections += 1
    // console.log(socket)
    socket.write("hello")
    if (sockets.length < game.getNumberOfPlayers()) {
        sockets.push(socket)
        socket.on("data", (data) => {
            const cmd = data.toString().split(" ")
                switch (cmd[0]) {
                    case "name":
                        const newPlayer = new Player()
                        newPlayer.setName(cmd[1])
                        game.addPlayer(newPlayer)
                        if (connections == game.getNumberOfPlayers()) {
                            console.log("We have enough players. Let's begin the game")
                            game.nextPlayer()                 
                            sockets.forEach((value, index) => {
                                if (index == game.getPlayerActiveIndex()) {
                                    value.write(`pick ${game.points} ${game.playingPile.getNumberOfCards()}`)
                                } else {
                                    value.write(`wait_pick ${game.getPlayerActive().getName()} ${game.points} ${game.playingPile.getNumberOfCards()}`)
                                }
                            })
                        }
                        break

                    case "pick":
                        if (! game.isNoMoreCardsLeft()) {
                            hints = []
                            hint_string = ""
                            console.log("pick")
                            picked = playingPile.randomPick(1)[0]
                            sockets.forEach((socket, index) => {
                                if (index == game.getPlayerActiveIndex()) {
                                    socket.write(`choose ${game.points} ${game.playingPile.getNumberOfCards()}`)
                                } else {
                                    socket.write(`wait_choose ${game.getPlayerActive().getName()} ${game.points} ${game.playingPile.getNumberOfCards()} ${picked.toString()}`)
                                }
                                // socket.write("start", "utf-8")
                            })
                        } else {
                            sockets.forEach((socket, index) => {
                                if (index == game.getPlayerActiveIndex()) {
                                    socket.write("close")
                                } else {
                                    socket.write("close")
                                }
                            })
                        }
                        break

                    case "choose":
                        pickedWord = new Word(picked.getWords()[Number(cmd[1]) - 1].getValue())
                        sockets.forEach((socket, index) => {
                            if (index == game.getPlayerActiveIndex()) {
                                socket.write(`wait_hint ${game.points} ${game.playingPile.getNumberOfCards()}`)
                            } else {
                                socket.write(`hint ${game.getPlayerActive().getName()} ${game.points} ${game.playingPile.getNumberOfCards()} ${pickedWord.getValue()}`)
                            }
                        })
                        break

                    case "hint":
                        hints.push(cmd[1])
                        hint_string += "," + cmd[1]
                        if (hints.length == game.getNumberOfPlayers() - 1) {
                            sockets.forEach((socket, index) => {
                                if (index == game.getPlayerActiveIndex()) {
                                    socket.write(`guess ${game.points} ${game.playingPile.getNumberOfCards()} ${hint_string}`)
                                } else {
                                    socket.write(`wait_guess ${game.getPlayerActive().getName()} ${game.points} ${game.playingPile.getNumberOfCards()} ${pickedWord.getValue()}`)
                                }
                            })
                        }
                        break

                    case "guess":
                        if (cmd[1] == "PASS TURN") {
                            result = "pass"
                        } else if (cmd[1] == pickedWord.getValue()) {
                            game.addPoints(1)
                            result = "right"
                        } else {
                            playingPile.randomPick(1)
                            result = "wrong"
                        }
                        sockets.forEach((socket, index) => {
                            if (index == game.getPlayerActiveIndex()) {
                                socket.write(`result ${result} ${pickedWord.getValue()}`)
                            } else {
                                socket.write(`other_result ${game.getPlayerActive().getName()} ${result}`)
                            }
                        })

                        game.nextPlayer()
                        break
                }
        })
    }

})



tcpServer.listen(constants.PORT, constants.HOST, () => {
    console.log(`Listening at http://${tcpServer.address().address}:${tcpServer.address().port}`)
})

