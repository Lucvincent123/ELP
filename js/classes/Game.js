class Game {
    constructor(numberOfPlayers) {  // numberOfPlayer : Int
        this.numberOfPlayers = numberOfPlayers
        this.points = 0
        this.players = new Array() // Array of Player objects
        // this.playingPile : Pile object
        // this.playerActiveIndex : Index of the active player (type Int)
        this.record = "" // record of the game
    }

    addPoints(points) {  // points : Int
        this.points += points
    }

    addPile(pile) { // pile : Pile object
        this.playingPile = pile
    }

    isNoMoreCardsLeft() {
        return this.playingPile.getNumberOfCards() <= 0 // bool
    }

    addPlayer(player) {  // player : Player object
        if (this.players.length < this.numberOfPlayers) {
            this.players.push(player)
            return true
        } else {
            console.log("This game has enough players")
            return false
        }
    }

    isFullPlayer() {
        return this.players.length >= this.numberOfPlayers
    }

    nextPlayer() {
        if (this.playerActiveIndex != undefined) {
            this.playerActiveIndex = (this.playerActiveIndex + 1) % this.numberOfPlayers
        } else {
            this.playerActiveIndex = 0
        }
    }

    getPlayerActiveIndex() {
        return this.playerActiveIndex
    }

    getPlayerActive() {
        return this.players[this.playerActiveIndex]
    }

    isAlreadyWritten(hint, playerIndex) {
        for (let i = 0; i < this.getNumberOfPlayers(); i++) {
            if (playerIndex != i && hint == this.players[i].getEasel()) {
                return true
            }
        }
        return false
    }

    showHints(wordToGuess) {  // Word Object
        this.players.forEach((player, index) => {
            if (index != this.playerActiveIndex) {
                // Check if the hint is invalid               or            is already written by another player
                if (wordToGuess.isSameWord(player.getEasel()) || this.isAlreadyWritten(player.getEasel(), index)) { 
                    console.log(`${player.getName()} : **********`)
                } else {
                    console.log(`${player.getName()} : ${player.getEasel()}`)
                }
            }
        })
    }

    showStatus() {
        console.log(`Points : ${this.points}`)
        console.log(`Cards left : ${this.playingPile.getNumberOfCards()}`)
    }

    showTurn() {
        console.log("#####################################################")
        this.showStatus()
        console.log(`${this.getPlayerActive().getName()}'s turn`)
    }

    getNumberOfPlayers() {
        return this.numberOfPlayers
    }

    log(string) {
        this.record += string
    }

    getRecord() {
        return this.record
    }
}

module.exports = Game