use relay3;

CREATE PROCEDURE get_conversations()
BEGIN

SELECT 
	JSON_ARRAYAGG(JSON_OBJECT(
		'id', c.id, 
        'zid', c.zid,
        'communityzid', c.community_zid,
        'text', c.text,
        'link', c.link,
        'img', c.img,
        'video', c.video,
        'public', c.public,
        'publicprice', c.public_price,
        'created', c.created,
        'updated', c.updated,
        'deleted', c.deleted,
        'comments', co.comments)
	) as 'Result'
FROM 
	relay3.conversations c
LEFT JOIN (
	SELECT 
		conversation_zid, 
		JSON_ARRAYAGG(JSON_OBJECT(
			'id', co.id, 
            'zid', co.zid,
            'userdid', co.user_did,
            'text', co.text,
            'link', co.link,
            'created', co.created,
            'updated', co.updated,
            'deleted', co.deleted)
		) AS comments
	FROM relay3.comments co
	GROUP BY co.conversation_zid
) co ON co.conversation_zid = c.zid;

END