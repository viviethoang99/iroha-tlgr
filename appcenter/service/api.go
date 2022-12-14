package serviceac

func CreateAPIWithClientParams(token string) *api {
	return &api{client: NewClient(token)}
}

type api struct {
	client Client
}
