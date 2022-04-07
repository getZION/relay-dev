
ALTER TABLE users
  ADD amount float NULL DEFAULT '0'
    AFTER name;

DROP TABLE IF EXISTS payments;
CREATE TABLE payments (
  id                    bigint        NOT NULL AUTO_INCREMENT,
  zid                   varchar(256)  NOT NULL,
  recipient_did         varchar(256)  NOT NULL,
  sender_did            varchar(256)  NOT NULL,
  amount                float         NOT NULL,
  type                  smallint      DEFAULT '1',
  status                smallint      DEFAULT '1',
  relevant_type         varchar(32)   NULL,
  relevant_zid          varchar(256)  NULL,
  recipient_node_pubkey varchar(256)  NOT NULL,
  recipient_relay_url   varchar(256)  NOT NULL,
  created               bigint        NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY id (id),
  UNIQUE KEY zid (zid),
  CONSTRAINT payment_recipient_did_users_did_foreign FOREIGN KEY (recipient_did) REFERENCES users (did) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT payment_sender_did_users_did_foreign FOREIGN KEY (sender_did) REFERENCES users (did) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

/*

- types 
    - 1 => boost
    - 2 => p2p
    - 3 => community_join
    - 4 => stake

- statuses
    - 1 => pending
    - 2 => completed
    - 3 => failed

*/
