INSERT INTO maat_news_dev.rss_feeds (
    id,
    title,
    link,
    description,
    author,
    categories,
    source,
    published_at,
    loaded_at
)
SELECT
    payload->>'id' AS id,
    payload->>'title' AS title,
    payload->>'link' AS link,
    payload->>'description' AS description,
    payload->>'author' AS author,
    ARRAY(SELECT jsonb_array_elements_text(payload->'categories')) AS categories,
    payload->>'source' AS source,
    payload->>'published_at' AS published_at,
    loaded_at
FROM maat_news_dev.rss_feeds_raw;
