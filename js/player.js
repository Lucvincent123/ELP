require("dotenv").config()

const net = require("net")

const readline = require("readline-sync")

const Player = require("./classes/Player")


const newPlayer = new Player()

console.log(newPlayer.getName())

newPlayer.setName(readline.question("Please enter your name: "))
const tcpClient = new net.Socket()
tcpClient.connect(process.env.PORT, process.env.HOST, () => {
    console.log("You entered the game")
    tcpClient.write(`name ${newPlayer.getName()}`)
})

tcpClient.on("data", (data) => {
    let cmd = data.toString().split()
    switch (cmd[0]) {
        case "start":
            console.log("start")
            tcpClient.write("ready", "utf-8")
    }
})
