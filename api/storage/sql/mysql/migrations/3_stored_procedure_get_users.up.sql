CREATE PROCEDURE get_users()
BEGIN

SELECT 
	JSON_ARRAYAGG(JSON_OBJECT(
		'Id', u.id, 
        'Did', u.did, 
        'Username', u.username,
        'Email', u.email,
        'Name', u.name,
        'Bio', u.bio,
        'Img', u.img,
        'PriceToMessage', u.price_to_message,
        'Created', u.created,
        'Updated', u.updated)
	) as 'Result'
FROM 
	users u;

END