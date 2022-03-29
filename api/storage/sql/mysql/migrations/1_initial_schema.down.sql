CREATE TABLE `communities` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `zid` varchar(255) NOT NULL,
  `name` varchar(150) NOT NULL,
  `owner_did` varchar(255) NOT NULL,
  `owner_username` varchar(255) NOT NULL,
  `description` varchar(250) NOT NULL,
  `escrow_amount` bigint NOT NULL,
  `img` varchar(255) DEFAULT NULL,
  `last_active` bigint DEFAULT NULL,
  `price_per_message` bigint NOT NULL,
  `price_to_join` bigint NOT NULL,
  `public` tinyint(1) DEFAULT '1',
  `created` bigint NOT NULL,
  `updated` bigint DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `zid` (`zid`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `conversations` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `zid` varchar(255) DEFAULT NULL,
  `community_zid` varchar(255) NOT NULL,
  `text` varchar(255) DEFAULT NULL,
  `link` varchar(255) DEFAULT NULL,
  `img` varchar(255) DEFAULT NULL,
  `video` varchar(255) DEFAULT NULL,
  `public` tinyint(1) DEFAULT '0',
  `public_price` bigint DEFAULT '0',
  `created` bigint NOT NULL,
  `updated` bigint DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `zid` (`zid`),
  KEY `community_zid` (`community_zid`),
  CONSTRAINT `conversations_community_zid_communities_Zid_foreign` FOREIGN KEY (`community_zid`) REFERENCES `communities` (`zid`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `did` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `bio` varchar(255) DEFAULT NULL,
  `img` varchar(255) DEFAULT NULL,
  `price_to_message` bigint DEFAULT NULL,
  `created` bigint DEFAULT NULL,
  `updated` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `did` (`did`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `tags` (
  `tag` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`tag`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `comments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `zid` varchar(255) NOT NULL,
  `conversation_zid` varchar(255) NOT NULL,
  `user_did` varchar(255) NOT NULL,
  `text` varchar(255) DEFAULT NULL,
  `link` varchar(255) DEFAULT NULL,
  `created` bigint NOT NULL,
  `updated` bigint DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `conversation_zid` (`conversation_zid`),
  KEY `user_did` (`user_did`),
  CONSTRAINT `comments_conversation_zid_conversations_Zid_foreign` FOREIGN KEY (`conversation_zid`) REFERENCES `conversations` (`zid`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `comments_user_did_users_Did_foreign` FOREIGN KEY (`user_did`) REFERENCES `users` (`did`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `community_tags` (
    `community_zid`   NVARCHAR(255)   NOT NULL UNIQUE,
    `tag`             NVARCHAR(64)    NOT NULL UNIQUE,
	FOREIGN KEY (community_zid) REFERENCES  communities(zid),
    FOREIGN KEY (tag)           REFERENCES  tags(tag)
);

CREATE PROCEDURE `get_communities` ()
BEGIN

SELECT 
	JSON_ARRAYAGG(JSON_OBJECT('id', c.id, 'name', c.name, 'zid', c.zid, 'tags', t.tags, 'users', u.users)) as 'Result'
FROM 
	relay3.communities c
LEFT JOIN (
	SELECT 
		community_id, 
		JSON_ARRAYAGG(ct.tag) AS tags
	FROM relay3.community_tags ct
	GROUP BY ct.community_id
) t ON t.community_id = c.id
LEFT JOIN (
	SELECT 
		CommunityZid, 
		JSON_ARRAYAGG(cu.UserDid) AS users
	FROM relay3.community_users cu
	GROUP BY cu.CommunityZid
) u ON u.CommunityZid = c.zid;

END;
