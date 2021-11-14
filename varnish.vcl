vcl 4.1;
/*
acl purge {
      "host.docker.internal";
      "localhost";
      "192.168.65.2";
}
*/
backend default {
  .host = "host.docker.internal";
  .port = "1323";
}

sub vcl_recv{
       if (req.method == "PURGE") {
       /*
                if (client.ip !~ purge) {
                        return(synth(405,"Not allowed."));
                }
                */
                return (purge);
        }

    return(hash);
}
