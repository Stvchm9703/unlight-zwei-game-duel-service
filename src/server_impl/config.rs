#[allow(unused_imports, dead_code, unreachable_code)]
use serde::{Deserialize, Serialize};
use std::fmt::Debug;
use std::fs::File;
use std::io::Read;
use std::ptr::null;
use toml;
// use std::file::File;

#[derive(Serialize, Deserialize, Debug)]
pub struct ConfTmp {
    APIServer: Option<CfAPIServer>,
    Database: Option<CfDatabase>,
    CacheDb: Option<CfDatabase>,
    AuthServer: Option<CfAPIServer>,
}
impl ConfTmp {
    // add code here
    pub fn new() -> ConfTmp {
        ConfTmp {
            APIServer: Option::default(),
            Database: Option::default(),
            CacheDb: Option::default(),
            AuthServer: Option::default(),
        }
    }
}
#[derive(Serialize, Deserialize, Debug)]
pub struct CfAPIServer {
    pub ConnType: Option<String>,
    pub ip: Option<String>,
    pub port: Option<u16>,
    pub max_pool_size: Option<u16>,
    pub api_refer_type: Option<String>,
    pub api_table_path: Option<String>,
    pub api_outpath: Option<String>,
}
impl CfAPIServer {
    // add code here
    pub fn new() -> CfAPIServer {
        CfAPIServer {
            ConnType: Option::default(),
            ip: Option::default(),
            port: Option::default(),
            max_pool_size: Option::default(),
            api_refer_type: Option::default(),
            api_outpath: Option::default(),
            api_table_path: Option::default(),
        }
    }
}

#[derive(Serialize, Deserialize, Debug)]
pub struct CfDatabase {
    pub connector: Option<String>,
    pub worker_node: Option<u16>,
    pub host: Option<String>,
    pub port: Option<u16>,
    pub username: Option<String>,
    pub password: Option<String>,
    pub database: Option<String>,
    pub filepath: Option<String>,
}

impl CfDatabase {
    // add code here
    pub fn new() -> CfDatabase {
        CfDatabase {
            connector: Option::default(),
            worker_node: Option::default(),
            host: Option::default(),
            port: Option::default(),
            username: Option::default(),
            password: Option::default(),
            database: Option::default(),
            filepath: Option::default(),
        }
    }
}

pub fn parse(path: &str) -> ConfTmp {
    let mut config_toml = String::new();

    let mut file = match File::open(path) {
        Ok(file) => file,
        Err(_) => {
            println!("Could not find config file in {}, using default!", path);
            return ConfTmp::new();
        }
    };

    file.read_to_string(&mut config_toml)
        .unwrap_or_else(|err| panic!("Error while reading config: [{}]", err));

    let config: ConfTmp = toml::from_str(&config_toml).unwrap();
    println!("{:#?}", config);
    return config;
}
