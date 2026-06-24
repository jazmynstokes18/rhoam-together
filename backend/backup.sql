--
-- PostgreSQL database dump
--

\restrict GSgquiNM8yvJaFR2mb1OEA01ILDeBVeprcUXNvmua77pmQ8e0TaWbHxFdy7Ymil

-- Dumped from database version 15.18 (Homebrew)
-- Dumped by pg_dump version 15.18 (Homebrew)

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
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: update_updated_at_column(); Type: FUNCTION; Schema: public; Owner: jazmynstokes
--

CREATE FUNCTION public.update_updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_updated_at_column() OWNER TO jazmynstokes;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: suggestions; Type: TABLE; Schema: public; Owner: jazmynstokes
--

CREATE TABLE public.suggestions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    trip_id uuid NOT NULL,
    created_by uuid NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    suggested_date date,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.suggestions OWNER TO jazmynstokes;

--
-- Name: trip_members; Type: TABLE; Schema: public; Owner: jazmynstokes
--

CREATE TABLE public.trip_members (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    trip_id uuid NOT NULL,
    user_id uuid NOT NULL,
    access_level character varying(50) NOT NULL,
    joined_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT trip_members_access_level_check CHECK (((access_level)::text = ANY ((ARRAY['read_only'::character varying, 'editor'::character varying, 'admin'::character varying])::text[])))
);


ALTER TABLE public.trip_members OWNER TO jazmynstokes;

--
-- Name: trips; Type: TABLE; Schema: public; Owner: jazmynstokes
--

CREATE TABLE public.trips (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    start_date date NOT NULL,
    end_date date NOT NULL,
    created_by uuid NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.trips OWNER TO jazmynstokes;

--
-- Name: users; Type: TABLE; Schema: public; Owner: jazmynstokes
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO jazmynstokes;

--
-- Data for Name: suggestions; Type: TABLE DATA; Schema: public; Owner: jazmynstokes
--

COPY public.suggestions (id, trip_id, created_by, title, description, suggested_date, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: trip_members; Type: TABLE DATA; Schema: public; Owner: jazmynstokes
--

COPY public.trip_members (id, trip_id, user_id, access_level, joined_at) FROM stdin;
\.


--
-- Data for Name: trips; Type: TABLE DATA; Schema: public; Owner: jazmynstokes
--

COPY public.trips (id, name, description, start_date, end_date, created_by, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: jazmynstokes
--

COPY public.users (id, email, password_hash, name, created_at, updated_at) FROM stdin;
\.


--
-- Name: suggestions suggestions_pkey; Type: CONSTRAINT; Schema: public; Owner: jazmynstokes
--

ALTER TABLE ONLY public.suggestions
    ADD CONSTRAINT suggestions_pkey PRIMARY KEY (id);


--
-- Name: trip_members trip_members_pkey; Type: CONSTRAINT; Schema: public; Owner: jazmynstokes
--

ALTER TABLE ONLY public.trip_members
    ADD CONSTRAINT trip_members_pkey PRIMARY KEY (id);


--
-- Name: trip_members trip_members_trip_id_user_id_key; Type: CONSTRAINT; Schema: public; Owner: jazmynstokes
--

ALTER TABLE ONLY public.trip_members
    ADD CONSTRAINT trip_members_trip_id_user_id_key UNIQUE (trip_id, user_id);


--
-- Name: trips trips_pkey; Type: CONSTRAINT; Schema: public; Owner: jazmynstokes
--

ALTER TABLE ONLY public.trips
    ADD CONSTRAINT trips_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: jazmynstokes
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: jazmynstokes
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_suggestions_created_by; Type: INDEX; Schema: public; Owner: jazmynstokes
--

CREATE INDEX idx_suggestions_created_by ON public.suggestions USING btree (created_by);


--
-- Name: idx_suggestions_suggested_date; Type: INDEX; Schema: public; Owner: jazmynstokes
--

CREATE INDEX idx_suggestions_suggested_date ON public.suggestions USING btree (suggested_date);


--
-- Name: idx_suggestions_trip_id; Type: INDEX; Schema: public; Owner: jazmynstokes
--

CREATE INDEX idx_suggestions_trip_id ON public.suggestions USING btree (trip_id);


--
-- Name: idx_trip_members_trip_id; Type: INDEX; Schema: public; Owner: jazmynstokes
--

CREATE INDEX idx_trip_members_trip_id ON public.trip_members USING btree (trip_id);


--
-- Name: idx_trip_members_user_id; Type: INDEX; Schema: public; Owner: jazmynstokes
--

CREATE INDEX idx_trip_members_user_id ON public.trip_members USING btree (user_id);


--
-- Name: idx_trips_created_by; Type: INDEX; Schema: public; Owner: jazmynstokes
--

CREATE INDEX idx_trips_created_by ON public.trips USING btree (created_by);


--
-- Name: idx_trips_dates; Type: INDEX; Schema: public; Owner: jazmynstokes
--

CREATE INDEX idx_trips_dates ON public.trips USING btree (start_date, end_date);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: jazmynstokes
--

CREATE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: suggestions suggestions_updated_at_trigger; Type: TRIGGER; Schema: public; Owner: jazmynstokes
--

CREATE TRIGGER suggestions_updated_at_trigger BEFORE UPDATE ON public.suggestions FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- Name: trips trips_updated_at_trigger; Type: TRIGGER; Schema: public; Owner: jazmynstokes
--

CREATE TRIGGER trips_updated_at_trigger BEFORE UPDATE ON public.trips FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- Name: users users_updated_at_trigger; Type: TRIGGER; Schema: public; Owner: jazmynstokes
--

CREATE TRIGGER users_updated_at_trigger BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- Name: suggestions suggestions_created_by_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jazmynstokes
--

ALTER TABLE ONLY public.suggestions
    ADD CONSTRAINT suggestions_created_by_fkey FOREIGN KEY (created_by) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: suggestions suggestions_trip_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jazmynstokes
--

ALTER TABLE ONLY public.suggestions
    ADD CONSTRAINT suggestions_trip_id_fkey FOREIGN KEY (trip_id) REFERENCES public.trips(id) ON DELETE CASCADE;


--
-- Name: trip_members trip_members_trip_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jazmynstokes
--

ALTER TABLE ONLY public.trip_members
    ADD CONSTRAINT trip_members_trip_id_fkey FOREIGN KEY (trip_id) REFERENCES public.trips(id) ON DELETE CASCADE;


--
-- Name: trip_members trip_members_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jazmynstokes
--

ALTER TABLE ONLY public.trip_members
    ADD CONSTRAINT trip_members_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: trips trips_created_by_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jazmynstokes
--

ALTER TABLE ONLY public.trips
    ADD CONSTRAINT trips_created_by_fkey FOREIGN KEY (created_by) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict GSgquiNM8yvJaFR2mb1OEA01ILDeBVeprcUXNvmua77pmQ8e0TaWbHxFdy7Ymil

