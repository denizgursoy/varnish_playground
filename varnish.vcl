vcl 4.0;

backend default {
    .host = "host.docker.internal";
    .port = "1323";
}

sub vcl_recv {
   set req.backend_hint=default;
}