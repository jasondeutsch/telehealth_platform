module Patient.Main exposing (..)

import Html exposing (..)
import Html.Attributes exposing ( class 
                                , enctype
                                , type_
                                , method
                                , name
                                , placeholder
                                , formaction )
import Http



-- Wires



main = 
    Html.program 
        { init = init
        , update = update
        , view = view
        , subscriptions = (\_ -> Sub.none)
        }



-- Model



type alias Model = { nextAppointmentTime : String }

type Msg = 
    GetNextAppointmentTime (Result Http.Error String)

init : (Model, Cmd Msg) 
init = ({ nextAppointmentTime = "No Appointment Scheduled" }, getNextAppointmentTime)



-- Update


update : Msg -> Model -> ( Model, Cmd msg )
update msg model = 
    case msg of
        GetNextAppointmentTime (Ok time) ->
            ({ model | nextAppointmentTime = time }, Cmd.none)
        GetNextAppointmentTime (Err _) ->
            (model, Cmd.none )
            


-- Server API


baseUrl : String
baseUrl = "http://localhost:8080/"

getNextAppointmentTime =
   Http.send GetNextAppointmentTime <| Http.getString (baseUrl ++ "string")



-- View


view : Model -> Html msg
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
    form [ formaction (baseUrl ++ "signup"), method "POST", enctype "application/x-www-form-urlencoded" ] 
        [ input [ type_ "text", placeholder "Email", name "Email"] []
        , input [ type_ "password", placeholder "Password", name "Password" ] [] 
        , input [ type_ "submit", placeholder "signup" ] []
        ]

