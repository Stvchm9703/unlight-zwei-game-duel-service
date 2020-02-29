use crate::server_impl::{
    ulz_proto::{Empty, GdGetInfoReq, GdadConfirmReq, GdadDiceResult, GdadResultResp},
    GameDuelServiceBackend,
};
use tonic::{Request, Response, Status};

pub fn ad_phase_confirm(
    this: &GameDuelServiceBackend,
    request: Request<GdadConfirmReq>,
) -> Result<Response<Empty>, Status> {
    Ok(Response::new(Empty::default()))
}
pub fn ad_phase_result(
    this: &GameDuelServiceBackend,
    request: Request<GdGetInfoReq>,
) -> Result<Response<GdadResultResp>, Status> {
    Err(unimplemented!())
}

pub fn ad_phase_dice_result(
    this: &GameDuelServiceBackend,
    request: Request<GdGetInfoReq>,
) -> Result<Response<GdadDiceResult>, Status> {
    Err(unimplemented!())
}
