---------------------------------------------------------------
-- App container
---------------------------------------------------------------
CREATE TABLE IF NOT EXISTS apple (
  id                 BIGINT PRIMARY KEY,
  name               TEXT,
  url                TEXT,
  icon               TEXT,
  kind               TEXT,
  version            TEXT,
  bundle_id          TEXT,
  author_id          BIGINT,
  author_name        TEXT,
  author_url         TEXT,
  vendor_name        TEXT,
  vendor_url         TEXT,
  copyright          TEXT,
  genre_id           BIGINT,
  genre_id_list      BIGINT [],
  genre              TEXT,
  genre_list         TEXT [],
  icon60             TEXT,
  icon100            TEXT,
  price              BIGINT,
  currency           TEXT,
  system             TEXT,
  features           TEXT [],
  devices            TEXT [],
  languages          TEXT [],
  platforms          TEXT [],
  rating             TEXT,
  reasons            TEXT [],
  size               BIGINT,
  cnt_rating         BIGINT,
  avg_rating         NUMERIC,
  cnt_rating_current BIGINT,
  avg_rating_current NUMERIC,
  vpp_device         BOOLEAN,
  game_center        BOOLEAN,
  screenshots        TEXT [],
  in_app_purchase    TEXT [],
  sibling_apps       BIGINT [],
  related_apps       BIGINT [],
  support_sites      JSONB,
  reviews            JSONB,
  extra              JSONB,
  description        TEXT,
  release_notes      TEXT,
  release_time       TIMESTAMPTZ,
  publish_time       TIMESTAMPTZ,
  crawled_time       TIMESTAMPTZ
);


COMMENT ON TABLE apple IS 'apple应用数据表';
COMMENT ON COLUMN apple.id IS '应用ID，又名iTunesID,trackID,整型';
COMMENT ON COLUMN apple.name IS '应用名称';
COMMENT ON COLUMN apple.url IS '应用页面URL';
COMMENT ON COLUMN apple.icon IS '应用图标URL(512px)';
COMMENT ON COLUMN apple.kind IS '应用类型(software)';
COMMENT ON COLUMN apple.version IS '应用版本';
COMMENT ON COLUMN apple.bundle_id IS '应用包名BundleID';
COMMENT ON COLUMN apple.author_id IS '作者ID';
COMMENT ON COLUMN apple.author_name IS '作者名';
COMMENT ON COLUMN apple.author_url IS '作者页面URL';
COMMENT ON COLUMN apple.vendor_name IS '厂商名称';
COMMENT ON COLUMN apple.vendor_url IS '厂商URL(外部)';
COMMENT ON COLUMN apple.copyright IS '版权信息';
COMMENT ON COLUMN apple.genre_id IS '首要类型ID';
COMMENT ON COLUMN apple.genre_id_list IS '类型ID数组';
COMMENT ON COLUMN apple.genre IS '首要类型名称';
COMMENT ON COLUMN apple.genre_list IS '类型名称数组';
COMMENT ON COLUMN apple.icon60 IS '图标URL(60px)';
COMMENT ON COLUMN apple.icon100 IS '图标URL(100px)';
COMMENT ON COLUMN apple.price IS '价格，0为免费';
COMMENT ON COLUMN apple.currency IS '货币';
COMMENT ON COLUMN apple.system IS '系统版本要求';
COMMENT ON COLUMN apple.features IS '特性，通常为iosUniversal';
COMMENT ON COLUMN apple.devices IS '支持的设备列表';
COMMENT ON COLUMN apple.languages IS '支持的语言，ISO2字符代号，已排序';
COMMENT ON COLUMN apple.platforms IS '支持的平台:iPad,iPhone,iPod,iWatch,AppleTV';
COMMENT ON COLUMN apple.rating IS '应用分级';
COMMENT ON COLUMN apple.reasons IS '分级原因';
COMMENT ON COLUMN apple.size IS '文件大小';
COMMENT ON COLUMN apple.cnt_rating IS '评分人数(历史所有版本)';
COMMENT ON COLUMN apple.avg_rating IS '平均评分(历史所有版本)';
COMMENT ON COLUMN apple.cnt_rating_current IS '评分人数(当前版本)';
COMMENT ON COLUMN apple.avg_rating_current IS '平均评分(当前版本)';
COMMENT ON COLUMN apple.vpp_device IS '是否支持VppDevice';
COMMENT ON COLUMN apple.game_center IS '是否启用游戏中心';
COMMENT ON COLUMN apple.screenshots IS '截图[URL数组]';
COMMENT ON COLUMN apple.in_app_purchase IS '应用内购`(rank,price,title)`';
COMMENT ON COLUMN apple.sibling_apps IS '同一开发者的应用';
COMMENT ON COLUMN apple.related_apps IS '苹果推荐的应用';
COMMENT ON COLUMN apple.support_sites IS '应用支持站点，KV JSON Object';
COMMENT ON COLUMN apple.reviews IS '用户评论`(user,rating,title,content)`';
COMMENT ON COLUMN apple.extra IS '额外信息，JSONB kv格式';
COMMENT ON COLUMN apple.description IS '应用描述';
COMMENT ON COLUMN apple.release_notes IS '最近更新内容';
COMMENT ON COLUMN apple.release_time IS '更新时间';
COMMENT ON COLUMN apple.publish_time IS '初始发布时间';
COMMENT ON COLUMN apple.crawled_time IS '抓取时间';
---------------------------------------------------------------


---------------------------------------------------------------
-- Task Queue
---------------------------------------------------------------
-- DROP TABLE apple_queue;
CREATE TABLE IF NOT EXISTS apple_queue (
  id TEXT PRIMARY KEY
);
COMMENT ON TABLE apple_queue IS 'Apple Task Queue';
-----------------------------------------
-- Function: add apple id to queue
CREATE OR REPLACE FUNCTION apple_aid(_id BIGINT)
  RETURNS VOID AS
$$BEGIN INSERT INTO apple_queue (id) VALUES ('!' || _id :: TEXT);
END;$$
LANGUAGE plpgsql VOLATILE;
COMMENT ON FUNCTION apple_aid(BIGINT) IS '向苹果队列中添加iTunesID任务';
-- SELECT apple_aid(1031569344)
-----------------------------------------
-- Function: add apple bundle id to queue
CREATE OR REPLACE FUNCTION apple_bid(_id TEXT)
  RETURNS VOID AS
$$BEGIN INSERT INTO apple_queue (id) VALUES ('@' || _id);
END;$$
LANGUAGE plpgsql VOLATILE;
COMMENT ON FUNCTION apple_bid(TEXT) IS '向苹果队列中添加BundleID任务';
-- SELECT apple_bid('com.tencent.xin');
-----------------------------------------
-- Function: add search keyword to queue
CREATE OR REPLACE FUNCTION apple_key(keyword TEXT)
  RETURNS VOID AS
$$BEGIN INSERT INTO apple_queue (id) VALUES ('#' || keyword);
END;$$
LANGUAGE plpgsql VOLATILE;
COMMENT ON FUNCTION apple_key(TEXT) IS '向苹果队列中添加关键词任务';
-- SELECT apple_key('蛤蛤');
-----------------------------------------


---------------------------------------------------------------


---------------------------------------------------------------
-- ORM Alternative choice: upsert manually
---------------------------------------------------------------
-- INSERT INTO apple (id, name, url, icon, kind, version, bundle_id,
--                    author_id, author_name, author_url, vendor_name, vendor_url, copyright,
--                    genre_id, genre_name, genres, genre_ids,
--                    icon60, icon100, price, currency,
--                    system, features, devices, languages, platforms, rating, reasons,
--                    size, cnt_rating, avg_rating, cnt_rating_current, avg_rating_current, vpp_device, game_center,
--                    screenshots, in_app_purchase, sibling_apps, related_apps, support_sites, reviews,
--                    description, release_notes, release_time, publish_time, crawled_time) VALUES
--   ($1, $2, $3, $4, $5, $6, $7, $8, $9,
--        $10, $11, $12, $13, $14, $15, $16, $17, $18, $19,
--                  $20, $21, $22, $23, $24, $25, $26, $27, $28, $29,
--                            $30, $31, $32, $33, $34, $35, $36, $37, $38, $39,
--                                      $40, $41, $42, $43, $44, $45, $46)
-- ON CONFLICT (id)
--   DO UPDATE SET
--
--     name               = EXCLUDED.name,
--     url                = EXCLUDED.url,
--     icon               = EXCLUDED.icon,
--     kind               = EXCLUDED.kind,
--     version            = EXCLUDED.version,
--     bundle_id          = EXCLUDED.bundle_id,
--     author_id          = EXCLUDED.author_id,
--     author_name        = EXCLUDED.author_name,
--     author_url         = EXCLUDED.author_url,
--     vendor_name        = EXCLUDED.vendor_name,
--     vendor_url         = EXCLUDED.vendor_url,
--     copyright          = EXCLUDED.copyright,
--     genre_id           = EXCLUDED.genre_id,
--     genre_name         = EXCLUDED.genre_name,
--     genres             = EXCLUDED.genres,
--     genre_ids          = EXCLUDED.genre_ids,
--     icon60             = EXCLUDED.icon60,
--     icon100            = EXCLUDED.icon100,
--     price              = EXCLUDED.price,
--     currency           = EXCLUDED.currency,
--     system             = EXCLUDED.system,
--     features           = EXCLUDED.features,
--     devices            = EXCLUDED.devices,
--     languages          = EXCLUDED.languages,
--     platforms          = EXCLUDED.platforms,
--     rating             = EXCLUDED.rating,
--     reasons            = EXCLUDED.reasons,
--     size               = EXCLUDED.size,
--     cnt_rating         = EXCLUDED.cnt_rating,
--     avg_rating         = EXCLUDED.avg_rating,
--     cnt_rating_current = EXCLUDED.cnt_rating_current,
--     avg_rating_current = EXCLUDED.avg_rating_current,
--     vpp_device         = EXCLUDED.vpp_device,
--     game_center        = EXCLUDED.game_center,
--     screenshots        = EXCLUDED.screenshots,
--     in_app_purchase    = EXCLUDED.in_app_purchase,
--     sibling_apps       = EXCLUDED.sibling_apps,
--     related_apps       = EXCLUDED.related_apps,
--     support_sites      = EXCLUDED.support_sites,
--     reviews            = EXCLUDED.reviews,
--     description        = EXCLUDED.description,
--     release_notes      = EXCLUDED.release_notes,
--     release_time       = EXCLUDED.release_time,
--     publish_time       = EXCLUDED.publish_time,
--     crawled_time       = EXCLUDED.crawled_time;
---------------------------------------------------------------


---------------------------------------------------------------
-- ETL Script
---------------------------------------------------------------
-- SELECT id, name, url, icon, kind, version, bundle_id, author_id, author_name, author_url, vendor_name, vendor_url, copyright, genre_id, array_to_json(genre_id_list) :: TEXT, genre, array_to_json(genre_list) :: TEXT, icon60, icon100, price, currency, system, array_to_json(features) :: TEXT, array_to_json(devices) :: TEXT, array_to_json(languages) :: TEXT, array_to_json(platforms) :: TEXT, rating, array_to_json(reasons) :: TEXT, size, cnt_rating, avg_rating, cnt_rating_current, avg_rating_current, vpp_device, game_center, array_to_json(screenshots) :: TEXT, array_to_json(in_app_purchase) :: TEXT, array_to_json(sibling_apps) :: TEXT, array_to_json(related_apps) :: TEXT, support_sites :: TEXT, reviews :: TEXT, extra :: TEXT, description, release_notes, release_time, publish_time, crawled_time FROM apple;
-----------------------------------------------------------------


---------------------------------------------------------------
-- MaxCompute DDL
---------------------------------------------------------------
-- CREATE TABLE apple (
--   id                 BIGINT COMMENT '应用ID，又名iTunesID,trackID,整型',
--   name               STRING COMMENT '应用名称',
--   url                STRING COMMENT '应用页面URL',
--   icon               STRING COMMENT '应用图标URL(512px)',
--   kind               STRING COMMENT '应用类型(software)',
--   version            STRING COMMENT '应用版本',
--   bundle_id          STRING COMMENT '应用包名BundleID',
--   author_id          BIGINT COMMENT '作者ID',
--   author_name        STRING COMMENT '作者名',
--   author_url         STRING COMMENT '作者页面URL',
--   vendor_name        STRING COMMENT '厂商名称',
--   vendor_url         STRING COMMENT '厂商URL(外部)',
--   copyright          STRING COMMENT '版权信息',
--   genre_id           BIGINT COMMENT '首要类型ID',
--   genre_id_list      STRING COMMENT '类型ID数组',
--   genre              STRING COMMENT '首要类型名称',
--   genre_list         STRING COMMENT '类型名称数组',
--   icon60             STRING COMMENT '图标URL(60px)',
--   icon100            STRING COMMENT '图标URL(100px)',
--   price              BIGINT COMMENT '价格，0为免费',
--   currency           STRING COMMENT '货币',
--   system             STRING COMMENT '系统版本要求',
--   features           STRING COMMENT '特性，通常为iosUniversal',
--   devices            STRING COMMENT '支持的设备列表',
--   languages          STRING COMMENT '支持的语言，ISO2字符代号，已排序',
--   platforms          STRING COMMENT '支持的平台:iPad,iPhone,iPod,iWatch,AppleTV',
--   rating             STRING COMMENT '应用分级',
--   reasons            STRING COMMENT '分级原因',
--   size               BIGINT COMMENT '文件大小',
--   cnt_rating         BIGINT COMMENT '评分人数(历史所有版本)',
--   avg_rating         DOUBLE COMMENT '平均评分(历史所有版本)',
--   cnt_rating_current BIGINT COMMENT '评分人数(当前版本)',
--   avg_rating_current DOUBLE COMMENT '平均评分(当前版本)',
--   vpp_device         BOOLEAN COMMENT '是否支持VppDevice',
--   game_center        BOOLEAN COMMENT '是否启用游戏中心',
--   screenshots        STRING COMMENT '截图[URL数组]',
--   in_app_purchase    STRING COMMENT '应用内购`(rank,price,title)`',
--   sibling_apps       STRING COMMENT '同一开发者的应用',
--   related_apps       STRING COMMENT '苹果推荐的应用',
--   support_sites      STRING COMMENT '应用支持站点，KV JSON Object',
--   reviews            STRING COMMENT '用户评论`(user,rating,title,content)`',
--   extra              STRING COMMENT '额外信息，JSONB kv格式',
--   description        STRING COMMENT '应用描述',
--   release_notes      STRING COMMENT '最近更新内容',
--   release_time       DATETIME COMMENT '更新时间',
--   publish_time       DATETIME COMMENT '初始发布时间',
--   crawled_time       DATETIME COMMENT '抓取时间'
-- ) PARTITIONED BY (ds string);
---------------------------------------------------------------