This is a starter project for a proxy to https://dialogflow.com/
To start call:

    >>> go run  <path_to_go>/go/src/leo/main.go

To enable demo chat at http://<base_url>:8080 use flag: -chat

Example:

    >>> go run  <path_to_go>/go/src/leo/main.go -chat.

If you run this on your local machine, go to:

    http://localhost:8080 to test.

To enable <askleo> endpoint use flag: -leo

Example:

    >>> go run  <path_to_go>/go/src/leo/main.go -leo.

After service initialization, to get a simple text response from leo agent go to:

    http://<base_url>:8080/askleo/<text_to_be_processed>

Note: <user_entered_text> will be sent using GET method.

To view documentation:

    >>> godoc -http=:6060 -goroot=<path_to_leo_package>

This will display docs at http://localhost:6060/pkg/leo
If you need to get a text representation of a particular
page just add:

    ?m=all,text
    
to the page url.    
