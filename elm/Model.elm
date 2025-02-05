module Model exposing(..)   


type Command = Left Int | Right Int | Forward Int | Repeat Int (List Command)     ------Create type Command and Model to exposing


type alias Model = 
  { content : List Command , input : String}