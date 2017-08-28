module Auth exposing (..)

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


type Role
    = Patient
    | Provider
    | Admin


type alias Email =
    String


type alias Password =
    String


type Msg
    = HandleSignup
    | Signup (Result Http.Error String)
    | HandleLogin
    | LoginResult (Result Http.Error String)
    | SetEmail Email
    | SetPassword Password


type alias Model =
    { email : Email
    , password : Password
    }


init : Model
init =
    { email = ""
    , password = ""
    }



-- Update


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        HandleSignup ->
            ( model, signup model.email model.password )

        Signup (Ok _) ->
            ( model, Cmd.none )

        Signup (Err _) ->
            ( model, Cmd.none )

        HandleLogin ->
            ( model, login model.email model.password )

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


signupUri : String
signupUri =
    "signup"


loginUri : String
loginUri =
    "login"


signup : Email -> Password -> Cmd Msg
signup e p =
    let
        url =
            baseUrl ++ signupUri

        body =
            Http.jsonBody <|
                JsonE.object
                    [ ( "email", JsonE.string e )
                    , ( "password", JsonE.string p )
                    ]

        decoder =
            JsonD.string
    in
        Http.send Signup <| Http.post url body decoder


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



-- Signup View


view : Model -> Html Msg
view model =
    signupForm


signupForm : Html Msg
signupForm =
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
            , value "Signup"
            , onClick HandleSignup
            ]
            []
        ]


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
