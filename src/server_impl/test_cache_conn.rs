use crate::server_impl;
use crate::server_impl::{
    config::CfDatabase
};
// use crate::server_impl

let test_cache_db = config::CfDatabase{
    connector : Some( "redis"),
    worker_node: Some(1),
    host: Some( "192.168.0.110"),
    port : Some(6379),
    max_pool_size : Some(1),
    username: Some(""),
    password
};
#[cfg(test)]
mod test_cache_conn{
    #[test]
    pub async fn test_conn() {
        let backend = GameDuelServiceBackend {
            rooms: Vec::new(),
            rds_conn: Box::new(redis::Client::open("redis://192.168.0.110:6379").unwrap()),
        };
        backend.init_conn(test_cache_db);

    }
}



fn main() {
    test_cache_conn::test_conn().await();
}