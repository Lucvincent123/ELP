class Player {
    constructor() {
        this.easel = "" // String
        // this.name  String
    }
    
    setName(name) {
        this.name = name // String
    }

    getName() {
        return this.name || "XXX"  // String
    }

    setEasel(hint) { // String
        this.easel = hint
    }

    getEasel() {
        return this.easel
    }

    // setSocket(socket) {
    //     this.socket = socket
    // }

    // getSocket() {
    //     return this.socket
    // }
}

module.exports = Player