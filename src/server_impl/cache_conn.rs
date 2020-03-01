use crate::server_impl::{config::CfDatabase, ulz_proto::GameDataSet, GameDuelServiceBackend};
use bytes::{Buf, BufMut, BytesMut};
use std::{error::Error, result};
use tonic::codec::{Codec, DecodeBuf, Decoder, EncodeBuf, Encoder, ProstCodec};

use prost::Message;
use redis::{Commands, ErrorKind, RedisError, RedisResult};

#[allow(unused_imports, dead_code, unreachable_code)]
#[tonic::async_trait]
pub trait RedisBox {
    async fn init_conn(&mut self, db_cf: CfDatabase) -> Result<bool, RedisError>;
    async fn get_data(mut self, room_key: &str) -> Result<GameDataSet, RedisError>;
    async fn set_room(mut self, _room: GameDataSet) -> Result<(), RedisError>;
}

#[tonic::async_trait]
impl RedisBox for GameDuelServiceBackend {
    async fn init_conn(&mut self, db_cf: CfDatabase) -> Result<bool, RedisError> {
        let conn = db_cf.connector.unwrap_or_else(|| "redis".to_string());
        let host = db_cf.host.unwrap_or_else(|| "127.0.0.1".to_string());
        let port = db_cf.port.unwrap_or_else(|| 6379);
        let addr = format!("{}://{}:{}", conn, &host, port);
        match redis::Client::open(addr) {
            Ok(result) => {
                // if self.rds_conn.is_empty() {
                self.rds_conn = Box::new(result);
                // } else {
                //     println!("self.rds_conn is not empty");
                // }
                Ok(true)
            }
            Err(e) => Err(e),
        }
    }

    async fn get_data(mut self, room_key: &str) -> Result<GameDataSet, RedisError> {
        let mut inst_conn = self.rds_conn.get_connection()?;
        let d: Vec<u8> = inst_conn.get(room_key)?;
        let mut buff = BytesMut::with_capacity(d.len());
        buff.put_slice(&d);

        let mut decoder = ProstCodec::<GameDataSet, GameDataSet>::default().decoder();
        let mut decb: DecodeBuf = DecodeBuf::new(&mut buff, d.len());
        match decoder.decode(&mut decb) {
            Ok(opt) => Ok(opt.unwrap_or_default()),
            Err(e) => Ok(GameDataSet::default()),
        }
    }
    async fn set_room(mut self, _room: GameDataSet) -> Result<(), RedisError> {
        let inst_conn = self.rds_conn.get_connection()?;
        let mut encoder = ProstCodec::<GameDataSet, GameDataSet>::default().encoder();
        let mut buff = BytesMut::new();
        let mut encb: EncodeBuf = EncodeBuf::new(&mut buff);
        let key = (_room.room_key).clone();
        encoder.encode(_room, &mut encb);
        self.rds_conn.set(key, buff.to_vec())?;
        // println!(buff.to_bytes());
        Ok(())
    }
}
