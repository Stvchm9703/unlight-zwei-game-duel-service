use crate::server_impl::{
    ulz_proto::{Empty, GameDataSet, GdBroadcastResp, GdCreateReq, GdGetInfoReq},
    GameDuelServiceBackend,
};
use tokio::sync::mpsc;
use tonic::{Request, Response, Status};
pub fn create_game(
    _this: &GameDuelServiceBackend,
    _request: Request<GdCreateReq>,
) -> Result<Response<GameDataSet>, Status> {
    Ok(Response::new(GameDataSet::default()))
}

pub fn get_game_data(
    _this: &GameDuelServiceBackend,
    _request: Request<GdGetInfoReq>,
) -> Result<Response<GameDataSet>, Status> {
    Ok(Response::new(GameDataSet::default()))
}

pub fn quit_game(
    _this: &GameDuelServiceBackend,
    _request: Request<GdCreateReq>,
) -> Result<Response<Empty>, Status> {
    Ok(Response::new(Empty::default()))
}

pub fn server_broadcast(
    _this: &GameDuelServiceBackend,
    _request: Request<GdGetInfoReq>,
) -> Result<Response<mpsc::Receiver<Result<GdBroadcastResp, Status>>>, Status> {
    // Ok()
    Err(unimplemented!())
}
