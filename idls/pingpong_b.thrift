namespace go api

struct Request {
        1: string message
}

struct Response {
        1: string message
}

service ServiceB {
    Response serviceB(1: Request req)
}
