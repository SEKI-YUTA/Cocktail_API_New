-- https://bipinparajuli.com.np/blog/many-to-many-relationship-in-postgresql
CREATE DATABASE cocktail_db;
\c cocktail_db

CREATE TABLE ingredient_categories(
    ingredient_category_id serial not null,
    name varchar(100),
    primary key(ingredient_category_id)
);

CREATE TABLE cocktail_categories(
    cocktail_category_id serial not null,
    name varchar(100),
    primary key(cocktail_category_id)
);

CREATE TABLE ingredients(
    ingredient_id serial not null,
    shortname varchar(100) DEFAULT '',
    longname varchar(100) DEFAULT '',
    -- longName varchar(100), バーボンウィスキー
    -- shortName varchar(100), ウィスキー
    -- category varchar(100), リキュール、ソフトドリンク、その他
    description varchar(1000) DEFAULT '',
    vol int DEFAULT 0,
    ingredient_category_id int,
    foreign key(ingredient_category_id) references ingredient_categories(ingredient_category_id),
    primary key(ingredient_id)
);

CREATE TABLE cocktails(
    cocktail_id serial not null,
    name varchar(100),
    description varchar(1000) DEFAULT '',
    cocktail_category_id int,
    foreign key(cocktail_category_id) references cocktail_categories(cocktail_category_id),
    parent_cocktail_id int,
    foreign key(parent_cocktail_id) references cocktails(cocktail_id),
    -- category varchar(100), ショート、ロング、その他
    -- parent 派生元がある場合、ホワイトレディ　-> サイドカーみたいな
    -- child 派生先がある場合、サイドカー　-> ホワイトレディみたいな
    vol int DEFAULT 0,
    ingredient_count int DEFAULT 0,
    primary key(cocktail_id)
);

CREATE TABLE cocktail_ingredients(
    cocktail_ingredient_id serial not null,
    cocktail_id int,
    ingredient_id int,
    foreign key(cocktail_id) references cocktails(cocktail_id),
    foreign key(ingredient_id) references ingredients(ingredient_id),
    -- CONSTRAINT cocktail_ingredients_pk PRIMARY KEY (cocktail_id, ingredient_id)
    primary key(cocktail_ingredient_id)
);

CREATE TABLE cocktail_parent_child(
    cocktail_parent_child_id serial not null,
    parent_id int,
    child_id int,
    foreign key(parent_id) references cocktails(cocktail_id),
    foreign key(child_id) references cocktails(cocktail_id),
    primary key(cocktail_parent_child_id)
);

-- insert into ingredients (name, isAlcohol) values ('ジン', TRUE);
-- insert into ingredients (name, isAlcohol) values ('ウォッカ', TRUE);
-- insert into ingredients (name, isAlcohol) values ('テキーラ', TRUE);
-- insert into ingredients (name, isAlcohol) values ('ラム', TRUE);
-- insert into ingredients (name, isAlcohol) values ('ファジーネーブル', TRUE);
-- insert into ingredients (name, isAlcohol) values ('カシス', TRUE);
-- insert into ingredients (name, isAlcohol) values ('ウィスキー', TRUE);
-- insert into ingredients (name, isAlcohol) values ('梅酒', TRUE);
-- insert into ingredients (name, isAlcohol) values ('焼酎', TRUE);
-- insert into ingredients (name, isAlcohol) values ('日本酒', TRUE);
-- insert into ingredients (name, isAlcohol) values ('グレナディンシロップ', FALSE);
-- insert into ingredients (name, isAlcohol) values ('ジンジャーエール', FALSE);
-- insert into ingredients (name, isAlcohol) values ('オレンジジュース', FALSE);
-- insert into ingredients (name, isAlcohol) values ('ライムジュース', FALSE);
-- insert into ingredients (name, isAlcohol) values ('烏龍茶', FALSE);
-- insert into ingredients (name, isAlcohol) values ('カルピス', FALSE);


-- insert into cocktails (name) values ('ジンバッグ');
-- insert into cocktails (name) values ('ウォッカバッグ');
-- insert into cocktails (name) values ('ラムバッグ');
-- insert into cocktails (name) values ('モスコミュール');
-- insert into cocktails (name) values ('テキーラサンライズ');
-- insert into cocktails (name) values ('ウーロンハイ');
-- insert into cocktails (name) values ('カルピスハイ');

-- insert into cocktail_ingredients (cocktail_id, ingredient_id) values (1, 1);
-- insert into cocktail_ingredients (cocktail_id, ingredient_id) values (1, 12);
-- insert into cocktail_ingredients (cocktail_id, ingredient_id) values (2, 2);
-- insert into cocktail_ingredients (cocktail_id, ingredient_id) values (2, 12);
-- insert into cocktail_ingredients (cocktail_id, ingredient_id) values (3, 4);
-- insert into cocktail_ingredients (cocktail_id, ingredient_id) values (3, 12);

-- ジャンクションテーブルの中身を確認
-- 中間テーブルイコールジャンクションテーブル？
--  select * from cocktail_ingredients INNER JOIN ingredients ON cocktail_ingredients.ingredient_id = ingredients.ingredient_id;
-- select i.name as name from cocktail_ingredients INNER JOIN ingredients i ON cocktail_ingredients.ingredient_id = i.ingredient_id;
-- select * from cocktail_ingredients INNER JOIN ingredients i ON cocktail_ingredients.ingredient_id = i.ingredient_id where cocktail_ingredients.cocktail_id=1;
-- カクテルとカクテルの素材をセットで取得
-- select * from cocktail_ingredients INNER JOIN ingredients i ON cocktail_ingredients.ingredient_id = i.ingredient_id INNER JOIN cocktails c ON cocktail_ingredients.cocktail_id = c.cocktail_id;
