module Main exposing (main)

import Navigation exposing (Location)
import UrlParser exposing (..)
import Page.Signup exposing (Model)


-- Wires



main = 
    Navigation.program 
        { init = init
        , update = update
        , view = view
        , subscriptions = (\_ -> Sub.none)
        }




-- Routing

type Route
    = NotFoundR
    | SignupR

matchers : Parser (Route -> a) a
matchers = 
    oneOf 
        [ map SignupR (s "signup") 
        ]

parseLocation : Location -> Route
parseLocation location = 
    case (parseHash matchers location) of 
        Just route -> 
            route

        Nothing
            NotFoundR

-- Model

type alias Model = { location : Route 
                   }

type Msg
    = OnLocationChange Location 

init : (Model, Cmd Msg) 



-- Update




update : Msg -> Model -> ( Model, Cmd Msg )
update msg model = 
    case msg of
        OnLocationChange l ->
            ( { model | location = parseLocation }, Cmd.none )


-- View



view : Model -> Html Msg
view model = 
    viewPage page

viewPage : page -> Html Msg
    case page of 
        


