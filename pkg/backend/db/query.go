package db

// group
const (
	INSERT_GROUP = "INSERT INTO public.\"group\" (name) VALUES ($1) RETURNING id;"

	SELECT_GROUP_ALL = "SELECT id, name FROM public.\"group\";"
)

// location
const (
	INSERT_LOCATION = "INSERT INTO public.\"location\" (name) VALUES ($1) RETURNING id;"

	SELECT_LOCATION_ALL = "SELECT id, name FROM public.\"location\";"
)

// season
const (
	SELECT_SEASON_ALL = "SELECT id, name FROM public.\"season\";"
)

// user
const (
	DELETE_USER_BY_ID = "DELETE FROM public.\"user\" WHERE id = $1;"

	INSERT_USER = "INSERT INTO public.\"user\" (firstname, lastname) VALUES ($1, $2) RETURNING id;"

	SELECT_USER_ALL     = "SELECT id, firstname, lastname FROM public.\"user\";"
	SELECT_USER_BY_ID   = "SELECT id, firstname, lastname FROM public.\"user\" WHERE id = $1;"
	SELECT_USERID_BY_ID = "SELECT id FROM public.\"user\" WHERE id = $1;" // check if user exists
)
