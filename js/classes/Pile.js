class Pile {
    constructor() {
        this.cards = new Array() // array of Card object
    }

    addCard(card) {
        this.cards.push(card)
    }

    pickCard(index) {
        return this.cards.splice(index, 1)[0]
    }

    randomPick(numberOfCards) {
        const picked = new Array()
        for (let i = 0; i < numberOfCards; i++) {
            picked.push(this.pickCard(Math.floor(Math.random() * this.cards.length)))
        }
        return picked
    }

    updateCards(cardList) {
        this.cards = cardList
    }

    isNoMoreCardsLeft() {
        return this.cards.length <= 0
    }

    getNumberOfCards() {
        return this.cards.length
    }
}

module.exports = Pile