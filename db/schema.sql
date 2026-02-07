SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public;


CREATE TABLE public.products (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    details jsonb,
    price bigint,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    CONSTRAINT products_pkey PRIMARY KEY (id)
);

CREATE TABLE public.users (
    user_id bigint NOT NULL,
    username text,
    balance numeric(10,2) DEFAULT 0 NOT NULL,
    recovery_key text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    firstname character varying(255),
    lastname character varying(255),
    role text DEFAULT 'user'::text,
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

CREATE TABLE public.cart (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id bigint NOT NULL,
    product_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now(),
    CONSTRAINT cart_pkey PRIMARY KEY (id),
    CONSTRAINT cart_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE,
    CONSTRAINT cart_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE
);

CREATE TABLE public.orders (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id bigint,
    product jsonb,
    amount numeric(10,2) NOT NULL,
    order_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT orders_pkey PRIMARY KEY (id)
);

CREATE TABLE public.payments (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id bigint NOT NULL,
    amount numeric(20,8),
    currency text NOT NULL,
    status text DEFAULT 'pending'::text,
    created_at timestamp with time zone DEFAULT now(),
    confirmed_at timestamp with time zone,
    address_in text,
    address_out text,
    txid_in text,
    txid_out text,
    value_coin numeric(20,8),
    value_forwarded_coin numeric(20,8),
    CONSTRAINT wallet_pkey PRIMARY KEY (id)
);