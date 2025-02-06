class Card {
    constructor(numberOfWords) {
        this.numberOfWords = numberOfWords
        this.words = new Array() //array of Word object
    }
    
    addWord(word) {
        if (this.words.length < this.numberOfWords) {
            this.words.push(word)
            return true
        } else {
            console.log("This card has enough words")
            return false
        }
    }

    updateCard(wordList) {
        this.words = wordList
    }

    getWords() {
        return this.words
    }

    show() {
        console.log("+============+")
        this.words.forEach((value, index) => {
            console.log(`+ ${index + 1}. ${value.value} +`)
        })
        console.log("+============+")
    }
}

module.exports = Card