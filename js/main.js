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
const wordList = fs.readFileSync(constants.DATA_PATH, "utf-8").split("\r\n").slice(2).map((value) => new Word(value.split(" ")[1]))
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

// get name
for (let i = 0; i < game.getNumberOfPlayers(); i++) {
    const newPlayer = new Player()
    newPlayer.setName(readline.question(`Player number ${i+1} name: `).trim()) // input player's name
    game.addPlayer(newPlayer)
}


while (! game.isNoMoreCardsLeft()) {
    game.nextPlayer()
    game.showTurn()
    console.log("Please pick one card")
    const pickedCard = playingPile.randomPick(1)[0]
    console.log("Card picked")
    pickedCard.show()
    console.log(`Now pick a number from 1 to ${pickedCard.numberOfWords}`)
    let wordIndex = Number(readline.question("").trim()) - 1
    var word = new Word(pickedCard.getWords()[wordIndex].getValue())
    // console.log(word)
    for (let i = 0; i < game.numberOfPlayers - 1; i++) {
        game.nextPlayer()
        game.showTurn()
        console.log(`The word to guess: ${word.value}`)
        console.log("Please write down the hint")
        game.getPlayerActive().easel = readline.question("").trim()
    }
    game.nextPlayer()
    game.showTurn()
    console.log("Here are the hints:")
    game.showHints(word)
    console.log("Write down your guess")
    console.log("If you don't want to guess, just tap PASS TURN")
    let guess = readline.question("").trim()
    if (guess == "PASS TURN") {
        continue
    }
    if (guess == word.value) {
        console.log("You have a good guess")
        console.log("You win one point")
        game.addPoints(1)
    } else {
        console.log("You have a wrong guess")
        console.log("You lost one another card")
        playingPile.randomPick(1)
    }
    console.log(`The correct answer is ${word.value}`)
}

game.showStatus()
game.log("JUST ONE")
fs.writeFile(constants.RECORD_PATH, game.getRecord(), () => {
    console.log("Record save")
})














