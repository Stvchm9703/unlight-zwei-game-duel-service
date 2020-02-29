#![allow(unused_imports, dead_code, unreachable_code)]

use std::collections::HashMap;
use std::hash::{Hash, Hasher};
use std::pin::Pin;
use std::sync::Arc;
use std::time::Instant;
use tokio::sync::mpsc;
use tonic::transport::Server;
use tonic::{Request, Response, Status};

pub mod ulz_proto {
    tonic::include_proto!("ulz_proto");
}

// use ulz_proto::game_duel_service_server::{GameDuelService, GameDuelServiceServer};

mod server_impl;
use server_impl::GameDuelServiceBackend;

fn main() {
    server_impl::run_server();
}