CREATE TABLE public.users
(
    email varchar(100) NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (email)
);

CREATE TABLE public.block
(
    id        int8         NOT NULL GENERATED ALWAYS AS IDENTITY,
    requestor varchar(100) NULL,
    target    varchar(100) NULL,
    CONSTRAINT block_pkey PRIMARY KEY (id),
    CONSTRAINT block_requestor_fkey FOREIGN KEY (requestor) REFERENCES users (email),
    CONSTRAINT block_target_fkey FOREIGN KEY (target) REFERENCES users (email)
);

CREATE TABLE public.friend
(
    id           int8         NOT NULL GENERATED ALWAYS AS IDENTITY,
    emailuserone varchar(100) NULL,
    emailusertwo varchar(100) NULL,
    CONSTRAINT friend_pkey PRIMARY KEY (id),
    CONSTRAINT friend_emailuserone_fkey FOREIGN KEY (emailuserone) REFERENCES users (email),
    CONSTRAINT friend_emailusertwo_fkey FOREIGN KEY (emailusertwo) REFERENCES users (email)
);

CREATE TABLE public."subscription"
(
    id        int8         NOT NULL GENERATED ALWAYS AS IDENTITY,
    requestor varchar(100) NULL,
    target    varchar(100) NULL,
    CONSTRAINT subscription_pkey PRIMARY KEY (id),
    CONSTRAINT subscription_requestor_fkey FOREIGN KEY (requestor) REFERENCES users (email),
    CONSTRAINT subscription_target_fkey FOREIGN KEY (target) REFERENCES users (email)
);
