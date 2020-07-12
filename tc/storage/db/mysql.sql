-- the table to store BranchSession data
CREATE TABLE IF NOT EXISTS `branch_table`
(
    `branch_id` VARCHAR(128) NOT NULL,
    `xid`       VARCHAR(128) NOT NULL,
    `status`    TINYINT,
    PRIMARY KEY (`branch_id`, `xid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

