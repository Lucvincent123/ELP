class Pile {
    constructor() {
        this.cards = new Array() // array of Card objects
    }

    addCard(card) {   // card : Card object
        this.cards.push(card)
    }

    pickCard(index) {  // index : Int
        return this.cards.splice(index, 1)[0]
    }

    randomPick(numberOfCards) { // numberOfCards : Int
        const picked = new Array()
        for (let i = 0; i < numberOfCards; i++) {
            // Pick a random card from the pile
            if (this.cards.length == 0) break
            picked.push(this.pickCard(Math.floor(Math.random() * this.cards.length)))
        }
        return picked
    }

    updateCards(cardList) {  // cardList : Array of Card objects
        this.cards = cardList
    }

    getNumberOfCards() {
        return this.cards.length  // Int
    }
}

module.exports = Pile