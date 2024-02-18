// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: feeds.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countFeeds = `-- name: CountFeeds :one
SELECT
    COUNT(*)
FROM
    feeds
`

func (q *Queries) CountFeeds(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countFeeds)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countItems = `-- name: CountItems :one
SELECT
    COUNT(*)
FROM
    items
`

func (q *Queries) CountItems(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countItems)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createFeed = `-- name: CreateFeed :one
INSERT INTO
    feeds (
        "url",
        created_at,
        updated_at,
        deleted_at,
        title,
        "description",
        link,
        feed_link,
        links,
        updated,
        updated_parsed,
        published,
        published_parsed,
        "language",
        copyright,
        generator,
        categories,
        custom,
        feed_type,
        feed_version
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20
    )
RETURNING
    id, url, created_at, updated_at, deleted_at, title, description, link, feed_link, links, updated, updated_parsed, published, published_parsed, language, copyright, generator, categories, custom, feed_type, feed_version
`

type CreateFeedParams struct {
	Url             string             `json:"url"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
	DeletedAt       pgtype.Timestamptz `json:"deleted_at"`
	Title           pgtype.Text        `json:"title"`
	Description     pgtype.Text        `json:"description"`
	Link            pgtype.Text        `json:"link"`
	FeedLink        pgtype.Text        `json:"feed_link"`
	Links           []string           `json:"links"`
	Updated         pgtype.Text        `json:"updated"`
	UpdatedParsed   pgtype.Timestamptz `json:"updated_parsed"`
	Published       pgtype.Text        `json:"published"`
	PublishedParsed pgtype.Timestamptz `json:"published_parsed"`
	Language        pgtype.Text        `json:"language"`
	Copyright       pgtype.Text        `json:"copyright"`
	Generator       pgtype.Text        `json:"generator"`
	Categories      []string           `json:"categories"`
	Custom          []byte             `json:"custom"`
	FeedType        pgtype.Text        `json:"feed_type"`
	FeedVersion     pgtype.Text        `json:"feed_version"`
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRow(ctx, createFeed,
		arg.Url,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.Title,
		arg.Description,
		arg.Link,
		arg.FeedLink,
		arg.Links,
		arg.Updated,
		arg.UpdatedParsed,
		arg.Published,
		arg.PublishedParsed,
		arg.Language,
		arg.Copyright,
		arg.Generator,
		arg.Categories,
		arg.Custom,
		arg.FeedType,
		arg.FeedVersion,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Description,
		&i.Link,
		&i.FeedLink,
		&i.Links,
		&i.Updated,
		&i.UpdatedParsed,
		&i.Published,
		&i.PublishedParsed,
		&i.Language,
		&i.Copyright,
		&i.Generator,
		&i.Categories,
		&i.Custom,
		&i.FeedType,
		&i.FeedVersion,
	)
	return i, err
}

const createFeedAuthor = `-- name: CreateFeedAuthor :one
INSERT INTO
    feed_authors (
        created_at,
        updated_at,
        deleted_at,
        "name",
        email,
        feed_id
    )
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING
    id, created_at, updated_at, deleted_at, name, email, feed_id
`

type CreateFeedAuthorParams struct {
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	Name      pgtype.Text        `json:"name"`
	Email     pgtype.Text        `json:"email"`
	FeedID    int64              `json:"feed_id"`
}

func (q *Queries) CreateFeedAuthor(ctx context.Context, arg CreateFeedAuthorParams) (FeedAuthor, error) {
	row := q.db.QueryRow(ctx, createFeedAuthor,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.Name,
		arg.Email,
		arg.FeedID,
	)
	var i FeedAuthor
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.Email,
		&i.FeedID,
	)
	return i, err
}

const createFeedDublinCore = `-- name: CreateFeedDublinCore :one
INSERT INTO
    feed_dublin_cores (
        created_at,
        updated_at,
        deleted_at,
        title,
        creator,
        author,
        "subject",
        "description",
        publisher,
        contributor,
        "date",
        "type",
        format,
        identifier,
        source,
        "language",
        relation,
        coverage,
        rights,
        feed_id
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20
    )
RETURNING
    id, created_at, updated_at, deleted_at, title, creator, author, subject, description, publisher, contributor, date, type, format, identifier, source, language, relation, coverage, rights, feed_id
`

type CreateFeedDublinCoreParams struct {
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	DeletedAt   pgtype.Timestamptz `json:"deleted_at"`
	Title       []string           `json:"title"`
	Creator     []string           `json:"creator"`
	Author      []string           `json:"author"`
	Subject     []string           `json:"subject"`
	Description []string           `json:"description"`
	Publisher   []string           `json:"publisher"`
	Contributor []string           `json:"contributor"`
	Date        []string           `json:"date"`
	Type        []string           `json:"type"`
	Format      []string           `json:"format"`
	Identifier  []string           `json:"identifier"`
	Source      []string           `json:"source"`
	Language    []string           `json:"language"`
	Relation    []string           `json:"relation"`
	Coverage    []string           `json:"coverage"`
	Rights      []string           `json:"rights"`
	FeedID      int64              `json:"feed_id"`
}

func (q *Queries) CreateFeedDublinCore(ctx context.Context, arg CreateFeedDublinCoreParams) (FeedDublinCore, error) {
	row := q.db.QueryRow(ctx, createFeedDublinCore,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.Title,
		arg.Creator,
		arg.Author,
		arg.Subject,
		arg.Description,
		arg.Publisher,
		arg.Contributor,
		arg.Date,
		arg.Type,
		arg.Format,
		arg.Identifier,
		arg.Source,
		arg.Language,
		arg.Relation,
		arg.Coverage,
		arg.Rights,
		arg.FeedID,
	)
	var i FeedDublinCore
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Creator,
		&i.Author,
		&i.Subject,
		&i.Description,
		&i.Publisher,
		&i.Contributor,
		&i.Date,
		&i.Type,
		&i.Format,
		&i.Identifier,
		&i.Source,
		&i.Language,
		&i.Relation,
		&i.Coverage,
		&i.Rights,
		&i.FeedID,
	)
	return i, err
}

const createFeedExtension = `-- name: CreateFeedExtension :one
INSERT INTO
    feed_extensions (
        created_at,
        updated_at,
        deleted_at,
        "name",
        "value",
        attrs,
        children,
        feed_id
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    id, created_at, updated_at, deleted_at, name, value, attrs, children, feed_id
`

type CreateFeedExtensionParams struct {
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	Name      pgtype.Text        `json:"name"`
	Value     pgtype.Text        `json:"value"`
	Attrs     []byte             `json:"attrs"`
	Children  []byte             `json:"children"`
	FeedID    int64              `json:"feed_id"`
}

func (q *Queries) CreateFeedExtension(ctx context.Context, arg CreateFeedExtensionParams) (FeedExtension, error) {
	row := q.db.QueryRow(ctx, createFeedExtension,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.Name,
		arg.Value,
		arg.Attrs,
		arg.Children,
		arg.FeedID,
	)
	var i FeedExtension
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.Value,
		&i.Attrs,
		&i.Children,
		&i.FeedID,
	)
	return i, err
}

const createFeedImage = `-- name: CreateFeedImage :one
INSERT INTO
    feed_images (
        created_at,
        updated_at,
        deleted_at,
        "url",
        title,
        feed_id
    )
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING
    id, created_at, updated_at, deleted_at, url, title, feed_id
`

type CreateFeedImageParams struct {
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	Url       pgtype.Text        `json:"url"`
	Title     pgtype.Text        `json:"title"`
	FeedID    int64              `json:"feed_id"`
}

func (q *Queries) CreateFeedImage(ctx context.Context, arg CreateFeedImageParams) (FeedImage, error) {
	row := q.db.QueryRow(ctx, createFeedImage,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.Url,
		arg.Title,
		arg.FeedID,
	)
	var i FeedImage
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Url,
		&i.Title,
		&i.FeedID,
	)
	return i, err
}

const createItem = `-- name: CreateItem :one
INSERT INTO
    items (
        created_at,
        updated_at,
        deleted_at,
        title,
        "description",
        content,
        link,
        links,
        updated,
        updated_parsed,
        published,
        published_parsed,
        "guid",
        categories,
        custom,
        feed_id
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16
    )
RETURNING
    id, created_at, updated_at, deleted_at, title, description, content, link, links, updated, updated_parsed, published, published_parsed, guid, categories, custom, feed_id
`

type CreateItemParams struct {
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
	DeletedAt       pgtype.Timestamptz `json:"deleted_at"`
	Title           pgtype.Text        `json:"title"`
	Description     pgtype.Text        `json:"description"`
	Content         pgtype.Text        `json:"content"`
	Link            pgtype.Text        `json:"link"`
	Links           []string           `json:"links"`
	Updated         pgtype.Text        `json:"updated"`
	UpdatedParsed   pgtype.Timestamptz `json:"updated_parsed"`
	Published       pgtype.Text        `json:"published"`
	PublishedParsed pgtype.Timestamptz `json:"published_parsed"`
	Guid            pgtype.Text        `json:"guid"`
	Categories      []string           `json:"categories"`
	Custom          []byte             `json:"custom"`
	FeedID          int64              `json:"feed_id"`
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRow(ctx, createItem,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.Title,
		arg.Description,
		arg.Content,
		arg.Link,
		arg.Links,
		arg.Updated,
		arg.UpdatedParsed,
		arg.Published,
		arg.PublishedParsed,
		arg.Guid,
		arg.Categories,
		arg.Custom,
		arg.FeedID,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.Link,
		&i.Links,
		&i.Updated,
		&i.UpdatedParsed,
		&i.Published,
		&i.PublishedParsed,
		&i.Guid,
		&i.Categories,
		&i.Custom,
		&i.FeedID,
	)
	return i, err
}

const createItemAuthor = `-- name: CreateItemAuthor :one
INSERT INTO
    item_authors (
        created_at,
        updated_at,
        deleted_at,
        "name",
        email,
        item_id
    )
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING
    id, created_at, updated_at, deleted_at, name, email, item_id
`

type CreateItemAuthorParams struct {
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	Name      pgtype.Text        `json:"name"`
	Email     pgtype.Text        `json:"email"`
	ItemID    int64              `json:"item_id"`
}

func (q *Queries) CreateItemAuthor(ctx context.Context, arg CreateItemAuthorParams) (ItemAuthor, error) {
	row := q.db.QueryRow(ctx, createItemAuthor,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.Name,
		arg.Email,
		arg.ItemID,
	)
	var i ItemAuthor
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.Email,
		&i.ItemID,
	)
	return i, err
}

const createItemDublinCore = `-- name: CreateItemDublinCore :one
INSERT INTO
    item_dublin_cores (
        created_at,
        updated_at,
        deleted_at,
        title,
        creator,
        author,
        "subject",
        "description",
        publisher,
        contributor,
        "date",
        "type",
        format,
        identifier,
        source,
        "language",
        relation,
        coverage,
        rights,
        item_id
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20
    )
RETURNING
    id, created_at, updated_at, deleted_at, title, creator, author, subject, description, publisher, contributor, date, type, format, identifier, source, language, relation, coverage, rights, item_id
`

type CreateItemDublinCoreParams struct {
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	DeletedAt   pgtype.Timestamptz `json:"deleted_at"`
	Title       []string           `json:"title"`
	Creator     []string           `json:"creator"`
	Author      []string           `json:"author"`
	Subject     []string           `json:"subject"`
	Description []string           `json:"description"`
	Publisher   []string           `json:"publisher"`
	Contributor []string           `json:"contributor"`
	Date        []string           `json:"date"`
	Type        []string           `json:"type"`
	Format      []string           `json:"format"`
	Identifier  []string           `json:"identifier"`
	Source      []string           `json:"source"`
	Language    []string           `json:"language"`
	Relation    []string           `json:"relation"`
	Coverage    []string           `json:"coverage"`
	Rights      []string           `json:"rights"`
	ItemID      int64              `json:"item_id"`
}

func (q *Queries) CreateItemDublinCore(ctx context.Context, arg CreateItemDublinCoreParams) (ItemDublinCore, error) {
	row := q.db.QueryRow(ctx, createItemDublinCore,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.Title,
		arg.Creator,
		arg.Author,
		arg.Subject,
		arg.Description,
		arg.Publisher,
		arg.Contributor,
		arg.Date,
		arg.Type,
		arg.Format,
		arg.Identifier,
		arg.Source,
		arg.Language,
		arg.Relation,
		arg.Coverage,
		arg.Rights,
		arg.ItemID,
	)
	var i ItemDublinCore
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Creator,
		&i.Author,
		&i.Subject,
		&i.Description,
		&i.Publisher,
		&i.Contributor,
		&i.Date,
		&i.Type,
		&i.Format,
		&i.Identifier,
		&i.Source,
		&i.Language,
		&i.Relation,
		&i.Coverage,
		&i.Rights,
		&i.ItemID,
	)
	return i, err
}

const createItemExtension = `-- name: CreateItemExtension :one
INSERT INTO
    item_extensions (
        created_at,
        updated_at,
        deleted_at,
        "name",
        "value",
        attrs,
        children,
        item_id
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    id, created_at, updated_at, deleted_at, name, value, attrs, children, item_id
`

type CreateItemExtensionParams struct {
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	Name      pgtype.Text        `json:"name"`
	Value     pgtype.Text        `json:"value"`
	Attrs     []byte             `json:"attrs"`
	Children  []byte             `json:"children"`
	ItemID    int64              `json:"item_id"`
}

func (q *Queries) CreateItemExtension(ctx context.Context, arg CreateItemExtensionParams) (ItemExtension, error) {
	row := q.db.QueryRow(ctx, createItemExtension,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.Name,
		arg.Value,
		arg.Attrs,
		arg.Children,
		arg.ItemID,
	)
	var i ItemExtension
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.Value,
		&i.Attrs,
		&i.Children,
		&i.ItemID,
	)
	return i, err
}

const createItemImage = `-- name: CreateItemImage :one
INSERT INTO
    item_images (
        created_at,
        updated_at,
        deleted_at,
        "url",
        title,
        item_id
    )
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING
    id, created_at, updated_at, deleted_at, url, title, item_id
`

type CreateItemImageParams struct {
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	Url       pgtype.Text        `json:"url"`
	Title     pgtype.Text        `json:"title"`
	ItemID    int64              `json:"item_id"`
}

func (q *Queries) CreateItemImage(ctx context.Context, arg CreateItemImageParams) (ItemImage, error) {
	row := q.db.QueryRow(ctx, createItemImage,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.Url,
		arg.Title,
		arg.ItemID,
	)
	var i ItemImage
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Url,
		&i.Title,
		&i.ItemID,
	)
	return i, err
}

const getFeed = `-- name: GetFeed :one
SELECT
    id, url, created_at, updated_at, deleted_at, title, description, link, feed_link, links, updated, updated_parsed, published, published_parsed, language, copyright, generator, categories, custom, feed_type, feed_version
FROM
    feeds
WHERE
    id = $1
`

func (q *Queries) GetFeed(ctx context.Context, id int64) (Feed, error) {
	row := q.db.QueryRow(ctx, getFeed, id)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Description,
		&i.Link,
		&i.FeedLink,
		&i.Links,
		&i.Updated,
		&i.UpdatedParsed,
		&i.Published,
		&i.PublishedParsed,
		&i.Language,
		&i.Copyright,
		&i.Generator,
		&i.Categories,
		&i.Custom,
		&i.FeedType,
		&i.FeedVersion,
	)
	return i, err
}

const getFeedAuthors = `-- name: GetFeedAuthors :many
SELECT
    id, created_at, updated_at, deleted_at, name, email, feed_id
FROM
    feed_authors
WHERE
    feed_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetFeedAuthorsParams struct {
	FeedID int64 `json:"feed_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetFeedAuthors(ctx context.Context, arg GetFeedAuthorsParams) ([]FeedAuthor, error) {
	rows, err := q.db.Query(ctx, getFeedAuthors, arg.FeedID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FeedAuthor{}
	for rows.Next() {
		var i FeedAuthor
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
			&i.Email,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeedDublinCores = `-- name: GetFeedDublinCores :many
SELECT
    id, created_at, updated_at, deleted_at, title, creator, author, subject, description, publisher, contributor, date, type, format, identifier, source, language, relation, coverage, rights, feed_id
FROM
    feed_dublin_cores
WHERE
    feed_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetFeedDublinCoresParams struct {
	FeedID int64 `json:"feed_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetFeedDublinCores(ctx context.Context, arg GetFeedDublinCoresParams) ([]FeedDublinCore, error) {
	rows, err := q.db.Query(ctx, getFeedDublinCores, arg.FeedID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FeedDublinCore{}
	for rows.Next() {
		var i FeedDublinCore
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Title,
			&i.Creator,
			&i.Author,
			&i.Subject,
			&i.Description,
			&i.Publisher,
			&i.Contributor,
			&i.Date,
			&i.Type,
			&i.Format,
			&i.Identifier,
			&i.Source,
			&i.Language,
			&i.Relation,
			&i.Coverage,
			&i.Rights,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeedExtensions = `-- name: GetFeedExtensions :many
SELECT
    id, created_at, updated_at, deleted_at, name, value, attrs, children, feed_id
FROM
    feed_extensions
WHERE
    feed_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetFeedExtensionsParams struct {
	FeedID int64 `json:"feed_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetFeedExtensions(ctx context.Context, arg GetFeedExtensionsParams) ([]FeedExtension, error) {
	rows, err := q.db.Query(ctx, getFeedExtensions, arg.FeedID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FeedExtension{}
	for rows.Next() {
		var i FeedExtension
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
			&i.Value,
			&i.Attrs,
			&i.Children,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeedImages = `-- name: GetFeedImages :many
SELECT
    id, created_at, updated_at, deleted_at, url, title, feed_id
FROM
    feed_images
WHERE
    feed_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetFeedImagesParams struct {
	FeedID int64 `json:"feed_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetFeedImages(ctx context.Context, arg GetFeedImagesParams) ([]FeedImage, error) {
	rows, err := q.db.Query(ctx, getFeedImages, arg.FeedID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FeedImage{}
	for rows.Next() {
		var i FeedImage
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Url,
			&i.Title,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeeds = `-- name: GetFeeds :many
SELECT
    id, url, created_at, updated_at, deleted_at, title, description, link, feed_link, links, updated, updated_parsed, published, published_parsed, language, copyright, generator, categories, custom, feed_type, feed_version
FROM
    feeds
ORDER BY
    created_at DESC
LIMIT
    $1
OFFSET
    $2
`

type GetFeedsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetFeeds(ctx context.Context, arg GetFeedsParams) ([]Feed, error) {
	rows, err := q.db.Query(ctx, getFeeds, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Feed{}
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Title,
			&i.Description,
			&i.Link,
			&i.FeedLink,
			&i.Links,
			&i.Updated,
			&i.UpdatedParsed,
			&i.Published,
			&i.PublishedParsed,
			&i.Language,
			&i.Copyright,
			&i.Generator,
			&i.Categories,
			&i.Custom,
			&i.FeedType,
			&i.FeedVersion,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItem = `-- name: GetItem :one
SELECT
    id, created_at, updated_at, deleted_at, title, description, content, link, links, updated, updated_parsed, published, published_parsed, guid, categories, custom, feed_id
FROM
    items
WHERE
    id = $1
`

func (q *Queries) GetItem(ctx context.Context, id int64) (Item, error) {
	row := q.db.QueryRow(ctx, getItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.Link,
		&i.Links,
		&i.Updated,
		&i.UpdatedParsed,
		&i.Published,
		&i.PublishedParsed,
		&i.Guid,
		&i.Categories,
		&i.Custom,
		&i.FeedID,
	)
	return i, err
}

const getItemAuthors = `-- name: GetItemAuthors :many
SELECT
    id, created_at, updated_at, deleted_at, name, email, item_id
FROM
    item_authors
WHERE
    item_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetItemAuthorsParams struct {
	ItemID int64 `json:"item_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetItemAuthors(ctx context.Context, arg GetItemAuthorsParams) ([]ItemAuthor, error) {
	rows, err := q.db.Query(ctx, getItemAuthors, arg.ItemID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ItemAuthor{}
	for rows.Next() {
		var i ItemAuthor
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
			&i.Email,
			&i.ItemID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemDublinCores = `-- name: GetItemDublinCores :many
SELECT
    id, created_at, updated_at, deleted_at, title, creator, author, subject, description, publisher, contributor, date, type, format, identifier, source, language, relation, coverage, rights, item_id
FROM
    item_dublin_cores
WHERE
    item_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetItemDublinCoresParams struct {
	ItemID int64 `json:"item_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetItemDublinCores(ctx context.Context, arg GetItemDublinCoresParams) ([]ItemDublinCore, error) {
	rows, err := q.db.Query(ctx, getItemDublinCores, arg.ItemID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ItemDublinCore{}
	for rows.Next() {
		var i ItemDublinCore
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Title,
			&i.Creator,
			&i.Author,
			&i.Subject,
			&i.Description,
			&i.Publisher,
			&i.Contributor,
			&i.Date,
			&i.Type,
			&i.Format,
			&i.Identifier,
			&i.Source,
			&i.Language,
			&i.Relation,
			&i.Coverage,
			&i.Rights,
			&i.ItemID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemExtensions = `-- name: GetItemExtensions :many
SELECT
    id, created_at, updated_at, deleted_at, name, value, attrs, children, item_id
FROM
    item_extensions
WHERE
    item_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetItemExtensionsParams struct {
	ItemID int64 `json:"item_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetItemExtensions(ctx context.Context, arg GetItemExtensionsParams) ([]ItemExtension, error) {
	rows, err := q.db.Query(ctx, getItemExtensions, arg.ItemID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ItemExtension{}
	for rows.Next() {
		var i ItemExtension
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
			&i.Value,
			&i.Attrs,
			&i.Children,
			&i.ItemID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemImages = `-- name: GetItemImages :many
SELECT
    id, created_at, updated_at, deleted_at, url, title, item_id
FROM
    item_images
WHERE
    item_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetItemImagesParams struct {
	ItemID int64 `json:"item_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetItemImages(ctx context.Context, arg GetItemImagesParams) ([]ItemImage, error) {
	rows, err := q.db.Query(ctx, getItemImages, arg.ItemID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ItemImage{}
	for rows.Next() {
		var i ItemImage
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Url,
			&i.Title,
			&i.ItemID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItems = `-- name: GetItems :many
SELECT
    id, created_at, updated_at, deleted_at, title, description, content, link, links, updated, updated_parsed, published, published_parsed, guid, categories, custom, feed_id
FROM
    items
WHERE
    feed_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetItemsParams struct {
	FeedID int64 `json:"feed_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetItems(ctx context.Context, arg GetItemsParams) ([]Item, error) {
	rows, err := q.db.Query(ctx, getItems, arg.FeedID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Item{}
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.Link,
			&i.Links,
			&i.Updated,
			&i.UpdatedParsed,
			&i.Published,
			&i.PublishedParsed,
			&i.Guid,
			&i.Categories,
			&i.Custom,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
