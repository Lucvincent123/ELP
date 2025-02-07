class Card {
    constructor(numberOfWords) {
        this.numberOfWords = numberOfWords // Int
        this.words = new Array() // array of Word object
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
        this.words = wordList // Word object
    }

    getWords() {
        return this.words // Array of Word object
    }

    show() {
        console.log("+============+")
        this.words.forEach((value, index) => {
            console.log(`+ ${index + 1}. ${value.value} +`)
        })
        console.log("+============+")
    }

    toString() {
        var msg = ""
        for (let i = 0; i < this.words.length; i++) {
            msg += this.words[i].getValue() + ","
        }
        return msg.slice(0, msg.length - 1)
    }
}

module.exports = Card