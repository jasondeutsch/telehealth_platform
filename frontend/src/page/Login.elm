module Page.Login exposing (..)

import Html exposing (..)
import Html.Attributes
    exposing
        ( class
        , name
        , placeholder
        , type_
        , value 
        )
import Html.Events exposing (onInput, onClick)
import Http
import Json.Encode as JsonE
import Json.Decode as JsonD


-- Signup and Login are so similiar, perhaps a user construct should be
-- brought to the top level.

type alias Model = 
    { email : Email
    , password : Password
    }

type alias Email = String
type alias Password = String

type Msg
    = HandleLogin
    | LoginResult (Result Http.Error String)
    | SetEmail Email
    | SetPassword Password

init : Model
init = 
    { email = ""
    , password = ""
    }


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        HandleLogin ->
           ( model , login model.email model.password ) 

        LoginResult (Ok _) ->
            ( model, Cmd.none )

        LoginResult (Err _) ->
            ( model, Cmd.none )

        SetEmail email ->
            ( { model | email = email }, Cmd.none )

        SetPassword pw ->
            ( { model | password = pw }, Cmd.none )



-- Server API

baseUrl : String
baseUrl =
    "http://localhost:8080/"


loginUri : String
loginUri =
    "login"


login : Email -> Password -> Cmd Msg 
login e p =
    let
        url =
            baseUrl ++ loginUri

        body =
            Http.jsonBody <|
                JsonE.object
                    [ ( "email", JsonE.string e )
                    , ( "password", JsonE.string p )
                    ]

        decoder =
            JsonD.string
    in
        Http.send LoginResult <| Http.post url body decoder



-- View

view : Model -> Html Msg
view model = loginForm
    

loginForm : Html Msg
loginForm =
    div []
        [ input
            [ type_ "text"
            , placeholder "Email"
            , onInput SetEmail
            ]
            []
        , input
            [ type_ "password"
            , placeholder "Password"
            , onInput SetPassword
            ]
            []
        , input
            [ type_ "submit"
            , value "Login"
            , onClick HandleLogin
            ]
            []
        ] 
