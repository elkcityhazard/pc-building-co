CREATE OR REPLACE TABLE pcbco_img_collections(
	 pcbco_img_collection_id BIGINT PRIMARY KEY AUTO_INCREMENT
	,pcbco_img_id BIGINT
	,pcbco_img_collection_created_at DATETIME DEFAULT NOW()
	,pcbco_img_collection_updated_at DATETIME DEFAULT NOW() ON UPDATE NOW()
	,pcbco_img_collection_version INT DEFAULT 1
	,CONSTRAINT fk_pcbco_img_collection_pcbo_img_id FOREIGN KEY (pcbco_img_id) REFERENCES pcbco_imgs(pcbco_img_id)
);