module Main exposing (main)

import Navigation exposing (Location)
import UrlParser exposing (..)
import Page.Signup
import Html exposing (Html, text)


-- Wires



main = 
    Navigation.program OnLocationChange 
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

        Nothing ->
            NotFoundR

-- Model


type alias Model = { location : Route 
                   , signupModel : Page.Signup.Model
                   }


type Msg
    = OnLocationChange Location 
    | SignupMsg Page.Signup.Msg 


init : Navigation.Location -> (Model, Cmd Msg) 
init location = 
    ( { location = parseLocation location
      , signupModel = Page.Signup.init
      }, Cmd.none )



-- Update




update : Msg -> Model -> ( Model, Cmd Msg )
update msg model = 
    case msg of
        OnLocationChange l ->
            ( { model | location = parseLocation l }, Cmd.none )
  
        SignupMsg m ->
            let 
                (subMdl, subCmd) = Page.Signup.update m model.signupModel
            in
                { model | signupModel = subMdl }
                   ! [ Cmd.map SignupMsg subCmd ] 

-- View



view : Model -> Html Msg
view model = page model

page : Model -> Html Msg
page model =
    case model.location of
        SignupR -> 
            Html.map SignupMsg <| Page.Signup.view model.signupModel

        NotFoundR ->
            text "404: Not found"
        


