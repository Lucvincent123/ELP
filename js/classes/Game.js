class Game {
    constructor(numberOfPlayers) {
        this.points = 0
        this.numberOfPlayers = numberOfPlayers
        this.players = new Array()
    }

    addPoints(points) {
        this.points += points
    }

    addPile(pile) {
        this.playingPile = pile
    }

    isNoMoreCardsLeft() {
        return this.playingPile.isNoMoreCardsLeft()
    }

    addPlayer(player) {
        this.players.push(player)
    }

    isFullPlayer() {
        return this.players.length >= this.numberOfPlayers
    }


    nextPlayer() {
        if (this.playerActive != undefined) {
            this.playerActive = (this.playerActive + 1) % this.numberOfPlayers
        } else {
            this.playerActive = 0
        }

    }

    getPlayerActive() {
        return this.players[this.playerActive]
    }

    showHints() {
        this.players.forEach((player, index) => {
            if (index != this.playerActive) {
                console.log(`${player.getName()} : ${player.easel}`)
            }
        })
    }

    showStatus() {
        console.log(`Points : ${this.points}`)
        console.log(`Cards left : ${this.playingPile.getNumberOfCards()}`)
    }
}

module.exports = Game