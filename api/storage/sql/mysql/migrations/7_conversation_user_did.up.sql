
ALTER TABLE conversations
  ADD user_did varchar(256) NOT NULL AFTER community_zid,
  ADD CONSTRAINT conversation_user_did_users_Did_foreign FOREIGN KEY (user_did) REFERENCES users (did) ON DELETE RESTRICT ON UPDATE RESTRICT;