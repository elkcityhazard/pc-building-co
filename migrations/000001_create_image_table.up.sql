CREATE OR REPLACE TABLE pcbco_imgs(
   pcbco_img_id BIGINT PRIMARY KEY AUTO_INCREMENT
  ,pcbco_img_dir_path VARCHAR(250) NOT NULL DEFAULT ""
  ,pcbco_img_filename VARCHAR(250) NOT NULL DEFAULT "" UNIQUE
  ,pcbco_img_created_at DATETIME NOT NULL DEFAULT NOW()
  ,pcbco_img_updated_at DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW()
  ,pcbco_img_version INT NOT NULL DEFAULT 1
);

CREATE OR REPLACE INDEX idx_pcbo_img_filename ON pcbco_imgs(pcbco_img_filename);
