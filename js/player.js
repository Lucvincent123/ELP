require("dotenv").config()

const net = require("net")

const readline = require("readline").createInterface({
    input: process.stdin,
    output: process.stdout
})

const Player = require("./classes/Player")


const newPlayer = new Player()

console.log(newPlayer.getName())

readline.question("Please enter your name: ", (name) => {
    newPlayer.setName(name.trim())
    readline.close()
})

const tcpClient = new net.Socket()
// tcpClient.connect(process.env.PORT, process.env.HOST, () => {
//     console.log("You entered the game")
// })

// tcpClient.on("data", (data) => {
//     if (data == "name") {
//         tcpClient.write(`name ${newPlayer.get}`)
//     }
// })
