This is a starter project for a proxy to DialogFlow.com.
To start call
    >>> go run  <path_to_go>/go/src/leo/main.go

To enable demo chat at <base_url>:8080 use flag: -chat

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