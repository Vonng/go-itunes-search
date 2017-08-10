CREATE TABLE apple (
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