module Patient.Main exposing (..)

import Html exposing (..)
import Html.Attributes exposing ( class 
                                , name
                                , placeholder
                                , type_
                                )
import Html.Events exposing (onInput)
import Http
import Json.Encode as JsonE



-- Wires



main = 
    Html.program 
        { init = init
        , update = update
        , view = view
        , subscriptions = (\_ -> Sub.none)
        }



-- Model



type alias Model = { nextAppointmentTime : String 
                   , email    : String
                   , password : String 
                   }

type Msg  
    = GetNextAppointmentTime (Result Http.Error String)
    | SetCredentialEmail String
    | SetCredentialPassword String

init : (Model, Cmd Msg) 
init = ({ nextAppointmentTime = "No Appointment Scheduled" 
        , email = ""
        , password = ""}, getNextAppointmentTime)



-- Update




update : Msg -> Model -> ( Model, Cmd msg )
update msg model = 
    case msg of
        GetNextAppointmentTime (Ok time) ->
            ({ model | nextAppointmentTime = time }, Cmd.none)
        GetNextAppointmentTime (Err _) ->
            (model, Cmd.none )
        SetCredentialEmail email ->
            ({ model | email = email }, Cmd.none)
        SetCredentialPassword pw ->
            ({ model | password = pw }, Cmd.none)

            


-- Server API


baseUrl : String
baseUrl = "http://localhost:8080/"

signupUri : String
signupUri = "signup"

{-
signup email pw = 
   let
       json = JsonE.encode 0 [ ("Email", string email)
                             , ("Password", string pw)
                             ]
   in
       Http.send Signup <| Http.post (baseUrl ++ signupUri) json
-}


getNextAppointmentTime =
   Http.send GetNextAppointmentTime <| Http.getString (baseUrl ++ "string")



-- View



view : Model -> Html Msg
view model = 
    div [] [ div [ class "card"] [ signupForm ]
           , infoRow model.nextAppointmentTime 
           ]



infoRow time = div [] [ nextAppointmentInfo time, pairedProviderInfo ] 

nextAppointmentInfo : String -> Html msg
nextAppointmentInfo time = 
    div [] [ text time ]

pairedProviderInfo : Html msg
pairedProviderInfo = 
    div [] [ text "someone" ]


-- Signup View

signupForm = 
    div [] 
        [ input [ type_ "text", placeholder "Email" , onInput SetCredentialEmail ] []
        , input [ type_ "password", placeholder "Password", onInput SetCredentialPassword ] [] 
        , input [ type_ "submit", placeholder "signup" ] []
        ]

