module Page.Signup exposing (..)

import Html exposing (..)
import Html.Attributes exposing ( class 
                                , name
                                , placeholder
                                , type_
                                )
import Html.Events exposing (onInput, onClick)
import Http
import Json.Encode as JsonE
import Json.Decode as JsonD



type alias Model = { email    : String
                   , password : String 
                   }

type Msg  
    = HandleSignup
    | Signup (Result Http.Error String)
    | SetEmail String
    | SetPassword String


init : Model
init = { email = ""
       , password = ""
       }



 -- Update




update : Msg -> Model -> ( Model, Cmd Msg )
update msg model = 
    case msg of
        HandleSignup ->
            (model, signup model.email model.password)
        Signup (Ok _) ->
            (model, Cmd.none) 
        Signup (Err _) ->
            (model, Cmd.none)
        SetEmail email ->
            ({ model | email = email }, Cmd.none)
        SetPassword pw ->
            ({ model | password = pw }, Cmd.none)


            


-- Server API


baseUrl : String
baseUrl = "http://localhost:8080/"

signupUri : String
signupUri = "signup"


signup email pw = 
   let
       url = baseUrl ++ signupUri
       body = Http.jsonBody 
                    <| JsonE.object [ ("Email", JsonE.string email)
                                    , ("Password", JsonE.string pw)
                                    ]
       decoder = JsonD.string
   in
       Http.send Signup <| Http.post url body decoder




-- Signup View

signupForm : Html Msg
signupForm = 
    div [] 
        [ input [ type_ "text"
                , placeholder "Email" 
                , onInput SetEmail ] 
                []
        , input [ type_ "password"
                , placeholder "Password"
                , onInput SetPassword ] 
                [] 
        , input [ type_ "submit"
                , placeholder "signup"
                , onClick HandleSignup ] 
                []
        ]
                  
