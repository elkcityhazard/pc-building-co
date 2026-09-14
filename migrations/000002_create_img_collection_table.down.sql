ALTER TABLE IF EXISTS pcbco_img_collections
	DROP FOREIGN KEY IF EXISTS fk_pcbco_img_collection_pcbo_img_id;

DROP TABLE IF EXISTS pcbco_img_collections;