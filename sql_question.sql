-- 題目1

--- 使用MySQL,InnoDB

--- 建立 property table
CREATE TABLE `property` (
  `id` int NOT NULL COMMENT '旅宿ID' AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) COLLATE 'utf8_general_ci' NOT NULL COMMENT '旅宿名稱'
) ENGINE='InnoDB' COLLATE 'utf8_general_ci';

--- 建立 room table
CREATE TABLE `room` (
  `id` int NOT NULL COMMENT '房間ID' AUTO_INCREMENT PRIMARY KEY,
  `property_id` int NOT NULL COMMENT '旅宿ID',
  `name` varchar(255) COLLATE 'utf8_general_ci' NOT NULL COMMENT '房間名稱'
) ENGINE='InnoDB' COLLATE 'utf8_general_ci';

--- 建立 orders table
CREATE TABLE `orders` (
  `id` int NOT NULL COMMENT '訂單ID' AUTO_INCREMENT PRIMARY KEY,
  `room_id` int NOT NULL COMMENT '房間ID',
  `price` decimal(20,4) NOT NULL COMMENT '價格',
  `created_at` timestamp NOT NULL COMMENT '建立時間'
) ENGINE='InnoDB' COLLATE 'utf8_general_ci';

--- 查詢語法
SELECT 
    COUNT(p.id) AS `order_count`, 
    p.id AS `property_id`,
    p.name AS `property_name` 
FROM `orders` o
JOIN room as r ON o.room_id = r.id
JOIN property as p ON r.property_id = p.id
WHERE o.created_at >= '2021-02-01 00:00:00' AND o.created_at < '2021-02-28 23:59:59'
GROUP BY p.id 
Order By order_count DESC,p.id DESC;

--- 調整資料表
-- 1. 視情況設定property room order的FK 來確保資料的正確性
-- 2. order.created_at 加上index (時間區間查詢應該很常使用到)
-- 3. 將 property.id 加入 order table
-- 4. 考慮加上 deleted_at 可進行軟刪除

--- 題目2
-- 查詢過慢可能原因
-- 1. 沒有設定index
-- 2. 語法沒有hit到index
-- 3. 語法調整

-- 解決方式
-- 1. 檢查index設定的方式好不好，如需調整則調整index
-- 2. 查詢過慢的語法可以下 slow query檢查並優化
-- 3. 是否因讀寫鎖問題造成鎖表，可考慮讀寫分離
-- 4. 非經常更新的資料，可於程式做快取機制
-- 5. 資料表是否有需要切partition
