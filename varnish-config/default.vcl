#
# This is an example VCL file for Varnish.
#
# It does not do anything by default, delegating control to the
# builtin VCL. The builtin VCL is called when there is no explicit
# return statement.
#
# See the VCL chapters in the Users Guide at https://www.varnish-cache.org/docs/
# and http://varnish-cache.org/trac/wiki/VCLExamples for more examples.

# Marker to tell the VCL compiler that this VCL has been adapted to the
# new 4.0 format.
vcl 4.0;

# Default backend definition. Set this to point to your content server.
#backend default {
#    .host = "127.0.0.1";
#    .port = "8989";
#}

backend lil1 {
    .host = "127.0.0.1";
    .port = "8989";
}

backend lil2 {
   .host = "127.0.0.1";
   .port = "8888";
}

import directors;
sub vcl_init {
  new cluster = directors.round_robin();
  cluster.add_backend(lil1);
  cluster.add_backend(lil2);
}

sub vcl_recv {
    if (req.method == "POST") {
        set req.backend_hint = cluster.backend();
        return (pass);
    } 
   
    if (req.url ~ "e$") {
      set req.backend_hint = lil2;
    } 
    
    if (req.url ~ "f$") {
      set req.backend_hint = lil1;
    }
   
    if (req.url ~ "notfound/") {
      return (pass);
    }	
    
    # Happens before we check if we have this in cache already.
    # 
    # Typically you clean up the request here, removing cookies you don't need,
    # rewriting the request, etc.
}

sub vcl_backend_response {
   if (beresp.status == 404) {
   	set beresp.ttl = 0s; 
   }

    # Happens after we have read the response headers from the backend.
    # 
    # Here you clean the response headers, removing silly Set-Cookie headers
    # and other mistakes your backend does.
}

sub vcl_deliver {
    # Happens when we have all the pieces we need, and are about to send the
    # response to the client.
    # 
    # You can do accounting or modifying the final object here.
}

