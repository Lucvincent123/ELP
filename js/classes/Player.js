class Player {
    constructor() {
        this.easel = ""
    }
    
    setName(name) {
        this.name = name
    }

    getName() {
        return this.name || "XXX"
    }

    // setSocket(socket) {
    //     this.socket = socket
    // }

    // getSocket() {
    //     return this.socket
    // }
}

module.exports = Player