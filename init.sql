--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5
-- Dumped by pg_dump version 14.5

-- Started on 2022-11-06 23:13:53

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

DROP DATABASE jokeapi;
--
-- TOC entry 3354 (class 1262 OID 32798)
-- Name: jokeapi; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE jokeapi WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';


ALTER DATABASE jokeapi OWNER TO postgres;

\connect jokeapi

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

--
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3355 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 211 (class 1259 OID 32801)
-- Name: jokes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.jokes (
    id bigint NOT NULL,
    text character varying(255) NOT NULL,
    author_id character varying(255) NOT NULL,
    description character varying(255),
    lang character varying(6) NOT NULL,
    uuid character varying(36) NOT NULL,
    added_at timestamp without time zone NOT NULL,
    modified_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.jokes OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 32800)
-- Name: jokes_author_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.jokes_author_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.jokes_author_id_seq OWNER TO postgres;

--
-- TOC entry 3356 (class 0 OID 0)
-- Dependencies: 210
-- Name: jokes_author_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.jokes_author_id_seq OWNED BY public.jokes.author_id;


--
-- TOC entry 209 (class 1259 OID 32799)
-- Name: jokes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.jokes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.jokes_id_seq OWNER TO postgres;

--
-- TOC entry 3357 (class 0 OID 0)
-- Dependencies: 209
-- Name: jokes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.jokes_id_seq OWNED BY public.jokes.id;


--
-- TOC entry 215 (class 1259 OID 32898)
-- Name: ratings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ratings (
    id bigint NOT NULL,
    stars numeric(5,3) NOT NULL,
    comment character varying(255),
    user_id character varying(36),
    joke_id character varying(36) NOT NULL,
    added_at timestamp without time zone NOT NULL,
    modified_at timestamp without time zone,
    deleted_at timestamp without time zone,
    uuid character varying(36)
);


ALTER TABLE public.ratings OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 32897)
-- Name: ratings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ratings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ratings_id_seq OWNER TO postgres;

--
-- TOC entry 3358 (class 0 OID 0)
-- Dependencies: 214
-- Name: ratings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ratings_id_seq OWNED BY public.ratings.id;


--
-- TOC entry 213 (class 1259 OID 32811)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    email character varying(120) NOT NULL,
    hash character varying(255) NOT NULL,
    uuid character varying(36) NOT NULL,
    added_at timestamp without time zone NOT NULL,
    modified_at timestamp without time zone,
    deleted_at timestamp without time zone,
    username character varying(64) NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 32810)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3359 (class 0 OID 0)
-- Dependencies: 212
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3175 (class 2604 OID 32804)
-- Name: jokes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jokes ALTER COLUMN id SET DEFAULT nextval('public.jokes_id_seq'::regclass);


--
-- TOC entry 3177 (class 2604 OID 32901)
-- Name: ratings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ratings ALTER COLUMN id SET DEFAULT nextval('public.ratings_id_seq'::regclass);


--
-- TOC entry 3176 (class 2604 OID 32814)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3344 (class 0 OID 32801)
-- Dependencies: 211
-- Data for Name: jokes; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.jokes VALUES (1, 'Who’s there?

Interrupting cow.

Interrupting

c–

MOO!', 'bb002697-0911-4866-8000-e1ff2ef992a6', 'ladlflñ', 'en-US', 'ac70acfd-cd3a-469e-8483-ca681dac44bb', '2022-11-02 02:46:05.423063', NULL, NULL);
INSERT INTO public.jokes VALUES (3, 'If you’re American when you go in the bathroom…

… and American when you come out, what are you in the bathroom?', 'bb002697-0911-4866-8000-e1ff2ef992a6', '', 'en-US', '3118ed8f-8cd5-410a-819c-9d2dc8a65be0', '2022-11-02 03:01:50.073954', NULL, NULL);
INSERT INTO public.jokes VALUES (2, '— Oiga, ¿el otorrino va por número?

— Van nombrando.

— Qué gran actor, pero no me cambie de tema.', 'bb002697-0911-4866-8000-e1ff2ef992a6', '', 'es-LA', '975dddb2-94ef-4376-9c71-380b47003488', '2022-11-02 02:55:11.047638', '2022-11-02 21:27:37.713483', NULL);


--
-- TOC entry 3348 (class 0 OID 32898)
-- Dependencies: 215
-- Data for Name: ratings; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.ratings VALUES (3, 2.000, 'Meh', '4e7fecd7-617d-42b1-a6fa-696131212c31', 'ac70acfd-cd3a-469e-8483-ca681dac44bb', '2022-11-02 15:08:01.015782', NULL, NULL, '0162bc39-d361-45e8-8a4b-c0236c00c5eb');
INSERT INTO public.ratings VALUES (2, 5.000, '', 'bb002697-0911-4866-8000-e1ff2ef992a6', 'ac70acfd-cd3a-469e-8483-ca681dac44bb', '2022-11-02 15:02:43.278682', NULL, NULL, '23c8f0ad-774a-42c9-b50f-46b55d8f7645');


--
-- TOC entry 3346 (class 0 OID 32811)
-- Dependencies: 213
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES (1, 'devasodin344@gmail.com', '$2a$10$9ZK3egf51g6PUf/gC1yl4.Y1HquCkF8kY86yiDVVNJQkm8r50a9Sm', 'bb002697-0911-4866-8000-e1ff2ef992a6', '2022-11-01 05:27:02.192625', NULL, NULL, 'devasodin');
INSERT INTO public.users VALUES (2, 'davidquinterogranadillo@gmail.com', '$2a$10$kacNxOlwibjUTLyeQAt5bu3w6ek4YVkYLnNb1S6W2gkwMKvtiDp4W', '4e7fecd7-617d-42b1-a6fa-696131212c31', '2022-11-01 15:51:12.519441', NULL, NULL, 'davidqui');


--
-- TOC entry 3360 (class 0 OID 0)
-- Dependencies: 210
-- Name: jokes_author_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.jokes_author_id_seq', 1, false);


--
-- TOC entry 3361 (class 0 OID 0)
-- Dependencies: 209
-- Name: jokes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.jokes_id_seq', 3, true);


--
-- TOC entry 3362 (class 0 OID 0)
-- Dependencies: 214
-- Name: ratings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ratings_id_seq', 3, true);


--
-- TOC entry 3363 (class 0 OID 0)
-- Dependencies: 212
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- TOC entry 3181 (class 2606 OID 32809)
-- Name: jokes jokes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jokes
    ADD CONSTRAINT jokes_pkey PRIMARY KEY (id);


--
-- TOC entry 3183 (class 2606 OID 32911)
-- Name: jokes jokes_unique_uuid; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jokes
    ADD CONSTRAINT jokes_unique_uuid UNIQUE (uuid);


--
-- TOC entry 3196 (class 2606 OID 32921)
-- Name: ratings ratings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT ratings_pk PRIMARY KEY (id);


--
-- TOC entry 3198 (class 2606 OID 32918)
-- Name: ratings ratings_un; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT ratings_un UNIQUE (uuid);


--
-- TOC entry 3185 (class 2606 OID 32818)
-- Name: users unique_email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT unique_email UNIQUE (email);


--
-- TOC entry 3188 (class 2606 OID 32816)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3190 (class 2606 OID 32862)
-- Name: users users_unique_uuid; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_unique_uuid UNIQUE (uuid);


--
-- TOC entry 3178 (class 1259 OID 32860)
-- Name: jokes_added_at_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX jokes_added_at_idx ON public.jokes USING btree (added_at);


--
-- TOC entry 3179 (class 1259 OID 32859)
-- Name: jokes_author_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX jokes_author_id_idx ON public.jokes USING btree (author_id);


--
-- TOC entry 3193 (class 1259 OID 32919)
-- Name: ratings_added_at_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ratings_added_at_idx ON public.ratings USING btree (added_at);


--
-- TOC entry 3194 (class 1259 OID 32903)
-- Name: ratings_joke_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ratings_joke_id_idx ON public.ratings USING btree (joke_id);


--
-- TOC entry 3199 (class 1259 OID 32902)
-- Name: ratings_user_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ratings_user_id_idx ON public.ratings USING btree (user_id);


--
-- TOC entry 3186 (class 1259 OID 32821)
-- Name: users_added_at_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX users_added_at_idx ON public.users USING btree (added_at);


--
-- TOC entry 3191 (class 1259 OID 32822)
-- Name: users_username_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_username_idx ON public.users USING btree (username);


--
-- TOC entry 3192 (class 1259 OID 32863)
-- Name: users_uuid_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_uuid_idx ON public.users USING btree (uuid);


--
-- TOC entry 3200 (class 2606 OID 32864)
-- Name: jokes jokes_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jokes
    ADD CONSTRAINT jokes_fk FOREIGN KEY (author_id) REFERENCES public.users(uuid);


--
-- TOC entry 3201 (class 2606 OID 32905)
-- Name: ratings ratings_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT ratings_fk FOREIGN KEY (user_id) REFERENCES public.users(uuid);


--
-- TOC entry 3202 (class 2606 OID 32912)
-- Name: ratings ratings_fk_jokes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT ratings_fk_jokes FOREIGN KEY (joke_id) REFERENCES public.jokes(uuid);


-- Completed on 2022-11-06 23:13:54

--
-- PostgreSQL database dump complete
--

