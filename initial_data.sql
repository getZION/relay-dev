-- Users

INSERT INTO `relay3`.`users` 
(`did`, `email`, `id`, `name`, `picture`, `price_to_message`, `username`) 
VALUES 
('did1', 'test1@org.com', '1', 'name1', 'pic1', '1', 'test_username1');

INSERT INTO `relay3`.`users` 
(`did`, `email`, `id`, `name`, `picture`, `price_to_message`, `username`) 
VALUES 
('did2', 'test2@org.com', '2', 'name2', '', '2', 'test_username2');



-- Communities

INSERT INTO `relay3`.`communities` 
(`created`, `deleted`, `description`, `escrow_amount`, `id`, `img`, `last_active`, `name`, `owner_did`, `owner_username`, `price_per_message`, `price_to_join`, `public`, `updated`, `zid`) 
VALUES 
('1', '1', 'desc', '10', '1', 'img1', '1', 'name1', 'did1', 'test_username1', '1', '10', '1', '1', 'zid1');



-- Community Users

INSERT INTO `relay3`.`community_users` (`CommunityZid`, `UserDid`) VALUES ('zid1', 'did1');
INSERT INTO `relay3`.`community_users` (`CommunityZid`, `UserDid`) VALUES ('zid1', 'did2');




-- Conversation

INSERT INTO `relay3`.`conversations` 
(`community_zid`, `created`, `deleted`, `id`, `img`, `link`, `public`, `public_price`, `text`, `updated`, `video`, `zid`) 
VALUES 
('zid1', '1', '0', '1', 'img1', 'link', '1', '1', 'text1', '1', 'video1', 'conversation_zid1');



-- Comments

INSERT INTO `relay3`.`comments` 
(`id`, `zid`, `user_did`, `conversation_zid`, `text`, `link`, `created`, `updated`, `deleted`) 
VALUES 
('1', 'zid1', 'did1', 'conversation_zid1', 'test_test', '', '1', '1' , '0');



