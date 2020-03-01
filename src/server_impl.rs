#![allow(unused_imports, dead_code, unreachable_code)]

use std::boxed::Box;
use std::collections::HashMap;
use std::hash::{Hash, Hasher};
use std::pin::Pin;
use std::sync::Arc;
use std::time::Instant;

use futures::{Stream, StreamExt};
use redis::{Client, Commands};
use tokio::sync::mpsc;
use tonic;
use tonic::transport::Server;
use tonic::{Request, Response, Status};

// #[derive(Debug)]
pub mod ulz_proto {
    tonic::include_proto!("ulz_proto");
}

// service.proto : generated by code
use ulz_proto::game_duel_service_server::{GameDuelService, GameDuelServiceServer};

// Data.proto
use ulz_proto::{
    CharCardEquSet, CharCardSet, EffectResult, EventCard, EventCardPos, EventCardType, GameDataSet,
    PlayerSide, RangeType, SignEq, SkillCardCond, SkillSet, StatusEffect, StatusSet,
};

// EventHookPhase.proto
use ulz_proto::{Empty, EventHookPhase, EventHookType};

// message.proto
use ulz_proto::{
    CastCmd, EcShortHand, GdBroadcastResp, GdChangeConfirmReq, GdCreateReq, GdGetInfoReq,
    GdInstanceDt, GdMoveConfirmReq, GdMoveConfirmResp, GdPhaseConfirmReq, GdPhaseConfirmResp,
    GdadConfirmReq, GdadDiceResult, GdadResultResp, MovePhaseOpt,
};

pub mod ad_phase;
pub mod basic_func;
pub mod cache_conn;
pub mod config;
pub mod draw_phase;
pub mod event_phase;
pub mod instance_ec;
pub mod move_phase;
use crate::server_impl::cache_conn::RedisBox;
pub async fn test_with_config(path: &str) -> GameDuelServiceBackend {
    let test_cache_db: config::CfDatabase = config::CfDatabase {
        connector: Some("redis".to_string()),
        worker_node: Some(1),
        host: Some("192.168.0.110".to_string()),
        port: Some(6379),
        username: Some("".to_string()),
        password: Some("".to_string()),
        database: Some("redis".to_string()),
        filepath: Some("".to_string()),
    };

    let mut test_gameset: ulz_proto::GameDataSet = ulz_proto::GameDataSet::default();
    test_gameset.room_key = "h1234".to_string();

    let ymp = config::parse(path).await;
    let mut backend = GameDuelServiceBackend {
        rooms: Vec::new(),
        rds_conn: Box::new(redis::Client::open("redis://127.0.0.1:6379").unwrap()),
    };

    match backend.init_cache_conn(test_cache_db).await {
        Ok(r) => {
            println!("ok");
        }
        Err(e) => {
            println!("{}", e);
        }
    };

    backend.set_room_data(test_gameset).await;
    let ky = backend.get_room_data("h1234").await;
    println!("{:?}", ky.unwrap());
    return backend;
}
pub async fn init_with_config(path: &str) -> GameDuelServiceBackend {
    let test_cache_db: config::CfDatabase = config::CfDatabase {
        connector: Some("redis".to_string()),
        worker_node: Some(1),
        host: Some("192.168.0.110".to_string()),
        port: Some(6379),
        username: Some("".to_string()),
        password: Some("".to_string()),
        database: Some("redis".to_string()),
        filepath: Some("".to_string()),
    };
    let ymp = config::parse(path).await;
    let mut backend = GameDuelServiceBackend {
        rooms: Vec::new(),
        rds_conn: Box::new(redis::Client::open("redis://127.0.0.1:6379").unwrap()),
    };
    match backend.init_cache_conn(test_cache_db).await {
        Ok(r) => {
            println!("ok");
        }
        Err(e) => {
            println!("{}", e);
        }
    };

    return backend;
}

#[tokio::main]
pub async fn run_server() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:10000".parse().unwrap();

    println!("RouteGuideServer listening on: {}", addr);

    let backend = init_with_config("").await;
    let svc = GameDuelServiceServer::new(backend);
    println!("running on: {}", addr);
    Server::builder().add_service(svc).serve(addr).await?;
    Ok(())
}

#[derive(Debug)]
pub struct GameDuelServiceBackend {
    rooms: Vec<Box<ulz_proto::GameDataSet>>,
    rds_conn: Box<redis::Client>,
}

// impl Copy for GameDuelServiceBackend{}


#[tonic::async_trait]
impl GameDuelService for GameDuelServiceBackend {
    // -------------------------------------------------------------------------------
    // basic function :
    //      forward to basic_func.rs
    // -------------------------------------------------------------------------------
    async fn create_game(
        &self,
        _request: Request<GdCreateReq>,
    ) -> Result<Response<GameDataSet>, tonic::Status> {
        return basic_func::create_game(&self, _request);
    }

    async fn get_game_data(
        &self,
        _request: Request<GdGetInfoReq>,
    ) -> Result<Response<GameDataSet>, tonic::Status> {
        return basic_func::get_game_data(&self, _request);
    }

    async fn quit_game(
        &self,
        _request: Request<GdCreateReq>,
    ) -> Result<Response<Empty>, tonic::Status> {
        return basic_func::quit_game(&self, _request);
    }

    type ServerBroadcastStream = mpsc::Receiver<Result<GdBroadcastResp, tonic::Status>>;
    async fn server_broadcast(
        &self,
        _request: Request<GdGetInfoReq>,
    ) -> Result<Response<Self::ServerBroadcastStream>, tonic::Status> {
        return basic_func::server_broadcast(&self, _request);
    }
    // -------------------------------------------------------------------------------
    //  instance get/set event-card:
    //      forward to instance_ec.rs
    // -------------------------------------------------------------------------------
    async fn inst_set_event_card(
        &self,
        _request: tonic::Request<GdInstanceDt>,
    ) -> Result<tonic::Response<Empty>, tonic::Status> {
        return instance_ec::inst_set_event_card(&self, _request);
    }
    async fn inst_get_event_card(
        &self,
        _request: tonic::Request<GdGetInfoReq>,
    ) -> Result<tonic::Response<GdInstanceDt>, tonic::Status> {
        return instance_ec::inst_get_event_card(&self, _request);
    }
    // -------------------------------------------------------------------------------
    //   Draw Phase confirm :
    //      forward to draw_phase.rs
    // -------------------------------------------------------------------------------
    async fn draw_phase_confirm(
        &self,
        _request: Request<GdGetInfoReq>,
    ) -> Result<Response<Empty>, tonic::Status> {
        return draw_phase::draw_phase_confirm(&self, _request);
    }
    // -------------------------------------------------------------------------------
    //  Move Phase :
    //      forward to move_phase.rs
    // -------------------------------------------------------------------------------
    async fn move_phase_confirm(
        &self,
        _request: Request<GdMoveConfirmReq>,
    ) -> Result<Response<Empty>, Status> {
        return move_phase::move_phase_confirm(&self, _request);
    }
    async fn move_phase_result(
        &self,
        _request: Request<GdGetInfoReq>,
    ) -> Result<Response<GdMoveConfirmResp>, Status> {
        return move_phase::move_phase_result(&self, _request);
    }
    // -------------------------------------------------------------------------------
    //  Event Phase :
    //      forward to event_phase.rs
    // -------------------------------------------------------------------------------
    async fn event_phase_result(
        &self,
        _request: Request<GdGetInfoReq>,
    ) -> Result<Response<GdPhaseConfirmResp>, tonic::Status> {
        return event_phase::event_phase_result(&self, _request);
    }
    async fn event_phase_confirm(
        &self,
        _request: Request<GdPhaseConfirmReq>,
    ) -> Result<Response<Empty>, tonic::Status> {
        return event_phase::event_phase_confirm(&self, _request);
    }
    // -------------------------------------------------------------------------------
    // Atk/Def Phase :
    //      forward to ad_phase.rs
    // -------------------------------------------------------------------------------
    async fn ad_phase_confirm(
        &self,
        _request: Request<GdadConfirmReq>,
    ) -> Result<Response<Empty>, Status> {
        return ad_phase::ad_phase_confirm(&self, _request);
    }
    async fn ad_phase_result(
        &self,
        _request: Request<GdGetInfoReq>,
    ) -> Result<Response<GdadResultResp>, Status> {
        return ad_phase::ad_phase_result(&self, _request);
    }

    async fn ad_phase_dice_result(
        &self,
        _request: Request<GdGetInfoReq>,
    ) -> Result<Response<GdadDiceResult>, Status> {
        return ad_phase::ad_phase_dice_result(&self, _request);
    }
    // -------------------------------------------------------------------------------
    //  Change character phase :
    //      !FIXME  not-impelement
    // -------------------------------------------------------------------------------
    async fn change_phase_confirm(
        &self,
        _request: tonic::Request<GdChangeConfirmReq>,
    ) -> Result<tonic::Response<Empty>, tonic::Status> {
        Ok(Response::new(Empty::default()))
    }
    async fn change_phase_result(
        &self,
        _request: tonic::Request<GdGetInfoReq>,
    ) -> Result<tonic::Response<Empty>, tonic::Status> {
        Ok(Response::new(Empty::default()))
    }
    // -------------------------------------------------------------------------------
    // -------------------------------------------------------------------------------
}
