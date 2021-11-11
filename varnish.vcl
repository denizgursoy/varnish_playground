vcl 4.0;

backend default {
  .host = "www.nytimes.com:80";
}

sub vcl_recv {
    if(req.method == "GET"){
        return (synth(404));
    }
}