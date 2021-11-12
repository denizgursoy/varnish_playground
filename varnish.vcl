vcl 4.0;

backend default {
    .host = "host.docker.internal";
    .port = "1323";
}

sub vcl_recv {
      if (req.method == "PRI") {
          return (synth(405));
      }
      if (!req.http.host &&
        req.esi_level == 0 &&
        req.proto ~ "^(?i)HTTP/1.1") {
          return (synth(400));
      }

      if (req.method != "GET" &&
        req.method != "HEAD" &&
        req.method != "PUT" &&
        req.method != "POST" &&
        req.method != "TRACE" &&
        req.method != "OPTIONS" &&
        req.method != "DELETE" &&
        req.method != "PATCH") {
                  return (pipe);
        }

      if (req.method != "GET" && req.method != "HEAD") {
          return (pass);
      }
      if (req.http.Authorization || req.http.Cookie) {
          return (pass);
      }
      return (hash);
  }

    sub vcl_hash {
        hash_data(req.url);
        if (req.http.host) {
            hash_data(req.http.host);
        } else {
            hash_data(server.ip);
        }
        return (lookup);
    }


sub vcl_miss {
      return (fetch);
  }

  sub vcl_purge {
        return (synth(200, "Purged"));
    }

sub vcl_pass {
      return (fetch);
  }


  sub vcl_pipe {

        return (pipe);
  }


  sub vcl_synth {
        set resp.http.Content-Type = "text/html; charset=utf-8";
        set resp.http.Retry-After = "5";
        set resp.body = {"<!DOCTYPE html>
    <html>
      <head>
        <title>"} + resp.status + " " + resp.reason + {"</title>
      </head>
      <body>
        <h1>Error "} + resp.status + " " + resp.reason + {"</h1>
        <p>"} + resp.reason + {"</p>
        <h3>Guru Meditation:</h3>
        <p>XID: "} + req.xid + {"</p>
        <hr>
        <p>Varnish cache server</p>
      </body>
    </html>
  "};
        return (deliver);
    }


       sub vcl_deliver {
          return (deliver);
      }


sub vcl_backend_fetch {
      if (bereq.method == "GET") {
          unset bereq.body;
      }
      return (fetch);
  }


