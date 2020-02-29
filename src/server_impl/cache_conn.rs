use crate::server_impl::{config::CfDatabase, ulz_proto::GameDataSet, GameDuelServiceBackend};
use bytes::{Buf, BufMut, BytesMut};
use std::{error::Error, result};

use prost::Message;
use redis::{Commands, ErrorKind, RedisError, RedisResult};
use tonic::codec::{Codec, DecodeBuf, Decoder, ProstCodec};

#[allow(unused_imports, dead_code, unreachable_code)]
impl GameDuelServiceBackend {
    pub async fn init_conn(&mut self, db_cf: CfDatabase) -> Result<bool, RedisError> {
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

    pub async fn get_data(&self, room_key: &str) -> Result<GameDataSet, RedisError> {
        let mut inst_conn = self.rds_conn.get_connection()?;
        let mut d: Vec<u8> = inst_conn.get(room_key)?;
        let mut buff = BytesMut::with_capacity(d.len());
        buff.put_slice(&d);

        let mut decoder = ProstCodec::<GameDataSet, GameDataSet>::default().decoder();
        let mut decb: DecodeBuf = DecodeBuf {
            buf: &mut buff,
            len: d.len(),
        };
        match decoder.decode(&mut decb) {
            Ok(opt) => Ok(opt.unwrap_or_default()),
            Err(e) => Ok(GameDataSet::default()),
        }
    }
    fn set_room() {}
}
