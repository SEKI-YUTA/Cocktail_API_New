-- 参考にしたサイト
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
