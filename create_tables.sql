-- START HERE
CREATE TABLE
    series (
        series_id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        "location" TEXT NOT NULL, -- LOCATION is listed as a reserved word for some databases.
        premiere_date DATE NOT NULL
    );

CREATE TABLE
    seasons (
        season_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        series_id TEXT NOT NULL REFERENCES series (series_id) ON DELETE CASCADE,
        season_number INTEGER NOT NULL,
        premiere_date DATE NOT NULL,
        finale_date DATE -- Can be null if finale hasn't aired yet.
    );

CREATE TABLE
    episodes (
        episode_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        season_id INTEGER NOT NULL REFERENCES seasons (season_id) ON DELETE CASCADE,
        episode_number INTEGER NOT NULL,
        episode_title TEXT NOT NULL,
        air_date DATE NOT NULL
    );

CREATE TABLE
    housewives (
        housewife_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        birth_date DATE NOT NULL,
        image_url TEXT
    );

CREATE TABLE
    seasons_housewife (
        season_housewife_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        season_id INTEGER NOT NULL REFERENCES seasons (season_id) ON DELETE CASCADE,
        housewife_id INTEGER NOT NULL REFERENCES housewives (housewife_id) ON DELETE CASCADE
    );

CREATE TABLE
    quotes (
        quote_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        housewife_id INTEGER NOT NULL REFERENCES housewives (housewife_id) ON DELETE CASCADE,
        content TEXT NOT NULL,
        tagline BOOLEAN NOT NULL
    );