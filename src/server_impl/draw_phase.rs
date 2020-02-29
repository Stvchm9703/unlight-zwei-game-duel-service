use crate::server_impl::{
    ulz_proto::{
        Empty,GdGetInfoReq
    }, 
    GameDuelServiceBackend};
use tonic::{Request, Response, Status};


pub fn draw_phase_confirm(
    this: &GameDuelServiceBackend,
    request: Request<GdGetInfoReq>,
) -> Result<Response<Empty>, Status> {
    Ok(Response::new(Empty::default()))
}
