CREATE TABLE parts_of_speech (
    id              bigint,
    word            bigint,
    part_of_speech  varchar(30) NOT NULL,
    UNIQUE(word, part_of_speech),
    CONSTRAINT      fk_word
        FOREIGN KEY (word)
            REFERENCES words(id)
);
