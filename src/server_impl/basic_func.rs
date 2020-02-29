use crate::server_impl::{
    ulz_proto::{Empty, GameDataSet, GdBroadcastResp, GdCreateReq, GdGetInfoReq},
    GameDuelServiceBackend,
};
use tokio::sync::mpsc;
use tonic::{Request, Response, Status};
pub fn create_game(
    this: &GameDuelServiceBackend,
    request: Request<GdCreateReq>,
) -> Result<Response<GameDataSet>, Status> {
    Ok(Response::new(GameDataSet::default()))
}

pub fn get_game_data(
    this: &GameDuelServiceBackend,
    request: Request<GdGetInfoReq>,
) -> Result<Response<GameDataSet>, Status> {
    Ok(Response::new(GameDataSet::default()))
}

pub fn quit_game(
    this: &GameDuelServiceBackend,
    request: Request<GdCreateReq>,
) -> Result<Response<Empty>, Status> {
    Ok(Response::new(Empty::default()))
}

pub fn server_broadcast(
    this: &GameDuelServiceBackend,
    request: Request<GdGetInfoReq>,
) -> Result<Response<mpsc::Receiver<Result<GdBroadcastResp, Status>>>, Status> {
    // Ok()
    Err(unimplemented!())
}
