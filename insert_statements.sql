INSERT INTO
    series (series_id, title, "location", premiere_date)
VALUES
    (
        'RHOBH',
        'Real Housewives of Beverly Hills',
        'California',
        '2010-10-14'
    ),
    (
        'RHOA',
        'Real Housewives of Atlanta',
        'Georgia',
        '2008-10-07'
    );

INSERT INTO
    seasons (
        series_id,
        season_number,
        premiere_date,
        finale_date
    )
VALUES
    ('RHOBH', 1, '2010-10-14', '2011-02-15'),
    ('RHOA', 1, '2008-10-07', '2008-11-25');

INSERT INTO
    episodes (
        season_id,
        episode_number,
        episode_title,
        air_date
    )
VALUES
    (
        (
            SELECT
                season_id
            FROM
                seasons
            WHERE
                series_id = 'RHOBH' AND season_number = 1
        ),
        1,
        'Life, Liberty and the Pursuit of Wealthiness',
        '2010-10-14'
    ),
    (
        (
            SELECT
                season_id
            FROM
                seasons
            WHERE
                series_id = 'RHOA' AND season_number = 1
        ),
        1,
        'Welcome One, Welcome ATL',
        '2008-10-07'
    );

INSERT INTO
    housewives (first_name, last_name, birth_date, image_url)
VALUES
    ('Kandi', 'Burruss', '1976-05-17', 'image.jpg'),
    ('Kyle', 'Richards', '1969-01-11', '');

INSERT INTO
    seasons_housewife (season_id, housewife_id)
VALUES
    (
        (
            SELECT
                season_id
            FROM
                seasons
            WHERE
                series_id = 'RHOBH'
                AND season_number = 1
        ),
        (
            SELECT
                housewife_id
            FROM
                housewives
            WHERE
                first_name = 'Kyle'
                AND last_name = 'Richards'
        )
    ),
    (
        (
            SELECT
                season_id
            FROM
                seasons
            WHERE
                series_id = 'RHOA'
                AND season_number = 1
        ),
        (
            SELECT
                housewife_id
            FROM
                housewives
            WHERE
                first_name = 'Kandi'
                AND last_name = 'Burruss'
        )
    );

INSERT INTO
    quotes (housewife_id, content, tagline)
VALUES
    (
        (
            SELECT
                housewife_id
            FROM
                housewives
            WHERE
                first_name = 'Kyle'
                AND last_name = 'Richards'
        ),
        'You''re such a liar, Camille!',
        'FALSE'
    ),
    (
        (
            SELECT
                housewife_id
            FROM
                housewives
            WHERE
                first_name = 'Kandi'
                AND last_name = 'Burruss'
        ),
        'I''m not about the drama. Don''t start none, won''t be none',
        'TRUE'
    );