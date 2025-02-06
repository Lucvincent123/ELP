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
  { content = [] , input = "", color = "cyan", largeur = 1 }



-- UPDATE


type Msg
    = Change String   --- Add evenement when we type
    | Submit  -- Add evenement submit
    | Color   -- Change color
    | Increment  --- Increment or Decrement StrokeWidth "largeur"
    | Decrement


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
        Color -> 
            case model.color of
                "cyan" -> { model | color = "red"}
                "red"  -> { model | color = "blue"}
                "blue"  -> { model | color = "green"}
                "green"  -> { model | color = "yellow"}
                "yellow"  -> { model | color = "cyan"}
                _  -> { model | color = "cyan"}
        Increment->
            { model | largeur = model.largeur + 1 }  

        Decrement->
            if model.largeur > 1 then
                { model | largeur = model.largeur - 1 }  
            else 
                model  

    
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
        [ button [ onClick Decrement ] [ Html.text "-" ]
        , Html.text ("Width: " ++ String.fromInt model.largeur)
        , button [ onClick Increment ] [ Html.text "+" ]
        ]
    , div [ Html.Attributes.style "margin-top" "20px", Html.Attributes.style "text-align" "center" ]  
        [ button [ onClick Submit ] [ Html.text "Draw" ]
        , button [ onClick Color ] [ Html.text "Change Color" ]
        , Html.text (model.color)]
    , div [ Html.Attributes.style "margin-top" "20px", Html.Attributes.style "text-align" "center" ]  
        [ Display.display model ]
    ]
