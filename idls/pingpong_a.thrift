namespace go api

struct Request {
        1: string message
}

struct Response {
        1: string message
}

service ServiceA {
    Response serviceA(1: Request req)
}
