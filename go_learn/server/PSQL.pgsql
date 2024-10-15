--INSERT INTO images_client_not_look VALUES (1,'2.jpg')

--DELETE FROM images_client_not_look

--INSERT INTO images_client_not_look VALUES ($1,$2)

--SELECT * FROM images_client_not_look WHERE name_file='4.jpg'

--CREATE INDEX idx_name_file ON images_client_not_look (name_file)

--SELECT EXISTS (SELECT * FROM images_client_not_look WHERE LOWER(TRIM(name_file)) = LOWER(TRIM('3.jpg')))

--SELECT NOT EXISTS(SELECT * FROM images_client_not_look WHERE name_file=$1);

SELECT name_file FROM images_client_not_look