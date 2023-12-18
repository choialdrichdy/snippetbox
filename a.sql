INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...
    A frog jumps into the pond,
    splash! Silence again.
    
    – Matsuo Bashō',
    DATETIME('now'),
    DATETIME(DATETIME('now'), '+1 year')
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry
    forest, winds howl in rage
    with no leaves to blow.
    
    – Natsume Soseki',
    DATETIME('now'),
    DATETIME(DATETIME('now'), '+1 year')
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning
    the mirror I stare into
    shows my father''s face.
    
    – Murakami Kijo',
    DATETIME('now'),
    DATETIME(DATETIME('now'), '+7 days')
);