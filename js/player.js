const net = require("net")

const readline = require("readline-sync")

const Player = require("./classes/Player")

const constants = require("./constant")



const newPlayer = new Player()

newPlayer.setName(readline.question("Please enter your name: ").trim())
const tcpClient = new net.Socket()
tcpClient.connect(constants.PORT, constants.HOST, () => {
    console.log("You entered the game")
})

tcpClient.on("data", (data) => {
    let cmd = data.toString().split(" ")
    switch (cmd[0]) {
        case "wait_pick":
            console.log("######################")
            console.log(`Points : ${cmd[2]}`)
            console.log(`Cards left : ${cmd[3]}`)
            console.log(`${cmd[1]} is picking a card`)
            break
            
        case "pick":
            console.log("######################")
            console.log(`Points : ${cmd[1]}`)
            console.log(`Cards left : ${cmd[2]}`)
            console.log("You are picking a card")
            tcpClient.write("pick")
            break
                
        case "wait_choose":
            console.log("######################")
            console.log(`Points : ${cmd[2]}`)
            console.log(`Cards left : ${cmd[3]}`)
            console.log("Card picked")
            showCard(cmd[4])
            console.log(`${cmd[1]} is choosing a number`)
            break
            
        case "choose":
            console.log("######################")
            console.log(`Points : ${cmd[1]}`)
            console.log(`Cards left : ${cmd[2]}`)
            console.log("Card picked")
            let number = readline.question(`Please pick a number from 1 to ${constants.NUMBER_OF_WORDS_PER_CARD}: `)
            tcpClient.write(`choose ${number}`)
            break
                
        case "wait_hint":
            console.log("######################")
            console.log(`Points : ${cmd[1]}`)
            console.log(`Cards left : ${cmd[2]}`)
            console.log("Waiting for hints")
            break
                    
        case "hint":
            console.log("######################")
            console.log(`Points : ${cmd[2]}`)
            console.log(`Cards left : ${cmd[3]}`)
            console.log(`Word to guess : ${cmd[4]}`)
            let hint = readline.question(`Give a hint for ${cmd[1]} : `).trim()
            tcpClient.write(`hint ${hint}`)
            break
                        
        case "wait_guess":
            console.log("######################")
            console.log(`Points : ${cmd[2]}`)
            console.log(`Cards left : ${cmd[3]}`)
            console.log(`Word to guess : ${cmd[4]}`)
            console.log(`Waiting for ${cmd[1]} to guess`)
            break
                            
                            
        case "guess":
            console.log("######################")
            console.log(`Points : ${cmd[1]}`)
            console.log(`Cards left : ${cmd[2]}`)
            showHints(cmd[3])
            let guess = readline.question(`Make a guess (PASS TURN if you don't want to) : `).trim()
            tcpClient.write(`guess ${guess}`)
            break
            
        case "result":
            console.log("######################")
            if (cmd[1] == "pass") {
                
            } else {
                console.log(`You got a ${cmd[1]} answer`)
                if (cmd[1] == "right") {
                    console.log("You win 1 point")
                } else {
                    console.log("You lost one card")
                }
            }
            console.log("Move to other player")
            tcpClient.write("pick")
            break
            
        case "other_result":
            console.log("######################")
            if (cmd[2] == "pass") {
                console.log(`${cmd[1]} pass turn`)
            } else {
                console.log(`${cmd[1]} got a ${cmd[2]} answer`)
                if (cmd[2] == "right") {
                    console.log("You win 1 point")
                } else {
                    console.log("You lost one card")
                }
            }
            console.log("Move to other player")
            break
            
        case "close":
            console.log("end")
            tcpClient.write("close")
            tcpClient.end()
        }
    })
tcpClient.write(`name ${newPlayer.getName()}`)

const showCard = (string) => {
    console.log("+=============+")
    string.split(",").forEach((word, index) => {
        console.log(`| ${index + 1}. ${word} |`)
    })
    console.log("+=============+")    
    
}

const showHints = (string) => {
    console.log("===============")
    console.log("Hints: ")
    string.split(",").forEach((hint) => {
        console.log(`- ${hint}`)
    })
    console.log("===============")    
    
}
