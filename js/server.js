require("dotenv").config()

// const net = require("net")
const fs = require("fs")
const readline = require("readline-sync")

const Pile = require("./classes/Pile")
const Card = require("./classes/Card")
const Word = require("./classes/Word")
const Game = require("./classes/Game")
const Player = require("./classes/Player")

const game = new Game(process.env.NUMBER_OF_PLAYERS)
const pile = new Pile()
// read card data from file
// fs.readFile(process.env.FILE_PATH, "utf-8", (err, data) => {
//     let numberOfWordPerCard = process.env.NUMBER_OF_WORDS_PER_CARD
//     let wordList = data.split("\r\n").slice(2).map((value) => new Word(value.split(" ")[1])) // array of word objects
//     for (let i = 0; i < wordList.length / numberOfWordPerCard; i++) {
//         let new_card = new Card(numberOfWordPerCard)
//         new_card.updateCard(wordList.slice(i * 5, i * 5 + 5))
//         game.addCard(new_card)
//     }
// })
const wordList = fs.readFileSync(process.env.FILE_PATH, "utf-8").split("\r\n").slice(2).map((value) => new Word(value.split(" ")[1]))
for (let i = 0; i < wordList.length / process.env.NUMBER_OF_WORDS_PER_CARD; i++) {
    const new_card = new Card(process.env.NUMBER_OF_WORDS_PER_CARD)
    new_card.updateCard(wordList.slice(i * 5, i * 5 + 5))
    pile.addCard(new_card)
}
// Pick a playing Pile
const playingPile = new Pile()
playingPile.updateCards(pile.randomPick(process.env.NUMBER_OF_CARDS_PER_GAME))

game.addPile(playingPile)
console.log(game.numberOfPlayers)

// readline.question(`Player number ${i+1} name: `, (name) => {
//     const newPlayer = new Player()
//     newPlayer.setName(name.trim())
//     game.addPlayer(newPlayer)
//     readline.close()
// })

for (let i = 0; i < game.numberOfPlayers; i++) {
    const newPlayer = new Player()
    newPlayer.setName(readline.question(`Player number ${i+1} name: `).trim())
    game.addPlayer(newPlayer)
}


while (! game.isNoMoreCardsLeft()) {
    game.nextPlayer()
    console.log("#####################################################")
    game.showStatus()
    console.log(`${game.getPlayerActive().getName()}'s turn`)
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
        console.log("#####################################################")
        game.showStatus()
        console.log(`The word to guess: ${word.value}`)
        console.log(`${game.getPlayerActive().getName()}'s turn`)
        console.log("Please write down the hint")
        game.getPlayerActive().easel = readline.question("").trim()
    }
    game.nextPlayer()
    console.log("#####################################################")
    game.showStatus()
    console.log(`${game.getPlayerActive().getName()}'s turn`)
    console.log("Here are the hints:")
    game.showHints()
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
// playingPile.randomPick(1)[0]

// console.log(game.players)
















// function getPlayer() {
//     return new Promise((resolve) => {
//         const tcpServer = net.createServer((so))
//     })
// }


// const tcpServer = net.createServer((socket) => {
//     if (! game.isFullPlayer()) {
//         var newPlayer = new Player()
//         game.addPlayer(newPlayer)
//     } else {
//         console.log("We have enough players")
//         console.log("Let's begin")
//     }

//     socket.on("data", (data) => {
//         let cmd = data.split(" ")
//         switch(cmd[0]) {
//             case "name":
//                 newPlayer.setName(cmd[1])
//                 break
//         }
//     })
// })



// tcpServer.listen(process.env.PORT, process.env.HOST, () => {
//     console.log(`Listening at http://${tcpServer.address().address}:${tcpServer.address().port}`)
// })

