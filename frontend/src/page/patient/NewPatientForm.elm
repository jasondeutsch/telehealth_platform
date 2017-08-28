module Page.Patient.NewPatientForm exposing (..)

import Html exposing (..)
import Http


type Msg 
    = Save
    | Save (Result Http.Error String)





-- Model


type alias Model = ""

-- Update

update : Msg -> Model -> (Cmd Msg)
update msg model = 
    case msg of
        NoOp -> (Model, Cmd.none)


-- Server API



--- View

view : Model -> Html Msg
view model = text "hello"



