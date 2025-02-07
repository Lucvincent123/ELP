class Word {
    constructor(word) {
        this.value = word // String
        this.sameWords = [word]
    }

    addSameWord(word) { // string
        this.sameWords.push(word)
    }

    isSameWord(string) { //String
        return this.sameWords.includes(string)
    }

    getValue() {
        return this.value // String
    }
}

module.exports = Word