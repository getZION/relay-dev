INSERT INTO `relay3`.`users` 
(`did`, `email`, `id`, `name`, `picture`, `price_to_message`, `username`) 
VALUES 
('did1', 'test1@org.com', '1', 'name1', '', '1', 'test_username1');

INSERT INTO `relay3`.`users` 
(`did`, `email`, `id`, `name`, `picture`, `price_to_message`, `username`) 
VALUES 
('did2', 'test2@org.com', '2', 'name2', '', '2', 'test_username2');

INSERT INTO `relay3`.`communities` 
(`created`, `deleted`, `description`, `escrow_amount`, `id`, `img`, `last_active`, `name`, `owner_did`, `owner_username`, `price_per_message`, `price_to_join`, `public`, `updated`, `zid`) 
VALUES 
('1', '1', 'desc', '10', '1', '', '1', 'comm1', 'did1', 'username', '1', '10', '1', '1', 'zid1');


INSERT INTO `relay3`.`community_users`
(`CommunityZid`, `UserDid`)
VALUES
(`zid1`, `did1`)
(`zid1`, `did2`)


INSERT INTO `relay3`.`conversations` 
(`community_zid`, `created`, `deleted`, `id`, `img`, `link`, `public`, `public_price`, `text`, `updated`, `video`, `zid`) 
VALUES 
('zid1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', 'conversation_zid1');



