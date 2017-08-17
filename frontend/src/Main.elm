module Main exposing (main)

import Navigation exposing (Location)
import Html exposing (Html, text)
import UrlParser exposing (..)
import Page.Signup as Signup
import Page.Login as Login
import Page.Patient.Dashboard as PatientDash

-- Wires


main  =
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
    | LoginR
    | PatientDashboardR


-- TODO remove or use me.
-- type SubMsg msg = Goto Route | Sub msg

matchers : Parser (Route -> a) a
matchers =
    oneOf
        [ map SignupR (s "signup")
        , map LoginR (s "login")
        , map PatientDashboardR (s "dashboard")
        ]


parseLocation : Location -> Route
parseLocation location =
    case (parseHash matchers location) of
        Just route ->
            route

        Nothing ->
            NotFoundR



-- Model


type alias Model =
    { location : Route
    , signupModel : Signup.Model
    , loginModel : Login.Model
    , patientDashModel : PatientDash.Model
    }


type Msg
    = OnLocationChange Location
    | SignupMsg Signup.Msg
    | LoginMsg Login.Msg
    | PatientDashMsg PatientDash.Msg




init : Navigation.Location -> ( Model, Cmd Msg )
init location =
     { location = parseLocation location
     , signupModel = Signup.init
     , loginModel = Login.init
     , patientDashModel = PatientDash.init
     }
       ! [Cmd.none] 
    




-- Update


{-| 
   The challenge in updating nested features is more apparent when a subview
   requires a command as it is rendered in a way not entirely distinct from 
   React's "ComponentDidMount".

   In the case of Page.Patient.Dashboard the `case parseLocation location of` 
   cathces the `OnLocationChange` for PatientDashboardR and runs the needed
   command.

   Future refactoring to to provide an abstraction of this is advised.
   Perhaps as Evancz suggested `type Submsg msg = Goto Route | Sub msg`. 


-} 


update :  Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of

       OnLocationChange location ->
       
            case parseLocation location of 

                PatientDashboardR ->
                    let
                        ( subMdl, subCmd ) =
                             PatientDash.update 
                                 PatientDash.GetNextApptTime model.patientDashModel

                    in
                        { model | location = parseLocation location
                                , patientDashModel = subMdl }
                                      ! [ Cmd.map PatientDashMsg subCmd ]          


                _ ->
                     ( { model | location = parseLocation location }, Cmd.none )

       SignupMsg m ->
            let
                ( subMdl, subCmd ) =
                    Signup.update m model.signupModel
            in
                { model | signupModel = subMdl }
                    ! [ Cmd.map SignupMsg subCmd ]

       LoginMsg m ->
            let 
                ( subMdl, subCmd ) = 
                    Login.update m model.loginModel
            in
                { model | loginModel = subMdl }
                    ! [ Cmd.map LoginMsg subCmd ]

       PatientDashMsg m ->
            let
                ( subMdl, subCmd ) =
                    PatientDash.update m model.patientDashModel
            in
                { model | patientDashModel = subMdl }
                    ! [ Cmd.map PatientDashMsg subCmd ] 




-- View



view : Model -> Html Msg
view model =
    page model 


page : Model -> Html Msg
page model =
    case model.location of
        LoginR ->
            Login.view model.loginModel
                |> Html.map LoginMsg
        SignupR ->
            Signup.view model.signupModel
                |> Html.map SignupMsg 

        PatientDashboardR ->
            PatientDash.view model.patientDashModel
                |> Html.map PatientDashMsg

        NotFoundR ->
            text "404: Not found"

