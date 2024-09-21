use push;

CREATE TABLE user_info
(
    userid BIGINT AUTO_INCREMENT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),

    ccode VARCHAR(8) NOT NULL DEFAULT '86', -- country code
    phone VARCHAR(32) NOT NULL, -- phone number

    PRIMARY KEY(userid)
) ENGINE=InnoDB AUTO_INCREMENT=1000000 PARTITION BY HASH(userid DIV 1000000) PARTITIONS 10;
CREATE INDEX user_info_idx0 ON user_info(phone, ccode); -- need program make a UNIQUE contract

CREATE TABLE msg_batch
(
    batch_id BIGINT AUTO_INCREMENT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    cron_time DATETIME NOT NULL DEFAULT '0001-01-01', -- cron time
    PRIMARY KEY(batch_id)
) ENGINE=InnoDB;

CREATE TABLE msg_status
(
    batch_id VARCHAR(64) NOT NULL, -- campaing batch
    userid BIGINT NOT NULL, -- phone number
    created_at DATETIME NOT NULL DEFAULT NOW(),
    sent_time DATETIME NOT NULL DEFAULT '0001-01-01', -- sent time
    readed_time DATETIME NOT NULL DEFAULT '0001-01-01', -- readed time
    PRIMARY KEY(batch_id, userid)
) ENGINE=InnoDB;
