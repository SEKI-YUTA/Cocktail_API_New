--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4 (Debian 15.4-1.pgdg120+1)
-- Dumped by pg_dump version 16.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cocktail_categories; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.cocktail_categories (
    cocktail_category_id integer NOT NULL,
    name character varying(100)
);


ALTER TABLE public.cocktail_categories OWNER TO root;

--
-- Name: cocktail_categories_cocktail_category_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.cocktail_categories_cocktail_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cocktail_categories_cocktail_category_id_seq OWNER TO root;

--
-- Name: cocktail_categories_cocktail_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cocktail_categories_cocktail_category_id_seq OWNED BY public.cocktail_categories.cocktail_category_id;


--
-- Name: cocktail_ingredients; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.cocktail_ingredients (
    cocktail_ingredient_id integer NOT NULL,
    cocktail_id integer,
    ingredient_id integer
);


ALTER TABLE public.cocktail_ingredients OWNER TO root;

--
-- Name: cocktail_ingredients_cocktail_ingredient_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.cocktail_ingredients_cocktail_ingredient_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cocktail_ingredients_cocktail_ingredient_id_seq OWNER TO root;

--
-- Name: cocktail_ingredients_cocktail_ingredient_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cocktail_ingredients_cocktail_ingredient_id_seq OWNED BY public.cocktail_ingredients.cocktail_ingredient_id;


--
-- Name: cocktail_parent_child; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.cocktail_parent_child (
    cocktail_parent_child_id integer NOT NULL,
    parent_id integer,
    child_id integer
);


ALTER TABLE public.cocktail_parent_child OWNER TO root;

--
-- Name: cocktail_parent_child_cocktail_parent_child_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.cocktail_parent_child_cocktail_parent_child_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cocktail_parent_child_cocktail_parent_child_id_seq OWNER TO root;

--
-- Name: cocktail_parent_child_cocktail_parent_child_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cocktail_parent_child_cocktail_parent_child_id_seq OWNED BY public.cocktail_parent_child.cocktail_parent_child_id;


--
-- Name: cocktails; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.cocktails (
    cocktail_id integer NOT NULL,
    name character varying(100),
    description character varying(1000) DEFAULT ''::character varying,
    cocktail_category_id integer,
    parent_cocktail_id integer,
    vol integer DEFAULT 0,
    ingredient_count integer DEFAULT 0
);


ALTER TABLE public.cocktails OWNER TO root;

--
-- Name: cocktails_cocktail_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.cocktails_cocktail_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cocktails_cocktail_id_seq OWNER TO root;

--
-- Name: cocktails_cocktail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cocktails_cocktail_id_seq OWNED BY public.cocktails.cocktail_id;


--
-- Name: ingredient_categories; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.ingredient_categories (
    ingredient_category_id integer NOT NULL,
    name character varying(100)
);


ALTER TABLE public.ingredient_categories OWNER TO root;

--
-- Name: ingredient_categories_ingredient_category_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.ingredient_categories_ingredient_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ingredient_categories_ingredient_category_id_seq OWNER TO root;

--
-- Name: ingredient_categories_ingredient_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.ingredient_categories_ingredient_category_id_seq OWNED BY public.ingredient_categories.ingredient_category_id;


--
-- Name: ingredients; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.ingredients (
    ingredient_id integer NOT NULL,
    shortname character varying(100) DEFAULT ''::character varying,
    longname character varying(100) DEFAULT ''::character varying,
    description character varying(1000) DEFAULT ''::character varying,
    vol integer DEFAULT 0,
    ingredient_category_id integer
);


ALTER TABLE public.ingredients OWNER TO root;

--
-- Name: ingredients_ingredient_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.ingredients_ingredient_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ingredients_ingredient_id_seq OWNER TO root;

--
-- Name: ingredients_ingredient_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.ingredients_ingredient_id_seq OWNED BY public.ingredients.ingredient_id;


--
-- Name: cocktail_categories cocktail_category_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktail_categories ALTER COLUMN cocktail_category_id SET DEFAULT nextval('public.cocktail_categories_cocktail_category_id_seq'::regclass);


--
-- Name: cocktail_ingredients cocktail_ingredient_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktail_ingredients ALTER COLUMN cocktail_ingredient_id SET DEFAULT nextval('public.cocktail_ingredients_cocktail_ingredient_id_seq'::regclass);


--
-- Name: cocktail_parent_child cocktail_parent_child_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktail_parent_child ALTER COLUMN cocktail_parent_child_id SET DEFAULT nextval('public.cocktail_parent_child_cocktail_parent_child_id_seq'::regclass);


--
-- Name: cocktails cocktail_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktails ALTER COLUMN cocktail_id SET DEFAULT nextval('public.cocktails_cocktail_id_seq'::regclass);


--
-- Name: ingredient_categories ingredient_category_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.ingredient_categories ALTER COLUMN ingredient_category_id SET DEFAULT nextval('public.ingredient_categories_ingredient_category_id_seq'::regclass);


--
-- Name: ingredients ingredient_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.ingredients ALTER COLUMN ingredient_id SET DEFAULT nextval('public.ingredients_ingredient_id_seq'::regclass);


--
-- Data for Name: cocktail_categories; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.cocktail_categories (cocktail_category_id, name) FROM stdin;
1	ショートカクテル
2	ロングカクテル
3	ノンアルコール
\.


--
-- Data for Name: cocktail_ingredients; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.cocktail_ingredients (cocktail_ingredient_id, cocktail_id, ingredient_id) FROM stdin;
1	1	1
2	1	2
3	2	1
4	2	3
5	3	4
6	3	5
7	4	4
8	4	2
9	4	6
10	5	7
11	5	2
12	6	8
13	6	2
14	7	9
15	7	10
16	7	11
17	8	1
18	8	10
19	8	11
20	9	2
21	9	12
\.


--
-- Data for Name: cocktail_parent_child; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.cocktail_parent_child (cocktail_parent_child_id, parent_id, child_id) FROM stdin;
\.


--
-- Data for Name: cocktails; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.cocktails (cocktail_id, name, description, cocktail_category_id, parent_cocktail_id, vol, ingredient_count) FROM stdin;
1	ジンバッグ	ジンベースのカクテル	2	\N	10	2
2	ジントニック	ジンベースのカクテル	2	\N	10	2
3	スクリュードライバー	ウォッカベースのカクテル	2	\N	10	2
4	モスコミュール	ウォッカベースのカクテル	2	\N	10	3
5	ラムバック		2	\N	10	2
6	ラムバック		2	\N	10	2
7	サイドカー		1	\N	20	3
9	シャーリーテンプル	ノンアルコールのカクテル	3	\N	0	2
8	ホワイトレディー	ジンベースのカクテル	1	7	20	3
\.


--
-- Data for Name: ingredient_categories; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.ingredient_categories (ingredient_category_id, name) FROM stdin;
1	スピリッツ
2	リキュール
3	ソフトドリンク
4	ウィスキー
5	ブランデー
6	焼酎
7	日本酒
\.


--
-- Data for Name: ingredients; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.ingredients (ingredient_id, shortname, longname, description, vol, ingredient_category_id) FROM stdin;
1	ジン	ドライ・ジン		40	1
2	ジンジャーエール	ジンジャーエール	ジンジャーの味がする炭酸飲料	0	3
3	トニックウォーター	トニックウォーター		0	3
4	ウォッカ	ウォッカ		40	1
5	オレンジジュース	オレンジジュース		0	3
6	ライムジュース	ライムジュース		0	3
7	ラム	ホワイトラム		40	1
8	ラム	ダークラム		40	1
9	ブランデー	コニャック		40	5
10	キュラソー	ホワイトキュラソー		40	1
11	レモンジュース	レモンジュース		0	3
12	グレナデンシロップ	グレナデンシロップ		0	3
\.


--
-- Name: cocktail_categories_cocktail_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.cocktail_categories_cocktail_category_id_seq', 3, true);


--
-- Name: cocktail_ingredients_cocktail_ingredient_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.cocktail_ingredients_cocktail_ingredient_id_seq', 21, true);


--
-- Name: cocktail_parent_child_cocktail_parent_child_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.cocktail_parent_child_cocktail_parent_child_id_seq', 1, false);


--
-- Name: cocktails_cocktail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.cocktails_cocktail_id_seq', 9, true);


--
-- Name: ingredient_categories_ingredient_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.ingredient_categories_ingredient_category_id_seq', 7, true);


--
-- Name: ingredients_ingredient_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.ingredients_ingredient_id_seq', 12, true);


--
-- Name: cocktail_categories cocktail_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktail_categories
    ADD CONSTRAINT cocktail_categories_pkey PRIMARY KEY (cocktail_category_id);


--
-- Name: cocktail_ingredients cocktail_ingredients_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktail_ingredients
    ADD CONSTRAINT cocktail_ingredients_pkey PRIMARY KEY (cocktail_ingredient_id);


--
-- Name: cocktail_parent_child cocktail_parent_child_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktail_parent_child
    ADD CONSTRAINT cocktail_parent_child_pkey PRIMARY KEY (cocktail_parent_child_id);


--
-- Name: cocktails cocktails_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktails
    ADD CONSTRAINT cocktails_pkey PRIMARY KEY (cocktail_id);


--
-- Name: ingredient_categories ingredient_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.ingredient_categories
    ADD CONSTRAINT ingredient_categories_pkey PRIMARY KEY (ingredient_category_id);


--
-- Name: ingredients ingredients_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.ingredients
    ADD CONSTRAINT ingredients_pkey PRIMARY KEY (ingredient_id);


--
-- Name: cocktail_ingredients cocktail_ingredients_cocktail_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktail_ingredients
    ADD CONSTRAINT cocktail_ingredients_cocktail_id_fkey FOREIGN KEY (cocktail_id) REFERENCES public.cocktails(cocktail_id);


--
-- Name: cocktail_ingredients cocktail_ingredients_ingredient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktail_ingredients
    ADD CONSTRAINT cocktail_ingredients_ingredient_id_fkey FOREIGN KEY (ingredient_id) REFERENCES public.ingredients(ingredient_id);


--
-- Name: cocktail_parent_child cocktail_parent_child_child_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktail_parent_child
    ADD CONSTRAINT cocktail_parent_child_child_id_fkey FOREIGN KEY (child_id) REFERENCES public.cocktails(cocktail_id);


--
-- Name: cocktail_parent_child cocktail_parent_child_parent_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktail_parent_child
    ADD CONSTRAINT cocktail_parent_child_parent_id_fkey FOREIGN KEY (parent_id) REFERENCES public.cocktails(cocktail_id);


--
-- Name: cocktails cocktails_cocktail_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktails
    ADD CONSTRAINT cocktails_cocktail_category_id_fkey FOREIGN KEY (cocktail_category_id) REFERENCES public.cocktail_categories(cocktail_category_id);


--
-- Name: cocktails cocktails_parent_cocktail_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cocktails
    ADD CONSTRAINT cocktails_parent_cocktail_id_fkey FOREIGN KEY (parent_cocktail_id) REFERENCES public.cocktails(cocktail_id);


--
-- Name: ingredients ingredients_ingredient_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.ingredients
    ADD CONSTRAINT ingredients_ingredient_category_id_fkey FOREIGN KEY (ingredient_category_id) REFERENCES public.ingredient_categories(ingredient_category_id);


--
-- PostgreSQL database dump complete
--

