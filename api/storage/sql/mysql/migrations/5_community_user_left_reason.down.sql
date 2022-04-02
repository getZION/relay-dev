
ALTER TABLE community_users
  ADD left_reason varchar(64) NOT NULL
    AFTER left_date;