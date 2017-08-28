module Page.Admin.Dashbaord exposing (..)


import Html exposing (..)
import Http
import Json.Decode as Decode



type Msg 
    = FetchPatients 
    | FetchPatients (Result Http.Error String)
    -- | Pair 
    -- | Unpair
    -- | ApproveProvider
    -- | RemoveProvider

type alias Patient = 
    { id : Int
    }

type alias Model = [List Patient]



-- Update



update : Msg -> Model -> (Model, Cmd Msg)
update msg model = 
    case msg of 
        FetchPatients ->
            (model, fetchPatients)

        FetchPatients (Ok patients) -> 
            ( { model | patientes = patients }, Cmd.none)

        FetchPAtients (Err _) ->
            (model, Cmd.none)




-- Server API


baseUrl : String
baseUrl = "http://localhost:8080"

fetchPatients : Cmd Msg
fetchPatients = 
    let
        url = baseUrl + "/admin/patients/all"
        decoder = Decode.list (Decode.int)

    in
        Http.send FetchPatients <| Http.get url decoder


-- View

view : Model -> Html Msg
view = text "admin/dashboard"
