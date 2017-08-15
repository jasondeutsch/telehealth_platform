module Page.Patient.Dashboard exposing (..)


import Html exposing (..)
import Http



-- Model & Messages


type alias Model =
    { nextApptTime : String }


type Msg
    = GetNextApptTime
    | NextApptTime (Result Http.Error String)


init : Model
init =
    { nextApptTime = ""  
    }


-- Update



update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
       GetNextApptTime ->
           (model, getNextApptTime)

       NextApptTime (Ok time) ->
            ( { model | nextApptTime = time }, Cmd.none )

       NextApptTime (Err _) ->
           ( model, Cmd.none )



-- Server


baseUrl : String
baseUrl =
    "http://localhost:8080/"



getNextApptTime =
    Http.send NextApptTime <| Http.getString (baseUrl ++ "string")



-- View


view : Model -> Html Msg
view model =
    div []
        [ infoRow model.nextApptTime
        ]


infoRow time =
    div [] 
        [ nextAppointmentInfo time
        , pairedProviderInfo ]


nextAppointmentInfo : String -> Html msg
nextAppointmentInfo time =
    div [] [ text time ]


pairedProviderInfo : Html msg
pairedProviderInfo =
    div [] [ text "someone" ]
