module Main exposing (..)
import Parser exposing (..)
import Svg exposing (..)
import Svg.Attributes exposing (..)
import Browser
import Html exposing (Html, div, input, text, button)
import Html.Attributes exposing (..)
import Html.Events exposing (onInput, onClick)
import Parsing exposing (..)
import Display exposing (..)
import Model exposing(..)




-- MAIN

main =
  Browser.sandbox { init = init, update = update, view = view }


-- MODEL



init : Model
init =
  { content = [] , input = ""}



-- UPDATE


type Msg
    = Change String   --- Add evenement when we type
    | Submit  -- Add evenement submit



update : Msg -> Model -> Model
update msg model =
    case msg of
        Change newContent ->
            { model | input = newContent }  

        Submit ->
            -- After submit, execute theses commands
            case run commandsParser model.input of
                Ok commands -> { model | content = commands }  
                Err _ -> { model | input = "Hehehhehe, you typed a wrong command, please retry" }    -- If wrong format, send this message


-- VIEW


view : Model -> Html Msg
view model =
    div []
    [ div [ Html.Attributes.style "text-align" "center", Html.Attributes.style "margin-top" "20px" ] -- Marging and Spacing, Position in the midle
        [ Html.text "Type in your code below:" ]
    , div [ Html.Attributes.style "margin-top" "20px", Html.Attributes.style "text-align" "center" ]  
        [ input
            [ placeholder "example:[Repeat 360 [ Right 1, Forward 5]]"
            , value (model.input)    --- Associate input to model.input
            , onInput Change
            , Html.Attributes.style "width" "50%" 
            , Html.Attributes.style "height" "20px"
            , Html.Attributes.style "display" "block"  
            , Html.Attributes.style "margin-left" "auto"
            , Html.Attributes.style "margin-right" "auto"
            ]
            []
        ]
    , div [ Html.Attributes.style "margin-top" "20px", Html.Attributes.style "text-align" "center" ]
        [ Html.text ("Received: " ++ model.input) ]
    , div [ Html.Attributes.style "margin-top" "20px", Html.Attributes.style "text-align" "center" ]  
        [ button [ onClick Submit ] [ Html.text "Draw" ] ]
    , div [ Html.Attributes.style "margin-top" "20px", Html.Attributes.style "text-align" "center" ]  
        [ Display.display model ]
    ]
