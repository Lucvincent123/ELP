module Parsing exposing (..)

import Parser exposing (..)
import Model exposing (..)



right : Parser Command             ---- Parse commande Right
right =
    succeed Right
        |. keyword "Right"
        |. spaces
        |= int

forward : Parser Command
forward =
    succeed Forward
        |. keyword "Forward"
        |. spaces
        |= int

left : Parser Command
left =
    succeed Left
        |. keyword "Left"
        |. spaces
        |= int


repeat : Parser Command
repeat =
    succeed Repeat
        |. keyword "Repeat"
        |. spaces
        |= int
        |. spaces
        |= lazy (\_ -> commandsParser)

parseCommand : Parser Command                ----- Parse une commande
parseCommand =
    oneOf
        [ right
        , forward
        , left
        , repeat
        ]

commandsParser : Parser (List Command)      ----- Parse une liste de commands separes par "," et cadre par "[]", c'est ce qu'on va utiliser
commandsParser =
    sequence
        { start = "["
        , separator = ","
        , end = "]"
        , spaces = spaces
        , item = parseCommand
        , trailing = Forbidden               ----- Forbiden : interdire "," apres le dernier element
        }


------Example: run commandsParser "[Forward 100]"  -> Ok [Forward 100]. We gonna use this in UPDATE function