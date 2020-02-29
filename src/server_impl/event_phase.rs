use crate::server_impl::{
    ulz_proto::{Empty, GdGetInfoReq, GdPhaseConfirmReq, GdPhaseConfirmResp},
    GameDuelServiceBackend,
};
use tonic::{Request, Response, Status};

pub  fn event_phase_result(
    this: &GameDuelServiceBackend,
    request: Request<GdGetInfoReq>,
) -> Result<Response<GdPhaseConfirmResp>, tonic::Status> {
    Ok(Response::new(GdPhaseConfirmResp::default()))
}

pub fn event_phase_confirm(
    this: &GameDuelServiceBackend,    
    request: Request<GdPhaseConfirmReq>,
) -> Result<Response<Empty>, tonic::Status> {
    Ok(Response::new(Empty::default()))
}
