module Display exposing (..)

import Svg exposing (..)
import Svg.Attributes exposing (..)
import Parsing exposing(..)
import Model exposing(..)
import Html exposing(..)

display : Model -> Html msg
display model =
    let
        start = { x = 200, y = 200, angle = 0 }  -- Begin in the middle
        path = tracePath start model.content   ---- please read function tracePath
    in
    div []
        [ svg [ Svg.Attributes.width "400", Svg.Attributes.height "400", viewBox "0 0 400 400" ]
            [ polyline [ points (formatPoints path), fill "none", stroke model.color, strokeWidth (String.fromInt model.largeur) ] [] 
    , rect [ x "0", y "0", width "400", height "400", fill "none", stroke "green", strokeWidth "5" ] []  -- Add rectangle in border
            ]  --- read function formatPoints 
        ]


    
type alias Turtle =        ---- Define a turtle with position x, y and angle
    { x : Float
    , y : Float
    , angle : Float
    }

tracePath : Turtle -> List Command -> List (Float, Float)    --- Create a list float float from the commands list
tracePath start commands =
    let
        -- Execute the command one by one 
        step turtle command =
            case command of
                Forward n ->
                    let
                        rad = degrees turtle.angle                         ------ If go foward, same angle, and x, y using sin,cos to calculate
                        newX = turtle.x + toFloat n * cos rad
                        newY = turtle.y - toFloat n * sin rad
                    in
                    ( { turtle | x = newX, y = newY }, [(newX, newY)] )   ------ update new values to turtle

                Left n ->
                    ( { turtle | angle = turtle.angle + toFloat n }, [] )   ------- If turn left, same x,y but angle change

                Right n ->
                    ( { turtle | angle = turtle.angle - toFloat n }, [] )

                -- I used recursivite for this Repeat function
                Repeat times subcommands ->
                        let                                                            
                            repeatedPoints = tracePath turtle (List.concat (List.repeat times subcommands))   --- use again tracePath to list in Repeat 
                            points = tracePath start subcommands                                                ---which is dupplicated by TIMES times.
                        in                                                                                      ---- So i used repeat and concat
                        (turtle, points ++ repeatedPoints)

        -- Write all the position that turtle arrive
        traceRecursively turtle remainingCommands =
            case remainingCommands of
                [] -> []
                cmd :: rest ->
                    let
                        (newT, points) = step turtle cmd
                    in
                    points ++ traceRecursively newT rest

    in
    -- Always start with the initial position
    (start.x, start.y) :: traceRecursively start commands




formatPoints : List (Float, Float) -> String       ------ svg.poliline accept only this string format "float,float float,float .....", so we gonna change the
formatPoints points =                                               --- list obtained in TracePath to this string format
    String.join " " (List.map (\(x, y) -> String.fromFloat x ++ "," ++ String.fromFloat y) points)