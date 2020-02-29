use crate::server_impl::{
    ulz_proto::{Empty, GdGetInfoReq, GdMoveConfirmReq, GdMoveConfirmResp},
    GameDuelServiceBackend,
};
use tonic::{Request, Response, Status};

pub fn move_phase_confirm(
    this: &GameDuelServiceBackend,
    request: Request<GdMoveConfirmReq>,
) -> Result<Response<Empty>, Status> {
    Ok(Response::new(Empty::default()))
}

pub fn move_phase_result(
    this: &GameDuelServiceBackend,
    request: Request<GdGetInfoReq>,
) -> Result<Response<GdMoveConfirmResp>, Status> {
    // Ok()
    Err(unimplemented!())
}
