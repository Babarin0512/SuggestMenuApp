#冷蔵庫内の食材の種類を示すカラムを登録(重複がないように登録する)
CREATE TABLE shokuzai_refrigerator(
    shokuzai_id int UNIQUE
);

#料理名,材料id,料理に必要な材料の合計を登録する。
CREATE TABLE app_db(
    ryouri_name VARCHAR(20),
    zairyou_id int,
    zairyou_sum int
);

INSERT INTO shokuzai_refrigerator
(shokuzai_id)
VALUES
(001)
(002)
(003)
(004)
(005)
(006)
(007)
(008)
(009)
 
INSERT INTO app_db
(ryouri_name, zairyou_id, zairyou_sum)
VALUES
('冷奴', 001, 2)
('冷奴', 002, 2)
('ステーキ', 001, 2)
('ちゃんこ鍋', 11, 1);

