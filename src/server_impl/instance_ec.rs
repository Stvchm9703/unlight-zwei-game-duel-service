use crate::server_impl::{ulz_proto, GameDuelServiceBackend};
use tonic::{Request, Response, Status};

pub fn inst_set_event_card(
    this: &GameDuelServiceBackend,
    request: Request<ulz_proto::GdInstanceDt>,
) -> Result<Response<ulz_proto::Empty>, Status> {
    // return
    Ok(Response::new(ulz_proto::Empty::default()))
}

pub fn inst_get_event_card(
    this: &GameDuelServiceBackend,
    request: Request<ulz_proto::GdGetInfoReq>,
) -> Result<Response<ulz_proto::GdInstanceDt>, Status> {
    // return
    Err(unimplemented!())
}
