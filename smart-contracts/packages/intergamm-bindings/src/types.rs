use schemars::JsonSchema;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema, Debug)]
#[serde(rename_all = "snake_case")]
pub struct IntergammAck {
    sequence: u64,
    error: Option<String>,
    response: Option<crate::msg::AckResponse>,
}