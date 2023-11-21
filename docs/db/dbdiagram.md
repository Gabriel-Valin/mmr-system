## DBML - DB Diagram IO

##### https://dbdiagram.io/

```dbml
Table players {
  id string [primary key]
  username string
  password string
  created_at timestamp
  updated_at timestamp
  deleted_at timestamp
}


Table rankeds_status {
  id string [primary key]
  player_id string [ref: - players.id]
  points float
  mmr float
  md5_completed boolean
  created_at timestamp
  updated_at timestamp
  deleted_at timestamp
}

Table matches {
  id string [primary key]
  a_team_players json
  b_team_players json
  a_team_mmr float
  b_team_mmr float
  team_won string
  mvp string
  created_at timestamp
  updated_at timestamp
  deleted_at timestamp
}

Table perfomances {
  id string [primary key]
  player_id string [ref: - players.id]
  kd float
  kda float
  won boolean
  points string
  match_rank_player integer
}
```
