create table container_balls
(
    id        integer not null
        constraint container_balls_pkey
            primary key,
    size      integer,
    player_id integer[],
    is_ready  boolean
);

alter table container_balls
    owner to postgres;

INSERT INTO public.container_balls (id, size, player_id, is_ready) VALUES (2, 4, null, false);
INSERT INTO public.container_balls (id, size, player_id, is_ready) VALUES (3, 8, null, false);
INSERT INTO public.container_balls (id, size, player_id, is_ready) VALUES (4, 2, null, false);
INSERT INTO public.container_balls (id, size, player_id, is_ready) VALUES (1, 5, '', false);
INSERT INTO public.container_balls (id, size, player_id, is_ready) VALUES (5, 3, '', false);