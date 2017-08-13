module Page.Patient.Dashboard exposing (..)


type alias Model =
    { nextAppointmentTime : String }


type Msg
    = GetNextAppointmentTime (Result Http.Error String)


init =
    ( { nextAppointmentTime = "None Scheduled" }, getNextAppointmentTime )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        GetNextAppointmentTime (Ok time) ->
            ( { model | nextAppointmentTime = time }, Cmd.none )

        GetNextAppointmentTime (Err _) ->
            ( model, Cmd.none )



-- Server


baseUrl : String
baseUrl =
    "http://localhost:8080/"


getNextAppointmentTime =
    Http.send GetNextAppointmentTime <| Http.getString (baseUrl ++ "string")



-- View


view : Model -> Html Msg
view model =
    div []
        [ infoRow model.nextAppointmentTime
        ]


infoRow time =
    div [] [ nextAppointmentInfo time, pairedProviderInfo ]


nextAppointmentInfo : String -> Html msg
nextAppointmentInfo time =
    div [] [ text time ]


pairedProviderInfo : Html msg
pairedProviderInfo =
    div [] [ text "someone" ]
