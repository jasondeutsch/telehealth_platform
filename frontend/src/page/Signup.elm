module Page.Signup exposing (..)

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


-- As noted in Login.elm, both modules are so similiar 
-- that either abstraction or lifting into top level may be appropriate.


type alias Model =
    { email : String
    , password : String
    }


type Msg
    = HandleSignup
    | Signup (Result Http.Error String)
    | SetEmail String
    | SetPassword String


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


signup email pw =
    let
        url =
            baseUrl ++ signupUri

        body =
            Http.jsonBody <|
                JsonE.object
                    [ ( "email", JsonE.string email )
                    , ( "password", JsonE.string pw )
                    ]

        decoder =
            JsonD.string
    in
        Http.send Signup <| Http.post url body decoder



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
